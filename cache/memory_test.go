package cache

import (
	"context"
	"testing"

	"github.com/binbinly/pkg/codec"
	"github.com/stretchr/testify/assert"
)

func Test_NewMemoryCache(t *testing.T) {
	asserts := assert.New(t)

	client := NewMemoryCache(WithCodec(codec.JSONEncoding{}), WithPrefix("demo"))
	asserts.NotNil(client)
}

func TestMemoStore_Set(t *testing.T) {
	asserts := assert.New(t)

	store := NewMemoryCache(WithCodec(codec.JSONEncoding{}), WithPrefix("demo"))
	err := store.Set(context.Background(), "test-key", "test-val", -1)
	asserts.NotNil(err)
}

func TestMemoStore_Get(t *testing.T) {
	asserts := assert.New(t)
	store := NewMemoryCache(WithCodec(codec.JSONEncoding{}), WithPrefix("demo"))
	ctx := context.Background()

	type testStruct struct {
		Name string
		Age  int
	}

	// normal
	{
		testKey := "test-key2"
		testVal := testStruct{
			Name: "test-name",
			Age:  18,
		}
		err := store.Set(ctx, testKey, &testVal, 60)
		asserts.Nil(err)

		var gotVal testStruct
		err = store.Get(ctx, testKey, &gotVal)
		asserts.Nil(err)
		asserts.NotNil(gotVal)
	}
}
