package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaEstados_20210701_215511 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaEstados_20210701_215511{}
	m.Created = "20210701_215511"

	migration.Register("CrearTablaEstados_20210701_215511", m)
}

// Run the migrations
func (m *CrearTablaEstados_20210701_215511) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE convenio.estados(id serial NOT NULL, nombre varchar(50) NOT NULL, descripcion varchar, codigo_abreavicion varchar(20), activo boolean NOT NULL, numero_orden numeric(5,2), fecha_creacion date, fecha_modificacion date, CONSTRAINT pk_estados PRIMARY KEY (id));")
	m.SQL("TABLE convenio.estados OWNER TO postgres;)")
}

// Reverse the migrations
func (m *CrearTablaEstados_20210701_215511) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS convenio.estados")
}
