package db

import (
	"github.com/Amit152116Kumar/chess_server/models"
	"github.com/Amit152116Kumar/chess_server/utils"
	"github.com/google/uuid"
)

func (db *Database) GetPassword(email string) (string, error) {
	var password string

	err := db.conn.QueryRow(`SELECT password FROM Users WHERE email = $1;`,
		email).Scan(&password)

	return password, err
}

func (db *Database) GetUserDetails(id uuid.UUID) (*models.User, error) {
	var user models.User
	err := db.conn.QueryRow(`SELECT id, username, email,created_at, updated_at  FROM Users WHERE id = $1;`, id).Scan(
		&user.ID, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	return &user, err
}

func (db *Database) GetGameDetails(gameID int) (models.Game, error) {
	var game models.Game
	err := db.conn.QueryRow(`SELECT white_id, black_id, time_control_id, start_time FROM Games WHERE id = ?;`, gameID).Scan(
		&game.WhitePlayerID, &game.BlackPlayerID, &game.TimeControl, &game.StartTime, &game.EndTime, &game.Status, &game.WinnerID)
	return game, err

}

func (db *Database) GetOngoingGameID(whitePlayer, blackPlayer string) (int, error) {
	var id int
	err := db.conn.QueryRow(`SELECT id FROM Games WHERE white_id = ? AND black_id = ? ;`,
		whitePlayer, blackPlayer).Scan(&id)
	return id, err
}

func (db *Database) GetChatMessages(gameID int) ([]models.ChatMessage, error) {
	rows, err := db.conn.Query(`SELECT sender_id, message FROM Chats WHERE game_id = ? ORDER BY sent_at DESC ;`, gameID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []models.ChatMessage
	for rows.Next() {
		var message models.ChatMessage
		err = rows.Scan(&message.SenderID, &message.Message, &message.SentAt)
		if err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}
	return messages, nil
}

func (db *Database) GetFriends(userID int) ([]models.Friends, error) {
	rows, err := db.conn.Query(`SELECT friend_id, status FROM Friends WHERE user_id = ?;`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var friends []models.Friends
	for rows.Next() {
		var friend models.Friends
		err = rows.Scan(&friend.FriendID, &friend.Status)
		if err != nil {
			return nil, err
		}
		friends = append(friends, friend)
	}
	return friends, nil
}

func (db *Database) GetMoveHistory(gameID int) ([]models.Move, error) {
	rows, err := db.conn.Query(`SELECT player_id, move_number, move, move_time FROM Moves WHERE game_id = ? ORDER BY move_number;`, gameID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var moves []models.Move
	for rows.Next() {
		var move models.Move
		err = rows.Scan(&move.PlayerID, &move.MoveNumber, &move.Move, &move.MoveTime)
		if err != nil {
			return nil, err
		}
		moves = append(moves, move)
	}
	return moves, nil
}

func (db *Database) GetUserGameHistory(userID int) ([]models.Game, error) {
	rows, err := db.conn.Query(`SELECT id, white_id, black_id, time_control_id, start_time, end_time, status, winner_id 
									FROM Games WHERE (white_id = ? OR black_id = ?) AND status = ? ORDER BY start_time DESC;`,
		userID, userID, utils.GameStatusFinished.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var games []models.Game
	for rows.Next() {
		var game models.Game
		err = rows.Scan(&game.ID, &game.WhitePlayerID, &game.BlackPlayerID, &game.TimeControl, &game.StartTime, &game.EndTime, &game.Status, &game.WinnerID)
		if err != nil {
			return nil, err
		}
		games = append(games, game)
	}
	return games, nil
}

func (db *Database) GetUserStats(user models.User) (models.Stats, error) {
	var stats = models.Stats{UserDetails: user}

	err := db.conn.QueryRow(`SELECT bullet, blitz, rapid, classical FROM Ratings WHERE user_id = ?;`,
		user.ID).Scan(
		&stats.Bullet, &stats.Blitz, &stats.Rapid, &stats.Classical)

	if err != nil {
		return stats, err
	}
	err = db.conn.QueryRow(`SELECT COUNT(*) FROM Games WHERE (white_id = ? OR black_id = ?) AND status = ?;`,
		user.ID, user.ID, utils.GameStatusFinished.String()).Scan(&stats.TotalGames)

	if err != nil {
		return stats, err
	}

	err = db.conn.QueryRow(`SELECT COUNT(*) FROM Games WHERE winner_id = ?;`, user.ID).Scan(&stats.Wins)
	if err != nil {
		return stats, err
	}

	err = db.conn.QueryRow(`SELECT COUNT(*) FROM Games WHERE (white_id = ? OR black_id = ?) AND status = ? AND winner_id != ?;`,
		user.ID, user.ID, utils.GameStatusFinished.String(), user.ID).Scan(&stats.Losses)

	if err != nil {
		return stats, err
	}

	stats.Draws = stats.TotalGames - stats.Wins - stats.Losses

	return stats, err
}
