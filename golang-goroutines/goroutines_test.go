package golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld(number int) {
	fmt.Println("number: ", number)
}

func TestCreateGoroutines(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go RunHelloWorld(i)
	}

	time.Sleep(1 * time.Second)

}
