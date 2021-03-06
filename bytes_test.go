// Copyright (c) 2020 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

package atomic

import (
	"sync"
	"testing"
)

func TestBytesEqual(t *testing.T) {
	var a = []byte{1, 2, 3}
	var b = []byte{1, 2, 3}
	if !bytesEqual(a, b) {
		t.Error("fail")
	}
}

func TestBytes(t *testing.T) {
	var val = []byte{1, 2, 3}
	addr := NewBytes(val)
	if !bytesEqual(addr.Load(), val) {
		t.Error(addr.Load())
	}
	addr.Store(val[:2])
	if !bytesEqual(addr.Load(), val[:2]) {
		t.Error(addr.Load())
	}
	var delta = val[2:]
	if !bytesEqual(addr.Add(delta), val) {
		t.Error(addr.Load())
	}
	if !bytesEqual(addr.Load(), val) {
		t.Error(addr.Load())
	}
	var new = []byte{4, 5, 6}
	if !bytesEqual(addr.Swap(new), val) {
		t.Error(addr.Load())
	}
	var old = new
	new = []byte{7, 8, 9}
	if !addr.CompareAndSwap(old, new) {
		t.Error(addr.Load())
	}
	if addr.CompareAndSwap(old, new) {
		t.Error(addr.Load())
	}

	addr = &Bytes{}
	if addr.Load() != nil || !bytesEqual(addr.Load(), []byte{}) {
		t.Error(addr.Load())
	}
}

func TestAddBytes(t *testing.T) {
	addr := NewBytes(nil)
	var wg sync.WaitGroup
	for i := 0; i < 8192; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			addr.Add(nil)
		}()
	}
	wg.Wait()
}

func TestCompareAndSwapBytes(t *testing.T) {
	addr := NewBytes(nil)
	var wg sync.WaitGroup
	for i := 0; i < 8192; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			addr.CompareAndSwap(nil, nil)
		}()
	}
	wg.Wait()
}

func TestSwapBytes(t *testing.T) {
	addr := NewBytes(nil)
	var wg sync.WaitGroup
	for i := 0; i < 8192; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			addr.Swap(nil)
		}()
	}
	wg.Wait()
}

func BenchmarkSwapBytes(b *testing.B) {
	addr := NewBytes(nil)
	for i := 0; i < b.N; i++ {
		addr.Swap(nil)
	}
}

func BenchmarkCompareAndSwapBytes(b *testing.B) {
	addr := NewBytes(nil)
	for i := 0; i < b.N; i++ {
		addr.CompareAndSwap(nil, nil)
	}
}

func BenchmarkAddBytes(b *testing.B) {
	addr := NewBytes(nil)
	for i := 0; i < b.N; i++ {
		addr.Add(nil)
	}
}

func BenchmarkStoreBytes(b *testing.B) {
	addr := NewBytes(nil)
	for i := 0; i < b.N; i++ {
		addr.Store(nil)
	}
}

func BenchmarkLoadBytes(b *testing.B) {
	addr := NewBytes(nil)
	for i := 0; i < b.N; i++ {
		addr.Load()
	}
}
