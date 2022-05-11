package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func ConnectDatabase() error {
	db, err := sql.Open("sqlite3", "./compendium.db")
	if err != nil {
		return err
	}
	DB = db
	return nil
}

func CreateTable(db *sql.DB) {
	tables := `CREATE TABLE IF NOT EXISTS user (id TEXT PRIMARY KEY, username TEXT NOT NULL, password TEXT NOT NULL);
	CREATE TABLE IF NOT EXISTS campaign (id INTEGER PRIMARY KEY AUTOINCREMENT, title TEXT, description TEXT, image TEXT);
	CREATE TABLE IF NOT EXISTS world (id INTEGER PRIMARY KEY AUTOINCREMENT, title TEXT, content TEXT, image TEXT, compaign_id, FOREIGN KEY(compaign_id) REFERENCES compaign(id));
	CREATE TABLE IF NOT EXISTS characters (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT, content TEXT, image TEXT, compaign_id, FOREIGN KEY(compaign_id) REFERENCES compaign(id));
	CREATE TABLE IF NOT EXISTS lore (id INTEGER PRIMARY KEY AUTOINCREMENT, title TEXT, image TEXT, content TEXT, compaign_id, FOREIGN KEY(compaign_id) REFERENCES compaign(id)) `
	_, err := db.Exec(tables)
	if err != nil {
		panic(err)
	}
}
