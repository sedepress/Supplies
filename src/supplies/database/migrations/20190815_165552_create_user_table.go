package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateUserTable_20190815_165552 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateUserTable_20190815_165552{}
	m.Created = "20190815_165552"

	migration.Register("CreateUserTable_20190815_165552", m)
}

// Run the migrations
func (m *CreateUserTable_20190815_165552) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("")
}

// Reverse the migrations
func (m *CreateUserTable_20190815_165552) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE user")
}
