package atomic

import (
	"sync"
	"testing"
)

func TestValue(t *testing.T) {
	var val = "Hello World"
	addr := NewValue(val)
	if addr.Load() != val {
		t.Error(addr.Load())
	}
	addr.v = nil
	if addr.Load() != nil {
		t.Error(addr.Load())
	}
	addr.Store(val[:5])
	if addr.Load().(string) != val[:5] {
		t.Error(addr.Load())
	}
}

func TestInitValue(t *testing.T) {
	addr := &Value{}
	var wg sync.WaitGroup
	for i := 0; i < 512; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			initValue(addr)
			if addr.v == nil {
				t.Error("should not be nil")
			}
		}()
	}
	wg.Wait()
}
