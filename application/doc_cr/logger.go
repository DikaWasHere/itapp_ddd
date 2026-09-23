package doc_cr
 
// Logger adalah interface untuk logging di application layer
type Logger interface {
    Info(msg string, args ...interface{})
    Error(msg string, err error, args ...interface{})
    Warn(msg string, args ...interface{})
}