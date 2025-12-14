package events

// Publisher is the interface for the event publisher
type Publisher interface {
	Publish(eventType string, payload any, metadata map[string]string) error
	Close() error
}
