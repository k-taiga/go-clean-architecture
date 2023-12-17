package main

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	// go-sqlite3を明示的に使っていないが内部的に利用をするために _ でimportする
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	// json:"id"でjsonに変換したときのオブジェクトの構造を定義
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func validateUser(name string, age int) error {
	if name == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "name is empty")
	}
	if len(name) > 100 {
		return echo.NewHTTPError(http.StatusBadRequest, "name is too long")
	}
	if age < 0 || age >= 200 {
		return echo.NewHTTPError(http.StatusBadRequest, "age must be between 0 and 200")
	}
	return nil
}

func main() {
	db := initDB("example.db")
	e := echo.New()
	e.Use(middleware.Logger())

	e.DELETE("/users/:id", func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		result, err := db.Exec("DELETE FROM users WHERE id = ?", id)

		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		rowsAffected, _ := result.RowsAffected()
		if rowsAffected == 0 {
			return echo.NewHTTPError(http.StatusNotFound, "not found")
		}
		// 204を返す
		return c.NoContent(http.StatusNoContent)
	})

	e.POST("/users", func(c echo.Context) error {
		// form経由の値を取得
		name := c.FormValue("name")
		// ageが文字列で来るのでstrconv ASCIItointeger(Atoi)でint型に変換する
		age, _ := strconv.Atoi(c.FormValue("age"))

		// sql.Result型のオブジェクトが返り値
		result, err := db.Exec("INSERT INTO users(name,age) VALUES (?,?)", name, age)

		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		id, _ := result.LastInsertId()

		// &Userはnew(User)とほぼ同じで生成されたオブジェクトのポインタを返すがstructの中身を指定して生成できる
		return c.JSON(http.StatusOK, &User{ID: int(id), Name: name, Age: age})
	})

	e.GET("/users", func(c echo.Context) error {
		rows, err := db.Query("SELECT id,name,age FROM users")
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		// rowsの中にdbQueryのコネクションなどのリソースも含まれているため不要になったらクローズ
		defer rows.Close()
		// User{}型のスライス
		users := []User{}

		// １レコードずつ回す
		for rows.Next() {
			var user User
			// rows.Scanでクエリで取得したID,Name,AgeをUser{}にバインド
			if err := rows.Scan(&user.ID, &user.Name, &user.Age); err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			users = append(users, user)
		}
		return c.JSON(http.StatusOK, users)
	})

	e.PUT("/users/:id", func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		name := c.FormValue("name")
		age, err := strconv.Atoi(c.FormValue("age"))

		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		// 右側がif文の条件式で左側がその実行式の前に実行するもの
		if err := validateUser(name, age); err != nil {
			return err
		}

		result, err := db.Exec("UPDATE users SET name = ?, age = ? where id = ?", name, age, id)

		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		// RowsAffected()で変更されたレコード数を取得
		rows, _ := result.RowsAffected()

		if rows == 0 {
			return echo.NewHTTPError(http.StatusNotFound, "not found")
		}

		// &Userはnew(User)とほぼ同じで生成されたオブジェクトのポインタを返すがstructの中身を指定して生成できる
		return c.JSON(http.StatusOK, &User{ID: id, Name: name, Age: age})
	})

	e.GET("/users/:id", func(c echo.Context) error {
		// c.Paramでパラメータを取得
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		// QueryRowで1レコードを取得
		row := db.QueryRow("SELECT id,name,age FROM users WHERE id = ?", id)

		var user User
		// Scanで取得したレコードをUser{}にバインド
		if err := row.Scan(&user.ID, &user.Name, &user.Age); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, user)
	})

	e.Start(":8080")

	// db, err := sql.Open("sqlite3", "./example.db")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // dbへの接続を最後閉じる
	// defer db.Close()

	// // ``だと複数行の文字列を生成できる
	// // ↓shift+fn+tabでcaps lockできるよ
	// createTableSQL := `CREATE TABLE IF NOT EXISTS users(
	// 	id INTEGER PRIMARY KEY AUTOINCREMENT,
	// 	name TEXT NOT NULL,
	// 	age INTEGER NOT NULL
	// );
	// `

	// _, err = db.Exec(createTableSQL)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println("Table created")
}
