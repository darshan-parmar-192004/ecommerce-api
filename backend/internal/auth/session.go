package auth

import (
	"encoding/json"
	"time"

	"backend/internal/cache"
)

func StoreSession(r *cache.RedisService, token string, customer any) error {

	key := "session:" + token

	data, err := json.Marshal(customer)
	if err != nil {
		return err
	}

	return r.Client.Set(
		cache.Ctx,
		key,
		data,
		time.Hour,
	).Err()
}

func GetSession(r *cache.RedisService, token string) ([]byte, error) {

	key := "session:" + token

	return r.Client.Get(
		cache.Ctx,
		key,
	).Bytes()
}

func DeleteSession(r *cache.RedisService, token string) error {

	key := "session:" + token

	return r.Client.Del(
		cache.Ctx,
		key,
	).Err()
}
