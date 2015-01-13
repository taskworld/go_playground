package main

// The DB interface defines methods to manipulate the albums.
type DB interface {
	Get(id int) *Task
	GetAll() []*Task
	Add(t *Task) (int, error)
	Update(t *Task) error
	Delete(id int)
}
