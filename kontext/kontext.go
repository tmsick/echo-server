package kontext

const (
	keyRequestID = "request_id"
)

type Kontext interface {
	// Get retrieves data from the context.
	Get(key string) any

	// Set saves data in the context.
	Set(key string, val any)
}
