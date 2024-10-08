package task

import (
	"bytes"
	"fmt"
	"text/tabwriter"
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/mergestat/timediff"
)

type Task struct {
	ID          string `gorm:"primaryKey"`
	Name        string
	CreatedAt   time.Time
	DueTo       time.Time
	CompletedAt time.Time
}

func NewTask(name string, dueDate time.Time) *Task {
	id, err := gonanoid.New()
	if err != nil {
		panic("error creating unique id")
	}
	return &Task{
		id,
		name,
		time.Now(),
		dueDate,
		time.Time{},
	}
}

func (t *Task) GetCreatedAt() string {
	if t.CreatedAt.IsZero() {
		return ""
	}
	return timediff.TimeDiff(t.CreatedAt)
}

func (t *Task) IsCompleted() bool {
	return !t.CompletedAt.IsZero()
}

func (t *Task) SetCompleted() {
	if t.IsCompleted() {
		fmt.Println("task " + t.ID + " already completed")
		return
	}
	t.CompletedAt = time.Now()
}

func (t *Task) GetCompleted() string {
	if t.CompletedAt.IsZero() {
		return ""
	}
	return timediff.TimeDiff(t.CompletedAt)
}

func (t *Task) GetDueTo() string {
	if t.DueTo.IsZero() {
		return ""
	}
	return timediff.TimeDiff(t.DueTo)
}

func (t *Task) String() string {
	buf := bytes.Buffer{}
	w := tabwriter.NewWriter(&buf, 0, 2, 4, ' ', 0)
	fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%v\n", t.ID, t.Name, t.GetCreatedAt(), t.GetDueTo(), t.GetCompleted())
	w.Flush()
	return buf.String()
}
