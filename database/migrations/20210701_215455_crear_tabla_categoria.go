package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaCategoria_20210701_215455 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaCategoria_20210701_215455{}
	m.Created = "20210701_215455"

	migration.Register("CrearTablaCategoria_20210701_215455", m)
}

// Run the migrations
func (m *CrearTablaCategoria_20210701_215455) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE convenio.categoria(id serial NOT NULL, nombre varchar(50) NOT NULL, descripcion varchar, codigo_abreavicion varchar(20), activo boolean NOT NULL, numero_orden numeric(5,2), id_pais_categoria integer, CONSTRAINT pk_categoria PRIMARY KEY (id));")
	m.SQL("ALTER TABLE convenio.categoria OWNER TO postgres;")
}

// Reverse the migrations
func (m *CrearTablaCategoria_20210701_215455) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS convenio.categoria;")
}
