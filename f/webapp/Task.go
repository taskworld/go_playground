package main

import "fmt"

// taskDB structure
type taskDB struct {
	seq int
	m   map[int]*Task
}

func (tDB *taskDB) Get(id int) *Task {
	return tDB.m[id]
}

func (tDB *taskDB) GetAll() []*Task {
	if len(tDB.m) == 0 {
		return nil
	}

	result := make([]*Task, len(tDB.m))
	i := 0

	for _, v := range tDB.m {
		result[i] = v
		i++
	}
	return result
}

func (tDB *taskDB) Add(t *Task) (int, error) {
	tDB.seq++

	t.ID = tDB.seq

	//storing value
	tDB.m[t.ID] = t
	return t.ID, nil
}

func (tDB *taskDB) Update(t *Task) error {
	tDB.m[t.ID] = t
	return nil
}

func (tDB *taskDB) Delete(id int) {
	delete(tDB.m, id)
}

var db DB

func init() {

	db = &taskDB{
		m: make(map[int]*Task),
	}

	db.Add(&Task{
		ID:    1,
		Title: "Ask Bob to collect payment from the Kasikorn Bank",
	})

	db.Add(&Task{
		ID:    2,
		Title: "Design OnBoarding process for upcoming Thurs",
	})

	db.Add(&Task{
		ID:    3,
		Title: "Implement hook for everyone to use when Data is updated",
	})

	db.Add(&Task{
		ID:    4,
		Title: "Inform about new architecture of backend system",
	})
}

// Task structure
type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// ToString representation of Task
func (t *Task) ToString() string {
	return fmt.Sprintf("id: %d, title: %s", t.ID, t.Title)
}
