package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaPaisCategoria_20210701_215525 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaPaisCategoria_20210701_215525{}
	m.Created = "20210701_215525"

	migration.Register("CrearTablaPaisCategoria_20210701_215525", m)
}

// Run the migrations
func (m *CrearTablaPaisCategoria_20210701_215525) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE convenio.pais_categoria(id serial NOT NULL, pais integer, CONSTRAINT pk_pais_categoria PRIMARY KEY (id));")
	m.SQL("ALTER TABLE convenio.pais_categoria OWNER TO postgres;")
}

// Reverse the migrations
func (m *CrearTablaPaisCategoria_20210701_215525) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS convenio.pais_categoria")
}
