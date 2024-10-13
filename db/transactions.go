package db

import (
	"fmt"
	"github.com/Amit152116Kumar/chess_server/models"
	"github.com/Amit152116Kumar/chess_server/utils"
	"time"
)

func (db *Database) AddUser(user *models.RegisterUserPayload) error {
	fmt.Println(user)
	_, err := db.conn.Exec(`INSERT INTO Users (first_name, last_name, username, email, password) VALUES ($1, $2, $3, $4, $5);`,
		user.FirstName, user.LastName, user.Username, user.Email, user.Password)
	return err
}

func (db *Database) UpdateAvatarURL(username, avatarURL string) error {
	_, err := db.conn.Exec(`UPDATE Users SET avatar_url = ? updated_at = ? WHERE username = ?;`, avatarURL, time.Now(), username)
	return err
}

func (db *Database) UpdateBio(username, bio string) error {
	_, err := db.conn.Exec(`UPDATE Users SET bio = ?  updated_at = ? WHERE username = ?;`,
		bio, time.Now(), username)
	return err
}

func (db *Database) UpdateUsername(oldUsername, newUsername string) error {
	_, err := db.conn.Exec(`UPDATE Users SET username = ? updated_at = ? WHERE username = ?;`,
		newUsername, time.Now(), oldUsername)
	return err
}

func (db *Database) UpdateEmail(username, email string) error {
	_, err := db.conn.Exec(`UPDATE Users SET email = ? updated_at = ? WHERE username = ?;`,
		email, time.Now(), username)
	return err
}

func (db *Database) UpdatePassword(username, password string) error {
	_, err := db.conn.Exec(`UPDATE Users SET password = ? updated_at = ? WHERE username = ?;`,
		password, time.Now(), username)
	return err
}

func (db *Database) UpdateRating(username, timeControl utils.TimeControl, rating int) error {
	_, err := db.conn.Exec(`UPDATE Ratings SET ? = ? WHERE user_id = (SELECT id FROM Users WHERE username = ?);`,
		timeControl.String(), rating, username)
	return err
}

func (db *Database) AddGame(whitePlayer, blackPlayer string, timeControl utils.TimeControl) error {
	_, err := db.conn.Exec(`INSERT INTO Games (
                   white_id, 
                   black_id, 
                   time_control_id
                   ) VALUES (
                             (SELECT id FROM Users WHERE username = ?),
                             (SELECT id FROM Users WHERE username = ?),
                             ?);`,
		whitePlayer, blackPlayer, timeControl.String())
	return err
}

func (db *Database) DeleteGame(gameID int) error {
	_, err := db.conn.Exec(`DELETE FROM Games WHERE id = ?;`, gameID)
	return err
}

func (db *Database) UpdateGameStatus(gameID int, result utils.GameResult, winner *string) error {

	if winner != nil {
		_, err := db.conn.Exec(`UPDATE Games SET status = ?, winner_id = (SELECT id FROM Users WHERE username = ?) WHERE id = ?;`,
			result.String(), *winner, gameID)
		return err
	} else {
		_, err := db.conn.Exec(`UPDATE Games SET status = ? WHERE id = ?;`, result.String(), gameID)
		return err
	}
}

func (db *Database) AddMove(gameID, playerID, moveNumber int, move string) error {
	_, err := db.conn.Exec(`INSERT INTO Moves (game_id, player_id, move_number, move) VALUES (?, ?, ?, ?);`,
		gameID, playerID, moveNumber, move)
	return err
}

func (db *Database) AddChatMessage(gameID int, sender, message string) error {
	_, err := db.conn.Exec(`INSERT INTO Chats (game_id, sender_id, message) VALUES 
                                                    (?, (SELECT id FROM Users WHERE username = ?), ?);`,
		gameID, sender, message)
	return err
}

func (db *Database) AddFriendRequest(sender, receiver string) error {
	_, err := db.conn.Exec(`INSERT INTO FriendRequests (user_id, friend_id) VALUES ((SELECT id FROM Users WHERE username = ?),
                                                        (SELECT id FROM Users WHERE username = ?));`,
		sender, receiver)
	return err
}

func (db *Database) UpdateFriendStatus(sender, receiver string, status utils.FriendStatus) error {
	_, err := db.conn.Exec(`UPDATE FriendRequests SET status = ? WHERE user_id = 
                                           (SELECT id FROM Users WHERE username = ?) 
                                       AND friend_id = (SELECT id FROM Users WHERE username = ?);`,
		status.String(), sender, receiver)
	return err
}

func (db *Database) DeleteFriendRequest(sender, receiver string) error {
	_, err := db.conn.Exec(`DELETE FROM FriendRequests WHERE user_id = 
										   (SELECT id FROM Users WHERE username = ?) 
									   AND friend_id = (SELECT id FROM Users WHERE username = ?);`,
		sender, receiver)
	return err
}
