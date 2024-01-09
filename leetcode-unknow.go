package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

type Detail struct {
	ID string
}

func Get(ctx context.Context, id string) (*Detail, error) {
	//you can call this function directly
	time.Sleep(time.Second * 3)
	if id == "3" {
		return nil, errors.New("error id is 3")
	}
	return &Detail{ID: id}, nil
}

func GetAll(ctx context.Context, ids []string) (map[string]*Detail, error) {
	var swg sync.WaitGroup
	var err error
	result := make(map[string]*Detail, len(ids))
	detailChan := make(chan *Detail, 1)
	doneChan := make(chan struct{}, 1)
	errChan := make(chan error, 1)
	defer close(detailChan)
	defer close(doneChan)
	defer close(errChan)

	ctx, cancel := context.WithCancel(ctx)

	for _, value := range ids {
		swg.Add(1)
		go func(ctx context.Context, v string) {
			defer swg.Done()
			res, err := Get(ctx, v)
			if err != nil {
				fmt.Println("get error ", err, v)
				errChan <- err
				return
			}

			select {
			case <-ctx.Done():
				return
			case detailChan <- res:
				return
			}

		}(ctx, value)
	}

	go func(ctx context.Context) {
		for value := range detailChan {
			fmt.Println("range ", value)
			select {
			case <-ctx.Done():
				return
			default:
				result[value.ID] = value
			}
		}
	}(ctx)

	go func() {
		swg.Wait()
		doneChan <- struct{}{}
	}()

	for {
		select {
		case err = <-errChan:
			fmt.Println("select error:", err)
			cancel()
		case <-doneChan:
			fmt.Println("select done")
			return result, err
		}
	}
}

// func main() {
// 	str := []string{"1", "2", "3", "4", "5", "6"}
// 	result, _ := GetAll(context.Background(), str)
// 	fmt.Println("end", result)
// }
