package db

type DbRepository interface {
	CheckTask(ChatID int64, task string) (bool, error)
	List(ChatID int64) (string, error)
	Add(ChatID int64, task string) error
	Delete(ChatID int64, task string) error
	Mark(ChatID int64, task string) error
	CheckAlreadyMark(ChatID int64, task string) (bool, error)
}