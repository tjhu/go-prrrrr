package stream

type Void struct{}
type Any interface{}

// StreamType
type StreamType int

const (
	StreamTypeSource StreamType = iota
	StreamTypeBatchedSource
	StreamTypeIntermediate
	StreamTypeBatchedIntermediate
	// Terminal nodes not physically created so there's no type for it.
)
