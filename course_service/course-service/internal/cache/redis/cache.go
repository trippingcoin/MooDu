package cache

import (
	"context"
	"cs/course-service/internal/model"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type AssignmentCache interface {
	Get(id string) (*model.Assignment, error)
	Set(assignment *model.Assignment) error
	Delete(id string) error
	List() ([]*model.Assignment, error)
	SetList([]*model.Assignment) error
}

type AssignmentRedisCache struct {
	rdb *redis.Client
	ctx context.Context
}

func NewAssignmentRedisCache(rdb *redis.Client) *AssignmentRedisCache {
	return &AssignmentRedisCache{
		rdb: rdb,
		ctx: context.Background(),
	}
}

func (c *AssignmentRedisCache) Get(id string) (*model.Assignment, error) {
	key := fmt.Sprintf("assignment:%s", id)
	val, err := c.rdb.Get(c.ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var assignment model.Assignment
	if err := json.Unmarshal([]byte(val), &assignment); err != nil {
		return nil, err
	}
	return &assignment, nil
}

func (c *AssignmentRedisCache) Set(assignment *model.Assignment) error {
	key := fmt.Sprintf("assignment:%s", assignment.ID)
	data, err := json.Marshal(assignment)
	if err != nil {
		return err
	}
	return c.rdb.Set(c.ctx, key, data, 10*time.Minute).Err()
}

func (c *AssignmentRedisCache) Delete(id string) error {
	key := fmt.Sprintf("assignment:%s", id)
	return c.rdb.Del(c.ctx, key).Err()
}

func (c *AssignmentRedisCache) List() ([]*model.Assignment, error) {
	val, err := c.rdb.Get(c.ctx, "assignment:list").Result()
	if err != nil {
		return nil, err
	}
	var assignments []*model.Assignment
	if err := json.Unmarshal([]byte(val), &assignments); err != nil {
		return nil, err
	}
	return assignments, nil
}

func (c *AssignmentRedisCache) SetList(assignments []*model.Assignment) error {
	data, err := json.Marshal(assignments)
	if err != nil {
		return err
	}
	return c.rdb.Set(c.ctx, "assignment:list", data, 10*time.Minute).Err()
}
