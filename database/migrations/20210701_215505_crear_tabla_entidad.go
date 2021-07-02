package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaEntidad_20210701_215505 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaEntidad_20210701_215505{}
	m.Created = "20210701_215505"

	migration.Register("CrearTablaEntidad_20210701_215505", m)
}

// Run the migrations
func (m *CrearTablaEntidad_20210701_215505) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE convenio.entidad(id serial NOT NULL, direccion varchar NOT NULL, correo varchar NOT NULL, telefono varchar NOT NULL, descripcion varchar, organizacion integer NOT NULL, id_naturaleza integer NOT NULL, CONSTRAINT pk_entidad PRIMARY KEY (id));")
	m.SQL("ALTER TABLE convenio.entidad OWNER TO postgres;")
}

// Reverse the migrations
func (m *CrearTablaEntidad_20210701_215505) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS convenio.entidad")
}
