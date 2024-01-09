package cache

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
	"time"
)

type RedisManager interface {
	// Close 关闭链接
	Close()
	// Exists 是否存在值
	Exists(key string) bool

	// Set 设置保存信息
	// 默认string类型
	Set(key, value string) (bool, error)

	// ExpireSet 设置保存信息，带过期时间
	// 默认string类型
	ExpireSet(key, value string, sec int64) (bool, error)

	// Get 读取信息
	// 默认string类型
	Get(key string) (string, error)

	// GetBytes 读取byte数据
	GetBytes(key string) ([]byte, error)

	// Delete 删除key
	Delete(key string)
}

type redisManager struct {
	opt  *RedisConnOpt
	pool *redis.Pool
}

// NewManager 初始化单利缓存管理
func NewManager(opt *RedisConnOpt) RedisManager {
	return &redisManager{
		opt:  opt,
		pool: newRedisPool(opt),
	}
}

func newRedisPool(opt *RedisConnOpt) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 30 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", opt.Host, opt.Port))

			if err != nil {
				log.Fatalf("redis dial err: %v", err)
				return nil, err
			}

			if opt.Password != "" {
				if _, err = c.Do("AUTH", opt.Password); err != nil {
					_ = c.Close()
					log.Fatalf("redis auth err: %v", err)
					return nil, err
				}
			}

			return c, nil
		},
	}
}

func (m *redisManager) Exists(key string) bool {
	conn := m.pool.Get()
	defer m.closeConn(conn)

	exists, err := conn.Do("EXISTS", key)
	if err != nil {
		log.Printf("redis conn do exists err: %v", err)
		return false
	}

	i, err := redis.Int(exists, err)
	if err != nil {
		log.Printf("redis reply to int err: %v", err)
		return false
	}
	return i > 0
}

func (m *redisManager) Close() {
	_ = m.pool.Close()
}

func (m *redisManager) Set(key, value string) (bool, error) {
	conn := m.pool.Get()
	defer m.closeConn(conn)

	reply, err := conn.Do("SET", key, value)
	if err != nil {
		log.Printf("redis conn do set err: %v", err)
		return false, err
	}

	str, err := redis.String(reply, err)
	if err != nil {
		log.Printf("redis reply to string err: %v", err)
		return false, err
	}

	return str == "OK", nil
}

func (m *redisManager) ExpireSet(key, value string, sec int64) (bool, error) {
	conn := m.pool.Get()
	defer m.closeConn(conn)

	reply, err := conn.Do("SETEX", key, sec, value)
	if err != nil {
		log.Printf("redis conn do set err: %v", err)
		return false, err
	}

	str, err := redis.String(reply, err)
	if err != nil {
		log.Printf("redis reply to string err: %v", err)
		return false, err
	}

	return str == "OK", nil
}

func (m *redisManager) Get(key string) (string, error) {
	conn := m.pool.Get()
	defer m.closeConn(conn)

	return redis.String(conn.Do("GET", key))
}

func (m *redisManager) GetBytes(key string) ([]byte, error) {
	conn := m.pool.Get()
	defer m.closeConn(conn)

	return redis.Bytes(conn.Do("GET", key))
}

func (m *redisManager) Delete(key string) {
	conn := m.pool.Get()
	defer m.closeConn(conn)

	// 用于循环查询key
	iter := 0
	var keys []string
	for {
		// SCAN cursor [MATCH pattern] [COUNT count]
		// cursor - 游标。
		// pattern - 匹配的模式。
		// count - 指定从数据集里返回多少元素，默认值为 10 。

		// redis 127.0.0.1:6379> scan 0 MATCH *11*
		// 1) "288"
		// 2) 1) "key:911"
		if arr, err := redis.Values(conn.Do("SCAN", iter, "MATCH", "*"+key+"*")); err != nil {
			log.Printf("redis SCAN err: %v", err)
			return
		} else {
			iter, _ = redis.Int(arr[0], nil)
			k, _ := redis.Strings(arr[1], nil)

			for _, value := range k {
				keys = append(keys, value)
			}
		}

		if iter == 0 {
			break
		}
	}

	// 开启事务
	_ = conn.Send("MULTI")
	for _, value := range keys {
		_ = conn.Send("DEL", value)
	}

	_, err := redis.Values(conn.Do("EXEC"))
	if err != nil {
		log.Printf("redis EXEC err: %v", err)
	}
}

func (m *redisManager) closeConn(conn redis.Conn) {
	_ = conn.Close()
}
