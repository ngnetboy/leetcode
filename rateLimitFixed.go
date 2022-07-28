package main

import (
	"errors"
	"fmt"
	"time"
)

type RateLimitFix struct {
	Interval      int64
	Limit         int64
	Count         int64
	LastTimeStamp time.Time
}

func (r *RateLimitFix) TryAcquire() error {
	now := time.Now()
	elapsedTime := now.Sub(r.LastTimeStamp).Seconds()
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

	for i := 0; i < 20; i++ {
		if err := rlf.TryAcquire(); err != nil {
			fmt.Println("acquire failed, err: ", err, i)
		}
		time.Sleep(time.Second * 1)
	}

}
