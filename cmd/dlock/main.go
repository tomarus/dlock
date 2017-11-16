// Command dlock aquires a distributed lock from redis.
// Exits with 0 if the lock was acquired successfully.
// Exits with 1 if the lock is already held by another process.
// Exits with -1 if a Redis error occurs.
// On (Redis) errors it outputs an errormessage to stderr.
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/tomarus/dlock"
)

func main() {
	nameOpt := flag.String("name", "", "Name of the lock")
	expireOpt := flag.Int("expire", 300, "Nr of seconds after which to expire lock")
	redisOpt := flag.String("redis", "redis:6379", "Address of Redis server")
	flag.Parse()

	dl, err := dlock.New(*redisOpt, *nameOpt, *expireOpt)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't initialize redis: %v\n", err)
		os.Exit(-1)
	}

	ok, err := dl.Acquire()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't acquire lock: %v\n", err)
		os.Exit(-1)
	}
	if ok {
		// lock acquired
		os.Exit(0)
	}
	// Lock already held.
	os.Exit(1)
}
