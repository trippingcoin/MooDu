package cache

import (
	"sync"

	"cs/course-service/internal/model"
)

type inMemoryCourseCache struct {
	courses map[string]*model.Course
	list    []*model.Course
	mu      sync.RWMutex
	hasList bool
}

func NewInMemoryCourseCache() CourseCache {
	return &inMemoryCourseCache{
		courses: make(map[string]*model.Course),
	}
}

func (c *inMemoryCourseCache) Get(id string) (*model.Course, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	course, ok := c.courses[id]
	return course, ok
}

func (c *inMemoryCourseCache) Set(id string, course *model.Course) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.courses[id] = course
}

func (c *inMemoryCourseCache) Delete(id string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.courses, id)
}

func (c *inMemoryCourseCache) List() ([]*model.Course, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if !c.hasList {
		return nil, false
	}
	return c.list, true
}

func (c *inMemoryCourseCache) SetList(courses []*model.Course) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.list = courses
	c.hasList = true
}

func (c *inMemoryCourseCache) ClearList() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.hasList = false
	c.list = nil
}
