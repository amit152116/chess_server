package db

const UserTable = `CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY ,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    username VARCHAR(10) NOT NULL UNIQUE,
    email VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    avatar_url VARCHAR(255) ,
    bio TEXT);`

const TimeControlTable = `BEGIN TRANSACTION;
					CREATE TABLE IF NOT EXISTS timecontrols (
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
					COMMIT;`

const RatingsTable = `CREATE TABLE IF NOT EXISTS ratings (
    id SERIAL PRIMARY KEY,
    user_id SERIAL NOT NULL,
    bullet INT DEFAULT 1200,
    blitz INT DEFAULT 1200,
    rapid INT DEFAULT 1200,
    classical INT DEFAULT 1200,
    FOREIGN KEY (user_id) REFERENCES Users(id));`

const SessionTable = `CREATE TABLE IF NOT EXISTS sessions (
    id SERIAL PRIMARY KEY,
    user_id SERIAL UNIQUE REFERENCES Users(id),
    session_token UUID UNIQUE NOT NULL,
    expires_at TIMESTAMP NOT NULL);`

const (
	GamesTable_IDX1 = `CREATE INDEX games_idx on games(white_id,black_id)`
	GamesTable      = `CREATE TABLE IF NOT EXISTS games (
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
    FOREIGN KEY (winner_id) REFERENCES Users(id));`
)

const MovesTable = `CREATE TABLE IF NOT EXISTS moves (
    id SERIAL PRIMARY KEY,
    game_id UUID NOT NULL,
    player_id SERIAL NOT NULL,
    move_number INT NOT NULL,
    move VARCHAR(5) NOT NULL,
    move_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (game_id) REFERENCES Games(id) ON DELETE CASCADE,
    FOREIGN KEY (player_id) REFERENCES Users(id));`

const PairingQueue = `CREATE TABLE IF NOT EXISTS pairingqueue (
	id SERIAL PRIMARY KEY,
    user_id SERIAL NOT NULL,
    time_control_id INTEGER NOT NULL,
    requested_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (time_control_id) REFERENCES TimeControls(id),
    FOREIGN KEY (user_id) REFERENCES Users(id));`

const ChatTable = `CREATE TABLE IF NOT EXISTS chats (
    id SERIAL PRIMARY KEY,
    game_id UUID NOT NULL,
    sender_id SERIAL NOT NULL,
    message TEXT NOT NULL,
    sent_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (game_id) REFERENCES Games(id),
    FOREIGN KEY (sender_id) REFERENCES Users(id));`

const FriendsTable = `CREATE TABLE IF NOT EXISTS	friends (
    user_id SERIAL NOT NULL,
    friend_id SERIAL NOT NULL,
    status VARCHAR(10) NOT NULL CHECK (status IN ('pending', 'accepted')) DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, friend_id),
    FOREIGN KEY (user_id) REFERENCES Users(id),
    FOREIGN KEY (friend_id) REFERENCES Users(id));`

const OngoingGamesTable = `CREATE VIEW ongoinggames AS
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
		g.start_time DESC;`
