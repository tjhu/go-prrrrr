package stream

type MapFn[T any, R any] func(T) R

type Void struct{}

// StreamType
type StreamType int

const (
	IntermediateType StreamType = iota
	SourceType
	// Unused. Terminal nodes not physically created.
	// TerminalType
)
