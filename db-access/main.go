package main

import (
	"database/sql"

	db "./db"
	model "./model"
)

func main() {

	x := db.Begin()
	var err error
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()
	// retrieve single person
	person := model.Person{}
	if err := tx.Getx(&person, model.QueryPersonByID(personID)); err != nil {
		return err
	}
	person.Lastname = "Doe"
	// update the person
	if err := tx.Update(&person); err != nil {
		return err
	}
	index := 0
	count := 50
	// retrieve multiple paged persons
	persons := []model.Person{}
	if err := ctx.tx.Selectx(&persons, model.QueryPersons(user), db.Limit(index, count)); err == sql.ErrNoRows {
	} else if err == nil {
	} else {
		return err
	}
	// count number of results
	total, err := ctx.tx.Countx(model.QueryPersons())
	if err != nil {
		return err
	}
}
