package plays

import (
	"errors"
	"sync"
	"time"
)

func WaitTimeout(entrances []string) ([]string, error) {
	var wg sync.WaitGroup

	resultChannel := make(chan string)
	defer close(resultChannel)

	for _, i := range entrances {
		wg.Add(1)
		go func(w *sync.WaitGroup, input string) {
			defer w.Done()
			println("loading " + input)
			time.Sleep(time.Second * 2)
			resultChannel <- input
			println("loaded " + input)
		}(&wg, i)
	}
	var result []string
	go func() {
		for res := range resultChannel {
			result = append(result, res)
		}
	}()

	inTime := make(chan interface{})
	go func() {
		defer close(inTime)
		wg.Wait()
	}()

	select {
	case <-time.After(time.Second * 1):
		return nil, errors.New("end with timeout")
	case <-inTime:
		return result, nil
	}
}
