package repository

import (
	"database/sql"
	"todoapp/model"
)

type TaskRepository interface {
	Create(task *model.Task) (int, error)
	Read(id int) (*model.Task, error)
	Update(task *model.Task) error
	Delete(id int) error
}

// privateな構造体を定義することで、外部からの直接のインスタンス化を防ぐ
// dbはsql.DBのポインタを持つ
type taskRepositoryImpl struct {
	db *sql.DB
}

// このpublicの関数で外部からインスタンス化できるようにする
func NewTaskRepository(db *sql.DB) *taskRepositoryImpl {
	// {db: db}で構造体の初期化を行いそのポインタを返す
	return &taskRepositoryImpl{db: db}
}

func (r *taskRepositoryImpl) Create(task *model.Task) (int, error) {
	stmt := `INSERT INTO tasks (title) VALUES (?) RETURNING id`
	// queryRowで実行してScanでtask.IDに値を入れる
	err := r.db.QueryRow(stmt, task.Title).Scan(&task.ID)
	return task.ID, err
}

func (r *taskRepositoryImpl) Read(id int) (*model.Task, error) {
	stmt := `SELECT id, title FROM tasks WHERE id = ?`
	task := model.Task{}
	// queryRowで実行してScanでtask.IDとtask.Titleに値を入れる
	err := r.db.QueryRow(stmt, id).Scan(&task.ID, &task.Title)
	return &task, err
}

func (r *taskRepositoryImpl) Update(task *model.Task) error {
	stmt := `UPDATE tasks SET title = ? WHERE id = ?`
	rows, err := r.db.Exec(stmt, task.Title, task.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return err
	}
	// 更新対象のレコードがない場合はsql.ErrNoRowsを返す
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (r *taskRepositoryImpl) Delete(id int) error {
	stmt := `DELETE FROM tasks WHERE id = ?`
	rows, err := r.db.Exec(stmt, id)
	if err != nil {
		return err
	}

	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return err
	}
	// 削除対象のレコードがない場合はsql.ErrNoRowsを返す
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}
