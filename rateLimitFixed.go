package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type RateLimitFix struct {
	sync.Mutex
	Interval      int64
	Limit         int64
	Count         int64
	LastTimeStamp time.Time
}

func (r *RateLimitFix) TryAcquireWait(waitSecond int) error {
	timer := time.NewTimer(time.Duration(waitSecond) * time.Second)
	done := make(chan bool)
	timeout := make(chan bool)

	go func() {
		for {
			err := r.TryAcquire()
			if err == nil {
				fmt.Println("acquire success")
				done <- true
				return
			}
			time.Sleep(time.Millisecond * 100)
			select {
			case <-timeout:
				fmt.Println("timeout quite acquire")
				return
			default:
				fmt.Println("try it again", time.Now())
			}
		}
	}()

	select {
	case <-done:
		fmt.Println("finish the acquireWait")
		timer.Stop()
		return nil
	case <-timer.C:
		timeout <- true
		return errors.New("timeout")
	}
}

func testTryAcquireWait() {
	rlf := &RateLimitFix{
		Interval: 2,
		Limit:    1,
	}

	err := rlf.TryAcquire()
	if err != nil {
		fmt.Println("acquire failed")
		return
	}
	fmt.Println("start time: ", time.Now())
	err = rlf.TryAcquireWait(5)
	fmt.Println("end time", time.Now())
	if err != nil {
		fmt.Println("acquireWait failed", err)
		return
	}

}

func (r *RateLimitFix) TryAcquire() error {
	now := time.Now()
	elapsedTime := now.Sub(r.LastTimeStamp).Seconds()
	r.Lock()
	defer r.Unlock()
	if int64(elapsedTime) >= r.Interval {
		r.LastTimeStamp = now
		r.Count = 1
		return nil
	}
	r.Count += 1
	if r.Count > r.Limit {
		return errors.New("rate limit")
	}
	return nil
}

func testTryAcquire() {
	rlf := &RateLimitFix{
		Interval: 2,
		Limit:    1,
	}

	go func() {
		for i := 0; i < 20; i++ {
			if err := rlf.TryAcquire(); err != nil {
				fmt.Println("acquire failed, err: ", err, i)
			}
			time.Sleep(time.Second * 1)
		}
	}()

	go func() {
		for i := 20; i < 40; i++ {
			if err := rlf.TryAcquire(); err != nil {
				fmt.Println("acquire failed, err: ", err, i)
			}
			time.Sleep(time.Second * 1)
		}
	}()
	time.Sleep(time.Second * 20)
}

func main() {
	testTryAcquireWait()
}
