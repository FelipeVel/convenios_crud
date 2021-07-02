package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearSchema_20210701_220752 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearSchema_20210701_220752{}
	m.Created = "20210701_220752"

	migration.Register("CrearSchema_20210701_220752", m)
}

// Run the migrations
func (m *CrearSchema_20210701_220752) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE SCHEMA convenio;")
	m.SQL("ALTER SCHEMA convenio OWNER TO postgres;")
	m.SQL("SET search_path TO pg_catalog,public,convenio;")
}

// Reverse the migrations
func (m *CrearSchema_20210701_220752) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP SCHEMA IF EXISTS convenio")
}
