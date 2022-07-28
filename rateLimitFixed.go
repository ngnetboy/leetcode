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

func main() {
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
