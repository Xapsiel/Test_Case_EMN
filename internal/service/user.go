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
		return fmt.Errorf("пропущены параметры для smtp")
	}
	token := uuid.New().String()
	subject := "Confirm Your Email"
	htmlBody := fmt.Sprintf("<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n<meta charset=\"UTF-8\">\n<meta http-equiv=\"X-UA-Compatible\" content=\"IE=edge\">\n<meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">\n<title>Подтвердите свой почтовый адрес</title>\n</head>\n<body style=\"font-family: Arial, sans-serif;\">\n<div style=\"max-width: 600px; margin: 0 auto; padding: 20px;\">\n  <h1 style=\"text-align: center; color: #333;\">Подтвердите свой почтовый адрес</h1>\n  <p style=\"margin-bottom: 20px;\">Если вы не создавали аккаунт на сайте SuperMegaWorker.com\n                    проигнорируйте данное сообщение или удалите его. Код подтверждения:</p>\n  <p  style=\"display: inline-block; padding: 10px 20px; background-color: #007bff; color: #fff; text-decoration: none; border-radius: 5px;\">%s</p>\n</div>\n</body>\n</html>", token)
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
	err = u.repo.UpdateToken(token, email)
	if err != nil {
		return fmt.Errorf("Ошибка обновления данных о токене")
	}
	return nil
}
