package redis

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"
	"github.com/muathendirangu/hex/internal/core/domain"
)

type MessageRedisRepository struct {
	db *redis.Client
}

func NewMessageRedisRepository(host string) *MessageRedisRepository {
	db := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: "",
		DB:       0,
	})
	return &MessageRedisRepository{
		db: db,
	}
}

func (r *MessageRedisRepository) SaveMessage(message domain.Message) error {
	json, err := json.Marshal(message)
	if err != nil {
		return err
	}
	r.db.HSet(context.Background(), "messages", message.ID, json)
	return nil
}

func (r *MessageRedisRepository) ReadMessage(id string) (*domain.Message, error) {
	value, err := r.db.HGet(context.Background(), "messages", id).Result()
	if err != nil {
		return nil, err
	}
	message := &domain.Message{}
	err = json.Unmarshal([]byte(value), &message)
	if err != nil {
		return nil, err
	}
	return message, nil
}

func (r *MessageRedisRepository) ReadMessages() ([]*domain.Message, error) {
	var messages []*domain.Message
	values, err := r.db.HGetAll(context.Background(), "messages").Result()
	if err != nil {
		return nil, err
	}
	for _, value := range values {
		message := &domain.Message{}
		err = json.Unmarshal([]byte(value), &message)
		if err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}
	return messages, nil
}
