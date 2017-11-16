// Package dlock provides a distributed lock mechanism using Redis.
// It uses redis atomic SETNX method to set the lock key.
package dlock

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

// DLock contains the mostly internal structure for handling locks.
type DLock struct {
	name   string
	expire int
	conn   redis.Conn
}

// New creates a new lock, but doesn't acquire it yet.
// It returns an error if Redis couldn't be contacte4d.
func New(addr, name string, expire int) (*DLock, error) {
	db, err := redis.Dial("tcp", addr)
	if err != nil {
		return nil, fmt.Errorf("dlock: can't connect to redis: %v", err)
	}
	return &DLock{name, expire, db}, nil
}

// Acquire aquires the lock, returning false if the lock is
// already held or true if the lock was acquired successfully.
// An error is returned if Redis couldn't be contacted or some
// other errors were returned from redis.
func (d *DLock) Acquire() (bool, error) {
	res, err := d.conn.Do("SETNX", d.name, "1")
	if err != nil {
		return false, fmt.Errorf("can't setnx lock %s in redis: %v", d.name, err)
	}
	t, err := redis.Int64(res, err)
	if err != nil {
		return false, fmt.Errorf("unexpected parse result from redis: %v", err)
	}
	if t == 0 {
		return false, nil
	}
	_, err = d.conn.Do("EXPIRE", d.name, d.expire)
	if err != nil {
		return false, fmt.Errorf("can't set expire in redis: %v", err)
	}
	return true, nil
}
