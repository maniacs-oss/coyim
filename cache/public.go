package cache

import "time"

type Simple interface {
	Get(key string) (interface{}, bool)
	GetOrCompute(key string, creator func(key string) interface{}) (interface{}, bool)
	Put(key string, value interface{}) bool
	PutIfAbsent(key string, creator func(key string) interface{}) bool
	Has(key string) bool
	Remove(key string) bool
	Clear()
}

type WithExpiry interface {
	Simple
	GetOrComputeTimed(key string, lifetime time.Duration, creator func(key string) interface{}) (interface{}, bool)
	PutTimed(key string, lifetime time.Duration, value interface{}) bool
	PutTimedIfAbsent(key string, lifetime time.Duration, creator func(key string) interface{}) bool
}
