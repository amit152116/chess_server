package db

import (
	"database/sql"
	"log"
)

func createUsersTable(db *sql.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS Users (
    id SERIAL PRIMARY KEY ,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    username VARCHAR(10) NOT NULL UNIQUE,
    email VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    avatar_url VARCHAR(255) ,
    bio TEXT);`)

	if err != nil {
		log.Panicln("Create USER Table: ", err)

	}

}

func createTimeControlTable(db *sql.DB) {
	_, err := db.Exec(`BEGIN TRANSACTION;
					CREATE TABLE IF NOT EXISTS TimeControls (
						id SMALLSERIAL PRIMARY KEY,
						name VARCHAR(10) UNIQUE NOT NULL
					);
					-- Only run this if the table creation was successful
					INSERT INTO TimeControls (name) 
					SELECT 'bullet' WHERE NOT EXISTS (SELECT 1 FROM TimeControls WHERE name = 'bullet')
					UNION ALL 
					SELECT 'blitz' WHERE NOT EXISTS (SELECT 1 FROM TimeControls WHERE name = 'blitz')
					UNION ALL 
					SELECT 'rapid' WHERE NOT EXISTS (SELECT 1 FROM TimeControls WHERE name = 'rapid')
					UNION ALL 
					SELECT 'classical' WHERE NOT EXISTS (SELECT 1 FROM TimeControls WHERE name = 'classical');
					COMMIT;`)

	if err != nil {
		log.Panicln("Create TIME CONTROL Table: ", err)
		return
	}
}

func createRatingsTable(db *sql.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS Ratings (
    id SERIAL PRIMARY KEY,
    user_id SERIAL NOT NULL,
    bullet INT DEFAULT 1200,
    blitz INT DEFAULT 1200,
    rapid INT DEFAULT 1200,
    classical INT DEFAULT 1200,
    FOREIGN KEY (user_id) REFERENCES Users(id));`)

	if err != nil {
		log.Panicln("Create RATING Table: ", err)
		return
	}
}

func createSessionsTable(db *sql.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS Sessions (
    id SERIAL PRIMARY KEY,
    user_id SERIAL UNIQUE REFERENCES Users(id),
    session_token UUID UNIQUE NOT NULL,
    expires_at TIMESTAMP NOT NULL);`)

	if err != nil {
		log.Panicln("Create SESSION Table: ", err)
		return
	}
}

func createGamesTable(db *sql.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS Games (
    id UUID PRIMARY KEY,
    white_id SERIAL NOT NULL,
    black_id SERIAL NOT NULL,
	time_control_id INTEGER NOT NULL,
	start_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    end_time TIMESTAMP DEFAULT NULL,
    status VARCHAR(10) NOT NULL CHECK (status IN ('in_progress', 'finished')) DEFAULT 'in_progress', 	
    game_result VARCHAR(10) CHECK (game_result IN ('resignation', 'timeout', 'checkmate', 'stalemate', 'draw')) DEFAULT NULL,
	winner_id SERIAL ,
	FOREIGN KEY (white_id) REFERENCES Users(id),
    FOREIGN KEY (black_id) REFERENCES Users(id),
	FOREIGN KEY (time_control_id) REFERENCES TimeControls(id),
    FOREIGN KEY (winner_id) REFERENCES Users(id));`)

	if err != nil {
		log.Panicln("Create GAME Table: ", err)
		return
	}
}

func createMovesTable(db *sql.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS Moves (
    id SERIAL PRIMARY KEY,
    game_id UUID NOT NULL,
    player_id SERIAL NOT NULL,
    move_number INT NOT NULL,
    move VARCHAR(5) NOT NULL,
    move_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (game_id) REFERENCES Games(id) ON DELETE CASCADE,
    FOREIGN KEY (player_id) REFERENCES Users(id));`)
	if err != nil {
		log.Panicln("Create MOVE Table: ", err)
		return
	}
}

func createMatchmakingQueue(db *sql.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS MatchmakingQueue (
	id SERIAL PRIMARY KEY,
    user_id SERIAL NOT NULL,
    time_control_id INTEGER NOT NULL,
    requested_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (time_control_id) REFERENCES TimeControls(id),
    FOREIGN KEY (user_id) REFERENCES Users(id));`)

	if err != nil {
		log.Panicln("Create MATCHMAKING QUEUE Table: ", err)
		return
	}
}

func createChatMessagesTable(db *sql.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS Chats (
    id SERIAL PRIMARY KEY,
    game_id UUID NOT NULL,
    sender_id SERIAL NOT NULL,
    message TEXT NOT NULL,
    sent_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (game_id) REFERENCES Games(id),
    FOREIGN KEY (sender_id) REFERENCES Users(id));`)
	if err != nil {
		log.Panicln("Create CHAT MESSAGE Table: ", err)
		return
	}
}
func createFriendsTable(db *sql.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS	Friends (
    user_id SERIAL NOT NULL,
    friend_id SERIAL NOT NULL,
    status VARCHAR(10) NOT NULL CHECK (status IN ('pending', 'accepted')) DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, friend_id),
    FOREIGN KEY (user_id) REFERENCES Users(id),
    FOREIGN KEY (friend_id) REFERENCES Users(id));`)
	if err != nil {
		log.Panicln("Create FRIENDS Table: ", err)
		return
	}
}
