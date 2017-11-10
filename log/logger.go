package log

// Logger interface for logging, currently stdio writer is implemented.
// More file based, database based loggers can be implemented.
type Logger interface {
	Log()
}

// Loggable - need to Replace with stdlib interface that has String method.
// models.StoreyResponse implements Loggable.
type Loggable interface {
	String() string
}
