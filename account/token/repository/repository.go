package repository

import (
	"github.com/go-redis/redis/v8"
	"github.com/jllanes-ss/avisos/account/domain"
	_tokenRedisRepository "github.com/jllanes-ss/avisos/account/token/repository/redis"
)

// NewTokenRepository is a factory for initializing User Repositories
func GetRepository(redisClient *redis.Client) domain.TokenRepository {
	return _tokenRedisRepository.NewTokenRepository(redisClient)
}
