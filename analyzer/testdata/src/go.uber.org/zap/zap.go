package zap

type Logger struct{}

func NewExample() *Logger {
	return &Logger{}
}

func (l *Logger) Info(msg string, fields ...any) {}
