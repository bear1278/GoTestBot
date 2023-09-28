package Mysql

import (
	"database/sql"
	"log"
	"time"

	tg "github.com/bear1278/GoTestBot/pkg/telegram"
)

type DbRepository struct {
	db *sql.DB
}

func NewDbRepository(db *sql.DB) *DbRepository {
	return &DbRepository{db: db}
}

func (d *DbRepository) CheckAlreadyMark(ChatID int64, task string) (bool, error) {
	rows, err := d.db.Query("select mark from tasks where chatid=? and name=?", ChatID, task)
	if err != nil {
		return false, err
	}
	defer rows.Close()
	mark := ""
	for rows.Next() {
		err = rows.Scan(&mark)
		if err != nil {
			return false, err
		}
	}
	if mark == "✅" {
		return true, nil
	} else {
		return false, nil
	}
}

func (d *DbRepository) CheckTask(ChatId int64, task string) (bool, error) {
	rows, err := d.db.Query("select name from tasks where chatid=? and name=?", ChatId, task)
	if err != nil {
		return false, err
	}
	defer rows.Close()
	name := ""
	for rows.Next() {
		err = rows.Scan(&name)
		if err != nil {
			return false, err
		}
	}
	if name != "" {
		return true, nil
	} else {
		return false, nil
	}
}

func (d *DbRepository) List(ChatID int64) (string, error) {
	rows, err := d.db.Query("select name, date, mark from tasks where chatid=?", ChatID)

	if err != nil {
		return "", err
	}
	var date string
	defer rows.Close()
	user := tg.NewUser(ChatID)
	for rows.Next() {
		t := tg.Task{}

		err := rows.Scan(&t.Name, &date, &t.Mark)
		t.Date, _ = time.Parse("2006-01-02", date)

		if err != nil {
			log.Print(err)
			continue
		}
		user.SetTask(t.Name, t.Date, t.Mark)

	}
	return user.GetListTask(), nil

}

func (d *DbRepository) Add(ChatID int64, task string) error {

	_, err := d.db.Exec("Insert into tasks values (?,?,?,?)", ChatID, task, time.Now().Format("2006-01-02"), "❗️")

	return err
}

func (d *DbRepository) Delete(ChatID int64, task string) error {
	_, err := d.db.Exec("Delete from tasks where chatid=? and name=?", ChatID, task)
	return err

}

func (d *DbRepository) Mark(ChatID int64, task string) error {

	_, err := d.db.Exec("Update tasks set mark='✅' where chatid=? and name=?", ChatID, task)
	return err
}
