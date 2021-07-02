package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaConvenios_20210701_215430 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaConvenios_20210701_215430{}
	m.Created = "20210701_215430"

	migration.Register("CrearTablaConvenios_20210701_215430", m)
}

// Run the migrations
func (m *CrearTablaConvenios_20210701_215430) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE convenio.convenio(id serial NOT NULL, descripcion varchar NOT NULL, responsable varchar NOT NULL, correo_responsable varchar NOT NULL, enlace varchar NOT NULL, id_pais_categoria integer, id_entidad integer, id_estados integer,CONSTRAINT pk_convenio PRIMARY KEY (id));")
	m.SQL("ALTER TABLE convenio.convenio OWNER TO postgres;")
}

// Reverse the migrations
func (m *CrearTablaConvenios_20210701_215430) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS convenio.convenio;")
}
