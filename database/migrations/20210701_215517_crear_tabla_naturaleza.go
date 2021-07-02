package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaNaturaleza_20210701_215517 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaNaturaleza_20210701_215517{}
	m.Created = "20210701_215517"

	migration.Register("CrearTablaNaturaleza_20210701_215517", m)
}

// Run the migrations
func (m *CrearTablaNaturaleza_20210701_215517) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE convenio.naturaleza(id serial NOT NULL, nombre varchar(50) NOT NULL, descripcion varchar, codigo_abreviacion varchar(20), activo boolean NOT NULL, numero_orden numeric(5,2), CONSTRAINT pk_naturaleza PRIMARY KEY (id));")
	m.SQL("ALTER TABLE convenio.naturaleza OWNER TO postgres;")
}

// Reverse the migrations
func (m *CrearTablaNaturaleza_20210701_215517) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS convenio.naturaleza")
}
