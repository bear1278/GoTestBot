package telegram

import (
	"time"
)

type Task struct{
	Name string
	Date time.Time 
	Mark string
}

func  newTask(n string, t time.Time, mark string) *Task {
	return &Task{Name: n, Date: t,Mark: mark}
}

func (t* Task) GetDataOfTask() string{
	return t.Name+" "+t.Date.Format("2006-01-02")+" "+t.Mark
}