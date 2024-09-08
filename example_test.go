package safeio_test

import (
	"bytes"
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/k1LoW/safeio"
)

func Example() {
	b := new(bytes.Buffer)
	w := safeio.NewWriter(b)

	ctx, cancel := context.WithCancel(context.Background())

	sg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		sg.Add(1)
		go func(i int) {
			j := 0
			for {
				select {
				case <-ctx.Done():
					sg.Done()
					return
				default:
					_, _ = w.Write([]byte(fmt.Sprintf("%d:%d\n", i, j)))
				}
				j++
			}
		}(i)
	}

	time.Sleep(100 * time.Millisecond)
	cancel()
	sg.Wait()
}
