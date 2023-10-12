package postgres

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/muathendirangu/hex/internal/core/domain"
)

type MessagePostgresRepository struct {
	db *gorm.DB
}

func NewMessagePostgresRepository() *MessagePostgresRepository {
	host := "localhost"
	port := "5432"
	user := "postgres"
	password := "pass1234"
	dbname := "postgres"

	conn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := gorm.Open("postgres", conn)
	if err != nil {
		panic("failed to connect to database")
	}
	db.AutoMigrate(&domain.Message{})
	return &MessagePostgresRepository{db: db}
}

func (repo *MessagePostgresRepository) SaveMessage(message domain.Message) error {
	req := repo.db.Create(&message)
	if req.RowsAffected == 0 {
		return fmt.Errorf("message not saved :%v", req.Error)
	}
	return nil
}

func (repo *MessagePostgresRepository) ReadMessage(id string) (*domain.Message, error) {
	var message domain.Message
	req := repo.db.First(&message, "id = ?", id)
	if req.Error != nil {
		return nil, req.Error

	}
	return &message, nil
}

func (repo *MessagePostgresRepository) ReadMessages() ([]*domain.Message, error) {
	var messages []*domain.Message
	req := repo.db.Find(&messages)
	if req.Error != nil {
		return nil, fmt.Errorf("failed to fetch messages: %w", req.Error)
	}
	return messages, nil
}
