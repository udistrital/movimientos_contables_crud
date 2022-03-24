package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearEsquemaMovimientosContables_20210210_124631 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearEsquemaMovimientosContables_20210210_124631{}
	m.Created = "20210210_124631"

	migration.Register("CrearEsquemaMovimientosContables_20210210_124631", m)
}

// Run the migrations
func (m *CrearEsquemaMovimientosContables_20210210_124631) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20210210_124631_crear_esquema_movimientos_contables_up.sql")

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
func (m *CrearEsquemaMovimientosContables_20210210_124631) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20210210_124631_crear_esquema_movimientos_contables_down.sql")

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
