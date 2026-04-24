package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/reneruprecht/alertbridge/backend/internal/alert/application"
	"github.com/reneruprecht/alertbridge/backend/internal/alert/domain"
)

type AlertCache struct {
	client *redis.Client
}

func NewAlertCache(client *redis.Client) *AlertCache {
	return &AlertCache{client: client}
}

func (r *AlertCache) ListAlerts(ctx context.Context) ([]application.AlertCacheDto, error) {

	var cursor uint64
	var alerts []application.AlertCacheDto

	for {
		keys, nextCursor, err := r.client.Scan(ctx, cursor, "alert:*", 100).Result()
		if err != nil {
			return nil, err
		}
		cursor = nextCursor

		pipe := r.client.Pipeline()
		cmds := make([]*redis.StringCmd, len(keys))
		for i, k := range keys {
			cmds[i] = pipe.Get(ctx, k)
		}
		pipe.Exec(ctx)

		for _, cmd := range cmds {
			val, _ := cmd.Result()
			var a alertDto
			err := json.Unmarshal([]byte(val), &a)
			if err != nil {
				return nil, err
			}

			alert := toCacheDto(a)
			alerts = append(alerts, alert)
		}

		if cursor == 0 {
			break
		}
	}

	return alerts, nil

}

func (r *AlertCache) Save(ctx context.Context, alert domain.Alert) error {

	alertDto := toDto(alert)

	json, err := json.Marshal(alertDto)

	if err != nil {
		return err
	}

	_, redisErr := r.client.Set(ctx, "alert:"+alertDto.Fingerprint, string(json), 15*time.Minute).Result()

	if redisErr != nil {
		return redisErr
	}

	return nil

}

func (r *AlertCache) DeleteByKey(ctx context.Context, alert domain.Alert) error {

	key := fmt.Sprintf("alert:%s", alert.Fingerprint)

	err := r.client.Del(ctx, key).Err()

	if err != nil {
		return err
	}

	return nil

}
