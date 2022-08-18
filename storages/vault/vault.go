package vault

type vaultStorage struct {
}

func NewVaultStorage(options ...string) *vaultStorage {
	rs := &vaultStorage{}
	return rs
}

// Incr increments the counter for a given rate limit key
func (m *vaultStorage) Incr(key string, expiry int, elastic_expiry bool, amount int) int {
	return 0
}

func (m *vaultStorage) Get(key string) int {
	return 0
}

func (m *vaultStorage) GetExpiry(key string) int {
	return 0
}
func (m *vaultStorage) Check() bool {
	return false
}
func (m *vaultStorage) Reset() int {
	return 0
}
func (m *vaultStorage) Clear(key string) {

}
func (m *vaultStorage) AcquireEntry(key string, limit int, expiry int, amount int) bool {
	return false
}
func (m *vaultStorage) GetMovingWindow(key string, limit int, expiry int) (int, int) {
	return 0, 0
}
