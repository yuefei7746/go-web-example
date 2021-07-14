package utils

import (
	"context"
	"github.com/go-redis/cache/v8"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	ctx := context.TODO()
	k := "mykey"
	v := "val"

	if err := GetCacheCli().Set(&cache.Item{
		Ctx:   ctx,
		Key:   k,
		Value: v,
		TTL:   time.Hour,
	}); err != nil {
		t.Fatal(err)
	}

	var wanted string
	if err := GetCacheCli().Get(ctx, k, &wanted); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, v, wanted)
}

func TestMembers(t *testing.T) {
	ctx := context.TODO()
	k := "device_white_list_test_600245"
	var v map[string]struct{}
	err := GetCacheCli().Once(&cache.Item{
		Ctx:   ctx,
		Key:   k,
		Value: &v,
		Do: func(item *cache.Item) (interface{}, error) {
			return GetRedisCli().SMembersMap(item.Ctx, item.Key).Result()
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.True(t, len(v) > 0)
}
