package redis

import (
	"errors"
	"github.com/gomodule/redigo/redis"
	"time"
)

type Client struct {
	client *redis.Pool
}

type ConnArgs struct {
	Enable      bool
	Host        string
	Port        int32
	Password    string
	Index       int32
	TTL         int32
	IdleTimeout time.Duration
}

func NewClient() (*Client, error) {
	//todo 参数通过go ini获取
	client := &Client{client: &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "localhost:6379")
			if err != nil {
				return nil, err
			}
			_, _ = c.Do("AUTH", "redis")
			return c, err
		},
	}}
	return client, nil
}

func (c *Client) GET(key string) (string, error) {
	if key == "" {
		return "", errors.New("params err")
	}
	res, err := c.client.Get().Do("GET", key)
	if err != nil {
		return "", err
	}
	result, ok := res.([]uint8)
	if !ok {
		return "", nil
	}
	return string(result), nil
}

func (c *Client) SET(key, value string) (int64, error) {
	if key == "" || value == "" {
		return 0, errors.New("params err")
	}
	ret, err := c.client.Get().Do("SET", key, value)
	if err != nil {
		return 0, err
	}
	result, ok := ret.(int64)
	if !ok {
		return 0, err
	}
	return result, nil
}

func (c *Client) EXISTS(key string) (bool, error) {
	if key == "" {
		return false, errors.New("params err")
	}
	ret, err := c.client.Get().Do("EXISTS", key)
	if err != nil {
		return false, err
	}
	return ret.(int64) == 1, nil
}

// EXPIRE 如果key不存在则set value
func (c *Client) EXPIRE(key, value, ttl string) (int64, error) {
	if key == "" || ttl == "" {
		return -1, errors.New("params err")
	}
	isExist, _ := c.EXISTS(key)
	var ret interface{}
	if isExist {
		ret, _ = c.client.Get().Do("EXPIRE", key, ttl)
		return ret.(int64), nil
	} else {
		_, _ = c.SET(key, value)
		ret, _ = c.client.Get().Do("EXPIRE", key, ttl)
		return ret.(int64), nil
	}
}
