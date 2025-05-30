package cache

import (
	"admin/pb"
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
)

type RedisAdminCache struct {
	client *redis.Client
}

func NewRedisAdminCache(client *redis.Client) *RedisAdminCache {
	return &RedisAdminCache{client: client}
}

func (c *RedisAdminCache) GetQueue() (*pb.QueueList, bool) {
	val, err := c.client.Get(context.Background(), "admin:queue").Result()
	if err != nil {
		return nil, false
	}
	var queue pb.QueueList
	if err := json.Unmarshal([]byte(val), &queue); err != nil {
		return nil, false
	}
	return &queue, true
}

func (c *RedisAdminCache) SetQueue(queue *pb.QueueList) error {
	data, err := json.Marshal(queue)
	if err != nil {
		return err
	}
	return c.client.Set(context.Background(), "admin:queue", data, 0).Err()
}

func (c *RedisAdminCache) GetSchedule(studentID string) (*pb.ScheduleResponse, bool) {
	val, err := c.client.Get(context.Background(), "admin:schedule:"+studentID).Result()
	if err != nil {
		return nil, false
	}
	var schedule pb.ScheduleResponse
	if err := json.Unmarshal([]byte(val), &schedule); err != nil {
		return nil, false
	}
	return &schedule, true
}

func (c *RedisAdminCache) SetSchedule(studentID string, schedule *pb.ScheduleResponse) error {
	data, err := json.Marshal(schedule)
	if err != nil {
		return err
	}
	return c.client.Set(context.Background(), "admin:schedule:"+studentID, data, 0).Err()
}
