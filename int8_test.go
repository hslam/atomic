// Copyright (c) 2020 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

package atomic

import (
	"sync"
	"testing"
)

func TestInt8(t *testing.T) {
	addr := NewInt8(1)
	if addr.Load() != 1 {
		t.Error(addr.Load())
	}
	addr.Store(2)
	if addr.Load() != 2 {
		t.Error(addr.Load())
	}
	var delta = int8(2)
	if addr.Add(delta) != 4 {
		t.Error(addr.Load())
	}
	if addr.Load() != 4 {
		t.Error(addr.Load())
	}
	var new = int8(5)
	if addr.Swap(new) != 4 {
		t.Error(addr.Load())
	}
	var old = new
	new = 6
	if !addr.CompareAndSwap(old, new) {
		t.Error(addr.Load())
	}
	if addr.CompareAndSwap(old, new) {
		t.Error(addr.Load())
	}
}

func TestAddInt8(t *testing.T) {
	addr := NewInt8(1)
	var wg sync.WaitGroup
	for i := 0; i < 512; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			addr.Add(1)
		}()
	}
	wg.Wait()
}

func TestCompareAndSwapInt8(t *testing.T) {
	addr := NewInt8(1)
	var wg sync.WaitGroup
	for i := 0; i < 512; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			addr.CompareAndSwap(1, 2)
		}()
	}
	wg.Wait()
}

func TestSwapInt8(t *testing.T) {
	addr := NewInt8(1)
	var wg sync.WaitGroup
	for i := 0; i < 512; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			addr.Swap(1)
		}()
	}
	wg.Wait()
}

func BenchmarkSwapInt8(b *testing.B) {
	addr := NewInt8(1)
	for i := 0; i < b.N; i++ {
		addr.Swap(1)
	}
}

func BenchmarkCompareAndSwapInt8(b *testing.B) {
	addr := NewInt8(1)
	for i := 0; i < b.N; i++ {
		addr.CompareAndSwap(1, 2)
	}
}

func BenchmarkAddInt8(b *testing.B) {
	addr := NewInt8(1)
	for i := 0; i < b.N; i++ {
		addr.Add(1)
	}
}

func BenchmarkStoreInt8(b *testing.B) {
	addr := NewInt8(1)
	for i := 0; i < b.N; i++ {
		addr.Store(1)
	}
}

func BenchmarkLoadInt8(b *testing.B) {
	addr := NewInt8(1)
	for i := 0; i < b.N; i++ {
		addr.Load()
	}
}
