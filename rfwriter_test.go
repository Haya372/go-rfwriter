package rfwriter

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

func ExampleRFWriter() {
	w, err := NewRFWriter("example.txt")
	if err != nil {
		os.Exit(1)
	}
	defer w.Close()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGUSR1)
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT)
	defer stop()
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-c:
				w.Reload()
			}
		}
	}()

	// main process with rfwriter
	// ...
}
