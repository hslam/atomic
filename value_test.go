package atomic

import (
	"sync"
	"testing"
)

func TestValue(t *testing.T) {
	var val = "Hello World"
	var addFunc AddFunc = func(old, delta interface{}) (new interface{}) {
		return old.(string) + delta.(string)
	}
	addr := NewValue(val, addFunc)
	if addr.Load() != val {
		t.Error(addr.Load())
	}
	addr.Store(val[:5])
	if addr.Load().(string) != val[:5] {
		t.Error(addr.Load())
	}
	var delta = val[5:]
	if addr.Add(delta) != val {
		t.Error(addr.Load())
	}
	if addr.Load() != val {
		t.Error(addr.Load())
	}
	var new = "Foo"
	if addr.Swap(new) != val {
		t.Error(addr.Load())
	}
	var old = new
	new = "Bar"
	if !addr.CompareAndSwap(old, new) {
		t.Error(addr.Load())
	}
	if addr.CompareAndSwap(old, new) {
		t.Error(addr.Load())
	}
	addr = &Value{}
	if addr.Load() != nil {
		t.Error(addr.Load())
	}
}

func TestAddValue(t *testing.T) {
	var addFunc AddFunc = func(old, delta interface{}) (new interface{}) {
		return old.(string) + delta.(string)
	}
	addr := NewValue("", nil)
	func() {
		defer func() {
			if err := recover(); err == nil {
				t.Error("should panic")
			}
		}()
		addr.Add("")
	}()
	addr.AddFunc = addFunc
	var wg sync.WaitGroup
	for i := 0; i < 8192; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			addr.Add("")
		}()
	}
	wg.Wait()
}

func TestCompareAndSwapValue(t *testing.T) {
	addr := NewValue("", nil)
	var wg sync.WaitGroup
	for i := 0; i < 8192; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			addr.CompareAndSwap("", "")
		}()
	}
	wg.Wait()
}

func TestSwapValue(t *testing.T) {
	addr := NewValue("", nil)
	var wg sync.WaitGroup
	for i := 0; i < 8192; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			addr.Swap("")
		}()
	}
	wg.Wait()
}

func BenchmarkSwapValue(b *testing.B) {
	addr := NewValue("", nil)
	for i := 0; i < b.N; i++ {
		addr.Swap("")
	}
}

func BenchmarkCompareAndSwapValue(b *testing.B) {
	addr := NewValue("", nil)
	for i := 0; i < b.N; i++ {
		addr.CompareAndSwap("", "")
	}
}

func BenchmarkAddValue(b *testing.B) {
	var addFunc AddFunc = func(old, delta interface{}) (new interface{}) {
		return old.(string) + delta.(string)
	}
	addr := NewValue("", addFunc)
	for i := 0; i < b.N; i++ {
		addr.Add("")
	}
}

func BenchmarkStoreValue(b *testing.B) {
	addr := NewValue("", nil)
	for i := 0; i < b.N; i++ {
		addr.Store("")
	}
}

func BenchmarkLoadValue(b *testing.B) {
	addr := NewValue("", nil)
	for i := 0; i < b.N; i++ {
		addr.Load()
	}
}

func TestFastValue(t *testing.T) {
	var val = "Hello World"
	var addFunc AddFunc = func(old, delta interface{}) (new interface{}) {
		return old.(string) + delta.(string)
	}
	addr := NewFastValue(val, addFunc)
	if addr.Load() != val {
		t.Error(addr.Load())
	}
	addr.Store(val[:5])
	if addr.Load().(string) != val[:5] {
		t.Error(addr.Load())
	}
	var delta = val[5:]
	if addr.Add(delta) != val {
		t.Error(addr.Load())
	}
	if addr.Load() != val {
		t.Error(addr.Load())
	}
	var new = "Foo"
	if addr.Swap(new) != val {
		t.Error(addr.Load())
	}
	var old = addr.Load()
	new = "Bar"
	if !addr.CompareAndSwap(old, new) {
		t.Error(addr.Load())
	}
	if addr.CompareAndSwap(old, new) {
		t.Error(addr.Load())
	}
	addr = &Value{fast: true}
	if addr.Load() != nil {
		t.Error(addr.Load())
	}
}

func TestAddFastValue(t *testing.T) {
	var addFunc AddFunc = func(old, delta interface{}) (new interface{}) {
		return old.(string) + delta.(string)
	}
	addr := NewFastValue("", nil)
	func() {
		defer func() {
			if err := recover(); err == nil {
				t.Error("should panic")
			}
		}()
		addr.Add("")
	}()
	addr.AddFunc = addFunc
	var wg sync.WaitGroup
	for i := 0; i < 8192; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			addr.Add("")
		}()
	}
	wg.Wait()
}

func testCompareAndSwapFastValue(t *testing.T) {
	addr := &Value{fast: true}
	var wg sync.WaitGroup
	for i := 0; i < 64; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			addr.CompareAndSwap("", "")
		}()
	}
	wg.Wait()
}

func TestCompareAndSwapFastValue(t *testing.T) {
	for i := 0; i < 8192; i++ {
		testCompareAndSwapFastValue(t)
	}
	addr := &Value{fast: true}
	func() {
		defer func() {
			if err := recover(); err == nil {
				t.Error("should panic")
			}
		}()
		addr.CompareAndSwap("", nil)
	}()
	addr.CompareAndSwap("", "")
	func() {
		defer func() {
			if err := recover(); err == nil {
				t.Error("should panic")
			}
		}()
		addr.CompareAndSwap(nil, "")
	}()
	func() {
		defer func() {
			if err := recover(); err == nil {
				t.Error("should panic")
			}
		}()
		addr.CompareAndSwap(1, "")
	}()
	func() {
		defer func() {
			if err := recover(); err == nil {
				t.Error("should panic")
			}
		}()
		addr.CompareAndSwap("", 1)
	}()
}

func TestSwapFastValue(t *testing.T) {
	addr := NewValue("", nil)
	var wg sync.WaitGroup
	for i := 0; i < 8192; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			addr.Swap("")
		}()
	}
	wg.Wait()
}

func BenchmarkSwapFastValue(b *testing.B) {
	addr := NewFastValue("", nil)
	for i := 0; i < b.N; i++ {
		addr.Swap("")
	}
}

func BenchmarkCompareAndSwapFastValue(b *testing.B) {
	addr := NewFastValue("", nil)
	for i := 0; i < b.N; i++ {
		addr.CompareAndSwap("", "")
	}
}

func BenchmarkAddFastValue(b *testing.B) {
	var addFunc AddFunc = func(old, delta interface{}) (new interface{}) {
		return old.(string) + delta.(string)
	}
	addr := NewFastValue("", addFunc)
	for i := 0; i < b.N; i++ {
		addr.Add("")
	}
}

func BenchmarkStoreFastValue(b *testing.B) {
	addr := NewFastValue("", nil)
	for i := 0; i < b.N; i++ {
		addr.Store("")
	}
}

func BenchmarkLoadFastValue(b *testing.B) {
	addr := NewFastValue("", nil)
	for i := 0; i < b.N; i++ {
		addr.Load()
	}
}
