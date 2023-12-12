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

func main() {
	db := initDB("example.db")
	e := echo.New()
	e.Use(middleware.Logger())

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
