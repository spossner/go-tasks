package store

import (
	"sync"

	"possner.de/tasks/internal/store/sqlstore"
	"possner.de/tasks/internal/task"
)

type Store interface {
	Add(name string) *task.Task
	Get(id string) *task.Task
	Delete(id string) error
	Complete(id string) error
	List() []*task.Task
}

var theStore Store
var lock = &sync.Mutex{}

func GetStore() Store {
	if theStore == nil {
		lock.Lock()
		defer lock.Unlock()

		if theStore == nil {
			theStore = sqlstore.NewSQLStore()
		}
	}
	return theStore
}
