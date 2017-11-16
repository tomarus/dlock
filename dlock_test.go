package dlock

import "testing"

const addr = "192.168.0.3:6379"

func TestAcquire(t *testing.T) {
	t.Skip("You need redis to run this test")

	lock, err := New(addr, "testlock", 10)
	if err != nil {
		t.Fatal(err)
	}

	ok, err := lock.Acquire()
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("Lock should be held")
	}

	// retry immediately, lock should  be held by other instance.
	ok, _ = lock.Acquire()
	if ok {
		t.Fatal("Lock should be not held")
	}
}
