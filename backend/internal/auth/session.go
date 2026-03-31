package auth

import (
	"encoding/json"
	"time"

	"backend/internal/services"
)

func StoreSession(r *services.RedisService, token string, customer any) error {
	if r.Client == nil {
		return nil
	}

	key := "session:" + token

	data, err := json.Marshal(customer)
	if err != nil {
		return err
	}

	return r.Client.Set(
		services.Ctx,
		key,
		data,
		time.Hour,
	).Err()
}

func GetSession(r *services.RedisService, token string) ([]byte, error) {
	if r.Client == nil {
		return nil, nil
	}

	key := "session:" + token

	return r.Client.Get(
		services.Ctx,
		key,
	).Bytes()
}

func DeleteSession(r *services.RedisService, token string) error {
	if r.Client == nil {
		return nil
	}

	key := "session:" + token

	return r.Client.Del(
		services.Ctx,
		key,
	).Err()
}
