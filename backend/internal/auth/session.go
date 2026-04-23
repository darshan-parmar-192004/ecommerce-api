package auth

import (
	"encoding/json"
	"fmt"

	"backend/internal/cache"
	"backend/internal/constants"
)

func StoreSession(r *cache.RedisService, token string, customer any) error {
	key := fmt.Sprintf("%s%s", constants.CacheKeySessionPrefix, token)

	data, err := json.Marshal(customer)
	if err != nil {
		return err
	}

	return r.Client.Set(
		cache.Ctx,
		key,
		data,
		constants.CacheSessionTTL,
	).Err()
}

func GetSession(r *cache.RedisService, token string) ([]byte, error) {
	key := fmt.Sprintf("%s%s", constants.CacheKeySessionPrefix, token)

	return r.Client.Get(
		cache.Ctx,
		key,
	).Bytes()
}

func DeleteSession(r *cache.RedisService, token string) error {
	key := fmt.Sprintf("%s%s", constants.CacheKeySessionPrefix, token)

	return r.Client.Del(
		cache.Ctx,
		key,
	).Err()
}
