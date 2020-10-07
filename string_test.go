// Copyright (c) 2020 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

package atomic

import (
	"sync"
	"sync/atomic"
	"testing"
)

func TestString(t *testing.T) {
	var val = "Hello World"
	addr := NewString(val)
	if addr.Load() != val {
		t.Error(addr.Load())
	}
	addr.v = nil
	if addr.Load() != "" {
		t.Error(addr.Load())
	}
	addr.Store(val[:5])
	if addr.Load() != val[:5] {
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

	addr = &String{v: &atomic.Value{}}
	if addr.Load() != "" {
		t.Error(addr.Load())
	}
}

func TestAddString(t *testing.T) {
	addr := NewString("")
	var wg sync.WaitGroup
	for i := 0; i < 512; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			addr.Add("")
		}()
	}
	wg.Wait()
}

func TestCompareAndSwapString(t *testing.T) {
	addr := NewString("")
	var wg sync.WaitGroup
	for i := 0; i < 512; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			addr.CompareAndSwap("", "")
		}()
	}
	wg.Wait()
}

func TestSwapString(t *testing.T) {
	addr := NewString("")
	var wg sync.WaitGroup
	for i := 0; i < 512; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			addr.Swap("")
		}()
	}
	wg.Wait()
}

func TestInitString(t *testing.T) {
	addr := &String{}
	var wg sync.WaitGroup
	for i := 0; i < 512; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			initString(addr)
			if addr.v == nil {
				t.Error("should not be nil")
			}
		}()
	}
	wg.Wait()
}
