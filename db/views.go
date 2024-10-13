package db

import (
	"database/sql"
	"log"
)

func createOngoingGamesView(db *sql.DB) {

	if _, err := db.Exec(`DROP View IF EXISTS OngoingGames;`); err != nil {
		log.Println("Drop ONGOING GAMES VIEW: ", err)
	}
	if _, err := db.Exec(`CREATE VIEW OngoingGames AS
	SELECT 
		g.id AS game_id,
		g.start_time,
		p1.username AS player1,
		p2.username AS player2,
		tc.name AS time_control
	FROM 
		Games g
	JOIN 
		Users p1 ON g.white_id = p1.id
	JOIN 
		Users p2 ON g.black_id = p2.id
	JOIN 
		TimeControls tc ON g.time_control_id = tc.id
	WHERE 
		g.status = 'in_progress'
	ORDER BY 
		g.start_time DESC;`); err != nil {
		log.Panicln("Create ONGOING GAMES VIEW: ", err)
		return
	}
}
