package ac

// DiscoverService discovery interface
type DiscoverService interface {
	GetService(key string) (string, error)
}
