// rfwriter is Reloadable File Writer for Golang.
package rfwriter

import (
	"io"
	"os"
	"sync"
)

/*
RFWriter is the interface that provides io.WriteCloser methods and Reload().
*/
type RFWriter interface {
	Write(p []byte) (n int, err error)
	Reload() error
	Close() error
}

type rfwriterImpl struct {
	io.WriteCloser
	filePath string
	mu       sync.Mutex
}

func (w *rfwriterImpl) Reload() error {
	w.mu.Lock()
	defer w.mu.Unlock()
	file, err := os.OpenFile(w.filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	w.WriteCloser = file
	return nil
}

// Create RFWriter instance
func NewRFWriter(filePath string) (RFWriter, error) {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}
	return &rfwriterImpl{
		WriteCloser: file,
		filePath:    filePath,
	}, nil
}
