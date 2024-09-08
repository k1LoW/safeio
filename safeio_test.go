package safeio

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestWriter(t *testing.T) {
	b := new(bytes.Buffer)
	w := NewWriter(b)

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

	got := map[int]int{}
	s := bufio.NewScanner(b)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		splitted := strings.Split(s.Text(), ":")
		if len(splitted) != 2 {
			t.Errorf("unexpected line: %s", s.Text())
		}
		ii, err := strconv.Atoi(splitted[0])
		if err != nil {
			t.Fatal(err)
		}
		jj, err := strconv.Atoi(splitted[1])
		if err != nil {
			t.Fatal(err)
		}
		if _, ok := got[ii]; ok {
			if got[ii]+1 != jj {
				t.Errorf("i: %d, j: %d", ii, jj)
			}
			got[ii] = jj
		} else {
			if jj != 0 {
				t.Errorf("i: %d, j: %d", ii, jj)
			}
			got[ii] = jj
		}
	}
}

func TestSwitch(t *testing.T) {
	b := new(bytes.Buffer)
	w := NewWriter(b)

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

	time.Sleep(50 * time.Millisecond)
	b2 := new(bytes.Buffer)
	w.Switch(b2)
	time.Sleep(50 * time.Millisecond)
	cancel()
	sg.Wait()

	got := map[int]int{}
	s := bufio.NewScanner(b)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		splitted := strings.Split(s.Text(), ":")
		if len(splitted) != 2 {
			t.Errorf("unexpected line: %s", s.Text())
		}
		ii, err := strconv.Atoi(splitted[0])
		if err != nil {
			t.Fatal(err)
		}
		jj, err := strconv.Atoi(splitted[1])
		if err != nil {
			t.Fatal(err)
		}
		if _, ok := got[ii]; ok {
			if got[ii]+1 != jj {
				t.Errorf("i: %d, j: %d", ii, jj)
			}
			got[ii] = jj
		} else {
			if jj != 0 {
				t.Errorf("i: %d, j: %d", ii, jj)
			}
			got[ii] = jj
		}
	}

	s2 := bufio.NewScanner(b)
	s2.Split(bufio.ScanLines)
	for s2.Scan() {
		splitted := strings.Split(s.Text(), ":")
		if len(splitted) != 2 {
			t.Errorf("unexpected line: %s", s.Text())
		}
		ii, err := strconv.Atoi(splitted[0])
		if err != nil {
			t.Fatal(err)
		}
		jj, err := strconv.Atoi(splitted[1])
		if err != nil {
			t.Fatal(err)
		}
		if _, ok := got[ii]; ok {
			if got[ii]+1 != jj {
				t.Errorf("i: %d, j: %d", ii, jj)
			}
			got[ii] = jj
		} else {
			if jj != 0 {
				t.Errorf("i: %d, j: %d", ii, jj)
			}
			got[ii] = jj
		}
	}
}
