package util

import (
	"bytes"
	"math"
	"math/rand"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandInt(t *testing.T) {
	t.Parallel()
	assert.Equal(t, true, RandInt(1, 2) == 1)
	assert.Equal(t, true, RandInt(-1, 0) == -1)
	assert.Equal(t, true, RandInt(0, 5) >= 0)
	assert.Equal(t, true, RandInt(0, 5) < 5)
	assert.Equal(t, 2, RandInt(2, 2))
	assert.Equal(t, 2, RandInt(3, 2))
}

func TestRandUint32(t *testing.T) {
	t.Parallel()
	assert.Equal(t, true, RandUint32(1, 2) == 1)
	assert.Equal(t, true, RandUint32(0, 5) < 5)
	assert.Equal(t, uint32(2), RandUint32(2, 2))
	assert.Equal(t, uint32(2), RandUint32(3, 2))
}

func TestFastIntn(t *testing.T) {
	t.Parallel()
	for i := 1; i < 10000; i++ {
		assert.Equal(t, true, FastRandn(uint32(i)) < uint32(i))
		assert.Equal(t, true, FastIntn(i) < i)
	}
	assert.Equal(t, 0, FastIntn(-2))
	assert.Equal(t, 0, FastIntn(0))
	assert.Equal(t, true, FastIntn(math.MaxUint32) < math.MaxUint32)
	assert.Equal(t, true, FastIntn(math.MaxInt64) < math.MaxInt64)
}

func BenchmarkRandInt(b *testing.B) {
	b.Run("RandInt", func(b *testing.B) {
		for i := 1; i < b.N; i++ {
			_ = RandInt(0, i)
		}
	})
	b.Run("RandUint32", func(b *testing.B) {
		for i := 1; i < b.N; i++ {
			_ = RandUint32(0, uint32(i))
		}
	})
	b.Run("FastIntn", func(b *testing.B) {
		for i := 1; i < b.N; i++ {
			_ = FastIntn(i)
		}
	})
	b.Run("Rand.Intn", func(b *testing.B) {
		for i := 1; i < b.N; i++ {
			_ = Rand.Intn(i)
		}
	})
	b.Run("std.rand.Intn", func(b *testing.B) {
		for i := 1; i < b.N; i++ {
			_ = rand.Intn(i)
		}
	})
}

func BenchmarkRandIntParallel(b *testing.B) {
	b.Run("RandInt", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = RandInt(0, math.MaxInt32)
			}
		})
	})
	b.Run("RandUint32", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = RandUint32(0, math.MaxInt32)
			}
		})
	})
	b.Run("FastIntn", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = FastIntn(math.MaxInt32)
			}
		})
	})
	b.Run("Rand.Intn", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = Rand.Intn(math.MaxInt32)
			}
		})
	})
	var mu sync.Mutex
	b.Run("std.rand.Intn", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				mu.Lock()
				_ = rand.Intn(math.MaxInt32)
				mu.Unlock()
			}
		})
	})
}

func TestRandString(t *testing.T) {
	t.Parallel()
	a, b := RandString(777), RandString(777)
	assert.Equal(t, 777, len(a))
	assert.Equal(t, false, a == b)
	assert.Equal(t, "", RandString(-1))
}

func TestRandBytes(t *testing.T) {
	t.Parallel()
	a, err := RandBytes(777)
	assert.Nil(t, err)
	b, err := RandBytes(777)
	assert.Nil(t, err)
	assert.Equal(t, 777, len(a))
	assert.Equal(t, 777, len(b))
	assert.Equal(t, false, bytes.Equal(a, b))
}

func TestFastRandBytes(t *testing.T) {
	t.Parallel()
	a, b := FastRandBytes(777), FastRandBytes(777)
	assert.Equal(t, 777, len(a))
	assert.Equal(t, 777, len(b))
	assert.Equal(t, false, bytes.Equal(a, b))
}

func TestRandHex(t *testing.T) {
	t.Parallel()
	assert.Equal(t, 32, len(RandHex(16)))
	assert.Equal(t, 14, len(RandHex(7)))
}

func BenchmarkRandBytes(b *testing.B) {
	b.Run("RandString", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = RandString(20)
		}
	})
	b.Run("RandBytes", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			RandBytes(20)
		}
	})
	b.Run("FastRandBytes", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = FastRandBytes(20)
		}
	})
}

func BenchmarkRandBytesParallel(b *testing.B) {
	b.Run("RandString", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				RandString(20)
			}
		})
	})
	b.Run("RandBytes", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				RandBytes(20)
			}
		})
	})
	b.Run("FastRandBytes", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				FastRandBytes(20)
			}
		})
	})
}
