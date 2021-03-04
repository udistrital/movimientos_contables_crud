package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type ModificarTablaTransaccion_20210304_182241 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &ModificarTablaTransaccion_20210304_182241{}
	m.Created = "20210304_182241"

	migration.Register("ModificarTablaTransaccion_20210304_182241", m)
}

// Run the migrations
func (m *ModificarTablaTransaccion_20210304_182241) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20210304_182241_modificar_tabla_transaccion_up.sql")

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
func (m *ModificarTablaTransaccion_20210304_182241) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20210304_182241_modificar_tabla_transaccion_down.sql")

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
