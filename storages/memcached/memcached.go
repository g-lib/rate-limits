package memcached

type memcachedStorage struct {
}

func NewMemcachedStorage(options ...string) *memcachedStorage {
	rs := &memcachedStorage{}
	return rs
}

// Incr increments the counter for a given rate limit key
func (m *memcachedStorage) Incr(key string, expiry int, elastic_expiry bool, amount int) int {
	return 0
}

func (m *memcachedStorage) Get(key string) int {
	return 0
}

func (m *memcachedStorage) GetExpiry(key string) int {
	return 0
}
func (m *memcachedStorage) Check() bool {
	return false
}
func (m *memcachedStorage) Reset() int {
	return 0
}
func (m *memcachedStorage) Clear(key string) {

}
func (m *memcachedStorage) AcquireEntry(key string, limit int, expiry int, amount int) bool {
	return false
}
func (m *memcachedStorage) GetMovingWindow(key string, limit int, expiry int) (int, int) {
	return 0, 0
}
