package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-martini/martini"
)

func GetTasks(r *http.Request, db DB) (int, []byte) {
	b, _ := json.Marshal(db.GetAll())
	return http.StatusOK, b
}

func GetTask(r *http.Request, db DB, parms martini.Params) (int, []byte) {
	id, err := strconv.Atoi(parms["id"])
	t := db.Get(id)
	if err != nil || t == nil {
		// Invalid id, or does not exist
		return http.StatusNotFound, []byte{}
	}
	b, _ := json.Marshal(t)
	return http.StatusOK, b
}
