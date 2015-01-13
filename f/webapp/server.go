package main

import "github.com/go-martini/martini"

func main() {

	m := martini.Classic()

	//inject database
	m.MapTo(db, (*DB)(nil))

	m.Group("/api/tasks", func(r martini.Router) {
		r.Get("", GetTasks)
		r.Get("/:id", GetTask)
	})

	m.RunOnAddr(":8989")
}
