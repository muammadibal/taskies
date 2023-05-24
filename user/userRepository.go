package user

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetByID(id string) (*User, error)
	GetAll() ([]*User, error)
	Create(user *User) error
	Update(user *User) error
	Delete(id string) error
}

type userRepository struct {
	db          *gorm.DB
	redisClient *redis.Client
}

func NewUserRepository(db *gorm.DB, redis *redis.Client) UserRepository {
	return &userRepository{db, redis}
}

func (r *userRepository) GetByID(id string) (*User, error) {
	return nil, nil
}
func (r *userRepository) GetAll() ([]*User, error) {
	// cachedUser, err := r.GetFromCache(fmt.Sprintf("user:%d", id))
	// if err == nil {
	// 	return cachedUser, nil
	// }

	// if err := r.SetCache(fmt.Sprintf("user:%d", id), &user); err != nil {
	// 	// Log the error, but it's not critical to the request
	// }
	return nil, nil
}
func (r *userRepository) Create(user *User) error {
	return nil
}
func (r *userRepository) Update(user *User) error {
	return nil
}
func (r *userRepository) Delete(id string) error {
	return nil
}
func (r *userRepository) GetFromCache(key string) (*User, error) {
	result, err := r.redisClient.Get(context.Background(), key).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, fmt.Errorf("key %s not found in cache", key)
		}
		return nil, err
	}

	var user User
	if err := json.Unmarshal(result, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) SetCache(key string, user *User) error {
	jsonData, err := json.Marshal(user)
	if err != nil {
		return err
	}

	err = r.redisClient.Set(context.Background(), key, jsonData, time.Hour).Err()
	if err != nil {
		return err
	}

	return nil
}
