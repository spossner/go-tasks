package sqlstore

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"possner.de/tasks/internal/task"
)

type SQLStore struct {
	db *gorm.DB
}

func NewSQLStore() *SQLStore {
	db, err := gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(task.Task{})
	return &SQLStore{db}
}

func (s *SQLStore) Add(name string) *task.Task {
	t := task.NewTask(name)
	s.db.Create(t)
	return t
}

func (s *SQLStore) Get(id string) *task.Task {
	var t task.Task
	s.db.First(&t, "id = ?", id)
	return &t
}

func (s *SQLStore) Delete(id string) error {
	s.db.Delete(&task.Task{}, "id = ?", id)
	return nil
}

func (s *SQLStore) Complete(id string) error {
	t := s.Get(id)
	t.SetCompleted()
	fmt.Println(t)
	s.db.Save(t)
	return nil
}

func (s *SQLStore) List() []*task.Task {
	var tasks []*task.Task
	s.db.Find(&tasks)
	return tasks
}
