package main

import (
	"github.com/muhwyndhamhp/spill/db"
)

func main() {
	db := db.GetDB()

	db.Debug()
}
