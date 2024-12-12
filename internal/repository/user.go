package repository

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"mobileTest_Case/internal/models"
	"time"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}
func (p *UserPostgres) Register(nickname, name, email string) error {
	var existingUser models.User
	query := "SELECT * FROM users WHERE email = $2"
	err := p.db.Get(&existingUser, query, email)

	if err == nil {
		return fmt.Errorf("Такой пользователь уже существует")
	}

	id := uuid.New().String()
	insertQuery := "INSERT INTO users (id,nickname, name, email,verified) VALUES ($1, $2, $3,$4,$5) "
	_ = p.db.QueryRow(insertQuery, id, nickname, name, email, false)

	return nil
}
func (p *UserPostgres) UpdateToken(token, email string) error {
	query := "INSERT INTO email_tokens (email, token, expires_at,created_at) VALUES ($1, $2, current_timestamp+interval '5 hours', current_timestamp) on conflict (email) do update  set token=$3, expires_at=current_timestamp+interval '5 hours', created_at=current_timestamp"
	_, err := p.db.Exec(query, email, token, token)
	if err != nil {
		return fmt.Errorf("Ошибка создания или обновления токена: %v", err)
	}
	return nil
}
func (p *UserPostgres) VerifyToken(token string) error {
	var email string
	var expires_at time.Time
	query := "SELECT email,expires_at from email_tokens where token=$1"
	err := p.db.QueryRow(query, token).Scan(&email, &expires_at)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("токен не найден или уже использован")
		}
		return fmt.Errorf("ошибка проверки токена: %v", err)
	}
	if time.Now().After(expires_at) {
		return fmt.Errorf("Время действия токена истекло")
	}
	query = "UPDATE users SET verified = true WHERE email = $1"
	_, err = p.db.Exec(query, email)
	if err != nil {
		return fmt.Errorf("ошибка обновления статуса пользователя: %v", err)
	}
	query = "DELETE FROM email_tokens WHERE token = $1"
	_, err = p.db.Exec(query, token)
	if err != nil {
		return fmt.Errorf("ошибка удаления токена: %v", err)
	}

	return nil
}

//func (p *UserPostgres) Verify(token string)
///
