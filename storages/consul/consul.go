package consul

type consulStorage struct {
}

func NewConsulStorage(options ...string) *consulStorage {
	rs := &consulStorage{}
	return rs
}

// Incr increments the counter for a given rate limit key
func (m *consulStorage) Incr(key string, expiry int, elastic_expiry bool, amount int) int {
	return 0
}

func (m *consulStorage) Get(key string) int {
	return 0
}

func (m *consulStorage) GetExpiry(key string) int {
	return 0
}
func (m *consulStorage) Check() bool {
	return false
}
func (m *consulStorage) Reset() int {
	return 0
}
func (m *consulStorage) Clear(key string) {

}
func (m *consulStorage) AcquireEntry(key string, limit int, expiry int, amount int) bool {
	return false
}
func (m *consulStorage) GetMovingWindow(key string, limit int, expiry int) (int, int) {
	return 0, 0
}
