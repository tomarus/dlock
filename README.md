[![Go Report Card](https://goreportcard.com/badge/github.com/tomarus/dlock)](https://goreportcard.com/report/github.com/tomarus/dlock)
[![GoDoc](https://godoc.org/github.com/tomarus/dlock?status.svg)](https://godoc.org/github.com/tomarus/dlock)

DLock
=====

DLock is a basic distributed Redis backed lock package and commandline utility written in Go.

It's usage lives in crontabs running on multiple servers of which only a single server should perform a task.

## Installation

```
go install github.com/tomarus/dlock/cmd/dlock
```

## Example cron entries

Running a tool the first day of the month:

```
0 0 1 * * root dlock -name testing -redis redis:6379 && your_script.sh
```

This crontab can be installed to many servers and it will run on only one server.

The dlock utility should not necessarily run from crontab ofcourse. This is just an example.

## Exit codes

dlock returns 0 when the lock was acquired successfully.

dlock returns 1 when the lock was already held by someone else.

dlock returns -1 when a redis error occurs. An errormessage is also logged to stderr in this case.

## Full options

```
dlock -name name_of_lock -redis your_host:6379 -expire 300
```

You cannot currently remove locks, you should let it expire instead. Expire should always be higher than the theoretical maximum runtime of your process and it should account for any time differences between servers.

## Code example

Suppose you have a daemon running on multiple servers which receive data at random intervals, like every few seconds of every few minutes.

Now you want to create a global report on any server, but not more than once every 3 minutes:

```
	func main() {
		dl, _ := dlock.New("addr:6379", "mykey", 180)
	}

	func loop() {
		receiveSomeData()
		ok, _ := dl.Acquire()
		if ok {
			makeReport()
		}
	}
```

## Disclaimer

There are probably better tools which solve this problem, like zookeeper or specialized scheduling containers.

This tool was just built in a few hours because I needed it, and I didn't want to have too much dependencies.
