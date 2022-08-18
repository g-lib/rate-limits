package redis

type redisStorage struct {
}

func NewredisStorage(options ...string) *redisStorage {
	rs := &redisStorage{}
	return rs
}

// Incr increments the counter for a given rate limit key
func (s *redisStorage) Incr(key string, expiry int, elastic_expiry bool, amount int) int {
	return 0
}

func (s *redisStorage) Get(key string) int {
	return 0
}

func (s *redisStorage) GetExpiry(key string) int {
	return 0
}
func (s *redisStorage) Check() bool {
	return false
}
func (s *redisStorage) Reset() int {
	return 0
}
func (s *redisStorage) Clear(key string) {

}
func (s *redisStorage) AcquireEntry(key string, limit int, expiry int, amount int) bool {
	return false
}
func (s *redisStorage) GetMovingWindow(key string, limit int, expiry int) (int, int) {
	return 0, 0
}
