// Copyright (c) 2020 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

package atomic

import (
	"sync"
	"testing"
)

func TestFloat32(t *testing.T) {
	addr := NewFloat32(3.4e38)
	if addr.Load() != 3.4e38 {
		t.Error(addr.Load())
	}
	addr.Store(0.2)
	if addr.Load() != 0.2 {
		t.Error(addr.Load())
	}
	var delta = float32(0.2)
	if addr.Add(delta) != 0.4 {
		t.Error(addr.Load())
	}
	if addr.Load() != 0.4 {
		t.Error(addr.Load())
	}
	var new = float32(0.5)
	if addr.Swap(new) != 0.4 {
		t.Error(addr.Load())
	}
	var old = new
	new = 0.6
	if !addr.CompareAndSwap(old, new) {
		t.Error(addr.Load())
	}
	if addr.CompareAndSwap(old, new) {
		t.Error(addr.Load())
	}
}

func TestAddFloat32(t *testing.T) {
	addr := NewFloat32(0.1)
	var wg sync.WaitGroup
	for i := 0; i < 8192; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			addr.Add(0.1)
		}()
	}
	wg.Wait()
}

func TestCompareAndSwapFloat32(t *testing.T) {
	addr := NewFloat32(0.1)
	var wg sync.WaitGroup
	for i := 0; i < 8192; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			addr.CompareAndSwap(0.1, 0.2)
		}()
	}
	wg.Wait()
}

func TestSwapFloat32(t *testing.T) {
	addr := NewFloat32(0.1)
	var wg sync.WaitGroup
	for i := 0; i < 8192; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			addr.Swap(0.1)
		}()
	}
	wg.Wait()
}
