package mongodb

type mongoDBStorage struct {
}

func NewMongoDBStorage(options ...string) *mongoDBStorage {
	rs := &mongoDBStorage{}
	return rs
}

// Incr increments the counter for a given rate limit key
func (m *mongoDBStorage) Incr(key string, expiry int, elastic_expiry bool, amount int) int {
	return 0
}

func (m *mongoDBStorage) Get(key string) int {
	return 0
}

func (m *mongoDBStorage) GetExpiry(key string) int {
	return 0
}
func (m *mongoDBStorage) Check() bool {
	return false
}
func (m *mongoDBStorage) Reset() int {
	return 0
}
func (m *mongoDBStorage) Clear(key string) {

}
func (m *mongoDBStorage) AcquireEntry(key string, limit int, expiry int, amount int) bool {
	return false
}
func (m *mongoDBStorage) GetMovingWindow(key string, limit int, expiry int) (int, int) {
	return 0, 0
}
