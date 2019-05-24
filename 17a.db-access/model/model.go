package model

import (
	"database/sql/driver"
	"time"

	db "github.com/arunpa0206/gotrainingv1/db-access/db"
	"github.com/jmoiron/sqlx"
)

type Gender string

var (
	GenderMale   Gender = "male"
	GenderFemale Gender = "female"
)

func (u *Gender) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	b := value.([]byte)
	*u = Gender(b)
	return nil
}
func (u Gender) Value() (driver.Value, error) {
	return string(u), nil
}

type Person struct {
	PersonID     utils.UUID `db:"person_id"`
	FirstName    string     `db:"first_name"`
	LastName     string     `db:"last_name"`
	Active       Bool       `db:"active"`
	Gender       Gender     `db:"gender"`
	ModifiedDate time.Time  `db:"modified_date"`
}

var (
	queryPersons      db.Query = "SELECT person_id, first_name, last_name, gender, active, modified_date FROM persons"
	queryPersonByID   db.Query = "SELECT person_id, first_name, last_name, gender, active, modified_date FROM persons WHERE person_id=:person_id"
	queryPersonInsert db.Query = "INSERT INTO persons (person_id, first_name, last_name, gender, active, modified_date) VALUES (:person_id, :first_name, :last_name, :gender, :active, :modified_date)"
	queryPersonUpdate db.Query = "UPDATE persons SET first_name=:first_name, last_name=:last_name, gender=:gender, modified_date=:modified_date, active=:active WHERE person_id=:person_id"
)

func QueryPersons(offset, count int) db.Queryx {
	return db.Queryx{
		Query:  queryPersons,
		Params: []interface{}{},
	}
}
func QueryPersonByID(personID utils.UUID) db.Queryx {
	return db.Queryx{
		Query: queryPersonByID,
		Params: []interface{}{
			personID,
		},
	}
}
func NewPerson() *Person {
	return &Person{PersonID: utils.NewUUID(), ModifiedDate: time.Now()}
}
func (s *Person) Insert(tx *sqlx.Tx) error {
	_, err := tx.NamedExec(string(queryPersonInsert), s)
	return err
}
func (s *Person) Update(tx *sqlx.Tx) error {
	s.ModifiedDate = time.Now()
	_, err := tx.NamedExec(string(queryPersonUpdate), s)
	return err
}
func (s *Person) Delete(tx *sqlx.Tx) error {
	s.Active = false
	return s.Update(tx)
}
