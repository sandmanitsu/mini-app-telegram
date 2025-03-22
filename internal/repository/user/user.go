package repository

import (
	"database/sql"
	"fmt"
	"log/slog"
	"mini-app-telegram/internal/domain"
	sl "mini-app-telegram/internal/logger"

	"github.com/Masterminds/squirrel"
)

type UserRepository struct {
	db     *sql.DB
	logger *slog.Logger
}

func NewUserRepository(db *sql.DB, logger *slog.Logger) *UserRepository {
	return &UserRepository{
		db:     db,
		logger: logger,
	}
}

func (u *UserRepository) CreateUser(user domain.User) error {
	const op = "repository.user.CreateUser"

	sql, args, err := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Insert("tg_user").
		Columns("tg_user_id", "username", "first_name", "last_name", "chat_id").
		Values(user.UserId, user.Username, user.FirstName, user.LastName, user.ChatId).
		ToSql()
	if err != nil {
		u.logger.Error(fmt.Sprintf("%s : building sql query", op), sl.Err(err))

		return err
	}

	_, err = u.db.Exec(sql, args...)
	if err != nil {
		u.logger.Error(fmt.Sprintf("%s: %s", op, sql), sl.Err(err))

		return err
	}

	return nil
}

func (u *UserRepository) GetUser(userId int64) (domain.User, error) {
	const op = "repository.user.GetUser"
	var user domain.User

	query, args, err := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Select("tg_user_id", "username", "first_name", "last_name", "chat_id").
		From("tg_user").
		Where("tg_user_id = ?", userId).
		ToSql()
	if err != nil {
		u.logger.Error(fmt.Sprintf("%s : building sql query", op), sl.Err(err))

		return user, err
	}

	err = u.db.QueryRow(query, args...).Scan(
		&user.UserId,
		&user.Username,
		&user.FirstName,
		&user.LastName,
		&user.ChatId,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			u.logger.Debug(fmt.Sprintf("%s : user not found! tg_user_id: %d", op, userId))
		} else {
			u.logger.Error(fmt.Sprintf("%s: %s", op, query), sl.Err(err))
		}

		return user, err
	}

	return user, nil
}
