package 单例Singleton

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestGetInstance(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(300)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			IncrCount1()
		}()
	}
	for i := 0; i < 200; i++ {
		go func() {
			defer wg.Done()
			IncrCount2()
		}()
	}
	wg.Wait()
	assert.Equal(t, int64(300), GetInstance().GetCount())
}
