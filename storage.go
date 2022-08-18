package ratelimits

import (
	"time"
)

type StorageBase interface {
	Incr() int
	Get(key string) int       // use key to get the counter value
	GetExpiry(key string) int // use key to get the expiry
	Check() bool              //  check if storage is healthy
	Reset() int               //  reset storage to clear limits
	Clear(key string)         // resets the rate limit key
}

type MovingWindowSupport interface {
	AcquireEntry(key string, limit int, expiry int, amount int) bool
	GetMovingWindow(key string, limit int, expiry int) (int, int)
}

type memoryStorage struct {
	Storage     *Counter
	Expirations map[string]int
	Event       map[string][]string
	Timer       *time.Timer
}

func NewMemoryStorage(options ...string) *memoryStorage {
	ms := &memoryStorage{
		Storage: NewCounter(),
	}
	return ms
}

// Incr increments the counter for a given rate limit key
func (m *memoryStorage) Incr(key string, expiry int, elastic_expiry bool, amount int) int {
	return 0
}

func (m *memoryStorage) Get(key string) int {
	return 0
}

func (m *memoryStorage) GetExpiry(key string) int {
	return 0
}
func (m *memoryStorage) Check() bool {
	return false
}
func (m *memoryStorage) Reset() int {
	return 0
}
func (m *memoryStorage) Clear(key string) {

}
func (m *memoryStorage) AcquireEntry(key string, limit int, expiry int, amount int) bool {
	return false
}
func (m *memoryStorage) GetMovingWindow(key string, limit int, expiry int) (int, int) {
	return 0, 0
}
