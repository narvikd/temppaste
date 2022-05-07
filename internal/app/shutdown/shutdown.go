package shutdown

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"temppaste/internal/app"
	"temppaste/pkg/errorskit"
)

// Register registers the app shutdown channels and functionality
func Register(app *app.App) {
	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel
	<-c
	appShutdown(app)
}

func appShutdown(app *app.App) {
	fmt.Println() // For new line
	log.Println("Gracefully shutting down...")
	// Cleanup
	shutdownFiber(app)
	// Finished Cleanup
	log.Println("Fiber was successful shutdown.")
	os.Exit(0) // Exit without an error
}

func shutdownFiber(app *app.App) {
	err := app.Shutdown()
	if err != nil {
		log.Fatalln(errorskit.Wrap(err, "there was a problem shutting down fiber"))
	}
}
