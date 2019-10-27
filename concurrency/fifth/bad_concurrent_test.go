package fifth

import (
	"testing"
	"time"
)

func Dowork(done <-chan interface{}, nums ...int) (<-chan interface{}, <-chan int) {
	heartbeat := make(chan interface{}, 1) // <1>
	intStream := make(chan int)
	go func() {
		defer close(heartbeat)
		defer close(intStream)

		time.Sleep(2 * time.Second)

		for _, n := range nums {
			select { // <2>
			case heartbeat <- struct{}{}:
			default: // <3>
			}

			select {
			case <-done:
				return
			case intStream <- n:
			}
		}
	}()

	return heartbeat, intStream
}

func TestDoWork_GenerateAllNumbers(t *testing.T) {
	done := make(chan interface{})
	defer close(done)

	intSlice := []int{0, 1, 2, 3}
	heartbeat, results := Dowork(done, intSlice...)

	// heartbeatのchannelを受信してから処理する。
	<-heartbeat

	for i, expected := range intSlice {
		select {
		case r := <-results:
			if r != expected {
				t.Errorf(
					"index %v: expected %v, but received %v,",
					i,
					expected,
					r,
				)
			}
		case <-time.After(1 * time.Second):
			t.Fatal("test timed out")
		}
	}
}