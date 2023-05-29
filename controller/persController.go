package controller

import (
	. "ffAPI/models"

	_ "github.com/go-sql-driver/mysql"
)

func GetPersons() []Person {
	results := ExecuteSQL("select PERS_NO, FIRSTNAME, LASTNAME from pers where IS_ACTIVE=1 order by LASTNAME")
	persons := []Person{}
	for results.Next() {
		var pers Person
		results.Scan(&pers.PersNo, &pers.Firstname, &pers.Lastname)
		persons = append(persons, pers)
	}
	return persons
}
