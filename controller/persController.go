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

func GetPersonsExtra() []PersonExtra {
	results := ExecuteSQL("select p.PERS_NO, p.FIRSTNAME, p.LASTNAME, p.USERNAME, f.FUNCTION_NO, f.FUNCTION_NAME, ac.CITY_NO, ac.CITY_NAME from pers p inner join atemschutzpflegestelle_cities ac on p.city_no = ac.CITY_NO inner join function f on p.FUNCTION_NO = f.FUNCTION_NO where p.IS_ACTIVE=1 order by p.LASTNAME")
	persons := []PersonExtra{}
	for results.Next() {
		var pers PersonExtra
		results.Scan(&pers.PersNo, &pers.Firstname, &pers.Lastname, &pers.Username, &pers.FunctionNo, &pers.FunctionName, &pers.CityNo, &pers.CityName)
		persons = append(persons, pers)
	}
	return persons
}

func GetFunctions() []Function {
	results := ExecuteSQL("select FUNCTION_NO, FUNCTION_NAME from function order by FUNCTION_NO")
	functions := []Function{}
	for results.Next() {
		var function Function
		results.Scan(&function.FunctionNo, &function.FunctionName)
		functions = append(functions, function)
	}
	return functions
}
