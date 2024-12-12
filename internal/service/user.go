package service

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/k3a/html2text"
	"gopkg.in/gomail.v2"
	"mobileTest_Case/internal/repository"

	"os"
	"strconv"
)

type UserService struct {
	repo repository.User
}

func (u *UserService) VerifyEmail(token string) error {
	return u.repo.VerifyToken(token)
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (u *UserService) Register(nickname, name, email string) error {
	err := u.repo.Register(nickname, name, email)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserService) SendToken(email string) error {
	// Получение переменных окружения для настройки SMTP
	from := os.Getenv("EMAIL_FROM")
	smtpPass := os.Getenv("SMTP_PASS")
	smtpUser := os.Getenv("SMTP_USER")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	// Проверка обязательных параметров
	if from == "" || smtpPass == "" || smtpUser == "" || smtpHost == "" || smtpPort == "" {
		return fmt.Errorf("missing required SMTP configuration")
	}
	token := uuid.New().String()
	subject := "Confirm Your Email"
	htmlBody := fmt.Sprintf("<h1>Welcome!</h1><p>Your token</p><p>%s</p>", token)
	// Создание сообщения
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", email)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", htmlBody)
	m.AddAlternative("text/plain", html2text.HTML2Text(htmlBody))

	// Установка SMTP клиента
	port, err := strconv.Atoi(smtpPort)
	if err != nil {
		return fmt.Errorf("invalid SMTP port: %v", err)
	}
	dialer := gomail.NewDialer(smtpHost, port, smtpUser, smtpPass)

	// Отправка сообщения
	if err := dialer.DialAndSend(m); err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}
	u.repo.UpdateToken(token, email)
	return nil
}
