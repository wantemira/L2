package main

import (
	"fmt"
	"time"
)

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Printf("done after %v", time.Since(start))
}

func or(channels ...<-chan interface{}) <-chan interface{} {
	result := make(chan interface{})
	if len(channels) == 0 {
		return nil
	}
	if len(channels) == 1 {
		return channels[0]
	}
	go func() {
		defer close(result)
		if len(channels) == 2 {
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		} else if len(channels) == 3 {
			select {
			case <-channels[0]:
			case <-channels[1]:
			case <-channels[2]:
			}
		} else {
			select {
			case <-channels[0]:
			case <-channels[1]:
			case <-channels[2]:
			case <-or(channels[3:]...):
			}
		}
	}()

	return result
}
