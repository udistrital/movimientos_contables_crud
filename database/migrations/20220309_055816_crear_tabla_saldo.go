package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaSaldo_20220309_055816 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaSaldo_20220309_055816{}
	m.Created = "20210304_182241"

	migration.Register("CrearTablaSaldo_20220309_055816", m)
}

// Run the migrations
func (m *CrearTablaSaldo_20220309_055816) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20220309_055816_crear_tabla_saldo_up.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}

}

// Reverse the migrations
func (m *CrearTablaSaldo_20220309_055816) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20220309_055816_crear_tabla_saldo_down.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}

}
