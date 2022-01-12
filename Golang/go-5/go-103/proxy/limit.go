package main

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
)

var counter = &LimitCounter{v: map[string]*Limit{}}

type LimitCounter struct{
	sync.Mutex
	v map[string]*Limit
}

type Limit struct {
	count int
	ttl time.Time
}

type LimitProxy struct {
	key string
	limit int
	ttl time.Duration
}

func NewLimitProxy(key string, limit int ttl time.Duration) LimitProxy {
	return LimitProxy{
		key: key,
		limit: limit,
		ttl: ttl,
	}
}

func ResetLimitHandler(c *fiber.Ctx) error {
		key:= strings.TrimPrefix(c.Path(),"/limit")
		if _,ok := counter.v[key]; !ok {
			return fiber.ErrNotFound
		}
		counter.Delete(key)
		c.Response().SetStatusCode(fiber.StatusNoContent)
		return nil
}
func (p LimitProxy) Accept(key string) bool {
	return p.key == key
}

func (p LimitProxy) Proxy(c *fiber.Ctx) error {
	path := c.Path()
	if r, ok := counter.v[path]; ok && r.count >= p.limit {
		if r.ttl.After(time.Now()) {
			c.Response().SetStatusCode(fiber.StatusTooManyRequests)
			fmt.Printf("request limit reached for %s \n", path)
			return fiber.ErrTooManyRequests
		} else {
			fmt.Printf("resetting counter values \n")
			counter.Set(path, p.ttl)
		}
	} else if !ok {
		counter.Set(path, p.ttl)
	}
	if err := c.SendString("Go Turkiye - 103 Http Package"); err != nil {
		return err
	}
	counter.Increment(path)
	return nil
}

func (l *LimitCounter) Set(key string, ttl time.Duration) {
	l.Lock()
	l.v[key] = &Limit{count: 0, ttl: time.Now().Add(ttl)}
	l.Unlock()
}

func (l *LimitCounter) Increment(key string) {
	l.Lock()
	l.v[key].count++
	l.Unlock()
}

func (l *LimitCounter) Delete(key string) {
	l.Lock()
	delete(l.v, key)
	l.Unlock()
}