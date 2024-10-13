package db

import (
	"database/sql"
	"github.com/Amit152116Kumar/chess_server/config"
	_ "github.com/lib/pq"
	"log"
	"sync"
)

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

func (db *Database) CreateAllTables() {
	createUsersTable(db.conn)
	createTimeControlTable(db.conn)
	createRatingsTable(db.conn)
	createSessionsTable(db.conn)
	createGamesTable(db.conn)
	createMovesTable(db.conn)
	createMatchmakingQueue(db.conn)
	createChatMessagesTable(db.conn)
	createFriendsTable(db.conn)

	// Views creation
	createOngoingGamesView(db.conn)

	log.Println("All tables created successfully")
}
