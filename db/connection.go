package db

import (
	"database/sql"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/amit152116/chess_server/config"
	_ "github.com/lib/pq"
)

const DatabaseVersion = 1

type Database struct {
	conn *sql.DB
}

var (
	Instance *Database
	once     sync.Once
)

func SetupDBConnection() *Database {
	once.Do(func() {
		dbConn, err := sql.Open("postgres", config.Cfg.GetConnectionString())
		if err != nil {
			log.Panicln("SetupDBConnection: ", err)
		}

		Instance = &Database{conn: dbConn}

		Instance.Ping()
	})
	log.Println("Database connection established")
	return Instance
}

func (db *Database) Ping() {
	err := db.conn.Ping()
	if err != nil {
		log.Panicln("Ping: ", err)
	}
	log.Println("Ping successful")
}

func (db *Database) Close() {
	if err := db.conn.Close(); err != nil {
		log.Panicln("Close: ", err)
	}
	log.Println("database Closed")
}

func (db *Database) GetDatabaseVersion() int {
	var currentVersion string
	err := db.conn.QueryRow(`SELECT value FROM metadata WHERE key = 'schema_version'`).Scan(&currentVersion)
	if err != nil {
		return 0
	}

	version, err := strconv.Atoi(currentVersion)
	if err != nil {
		return 0
	}
	return version
}

func (db *Database) SetConnectionSettings() {
	db.conn.SetMaxIdleConns(5)
	db.conn.SetMaxOpenConns(20)
	db.conn.SetConnMaxIdleTime(time.Duration(time.Minute * 1))
	db.conn.SetConnMaxLifetime(time.Duration(time.Minute * 10))
}

func (db *Database) CreateAllTables() {
	oldVersion := db.GetDatabaseVersion()
	if oldVersion == 0 {
		db.onCreateDatabase()
	} else if oldVersion == DatabaseVersion {
		return
	} else if oldVersion < DatabaseVersion {
		db.onUpgradeDatabase(oldVersion, DatabaseVersion)
	} else {
		log.Panicf("database schema version (%d) is higher than expected (%d)", oldVersion, DatabaseVersion)
	}
}

func (db *Database) onCreateDatabase() {
	result, err := db.conn.Exec(`CREATE TABLE IF NOT EXISTS metadata (
    key TEXT PRIMARY KEY, value TEXT NOT NULL);`)
	if err != nil {
		return
	}
	log.Println(result)
	db.updateDBVersion()
}

func (db *Database) updateDBVersion() {
	result, err := db.conn.Exec(`INSERT INTO metadata (key, value) VALUES ('$1', '$2');
		`, "schema_version", DatabaseVersion)
	if err != nil {
		return
	}
	log.Println(result)
}

func (db *Database) onUpgradeDatabase(oldVersion int, newVersion int) {
	switch newVersion {
	case 1:
	default:
	}
	db.updateDBVersion()
}

func (db *Database) onDropDatabase() {
}
