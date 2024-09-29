package middleware

import (
	"os"
	"runtime/pprof"
	"time"

	"github.com/adisnuhic/go-clean/pkg/log"
	"github.com/gin-gonic/gin"
)

// ProfileMiddleware wraps a Gin handler function to profile its execution
func ProfileMiddleware(next gin.HandlerFunc, logger log.ILogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		f, err := os.Create("/tmp/cpu.prof")
		if err != nil {
			logger.Fatalf("could not create CPU profile: %v", err)
		}
		defer f.Close()

		if err := pprof.StartCPUProfile(f); err != nil {
			logger.Fatalf("could not start CPU profile: %v", err)
		}
		defer pprof.StopCPUProfile()

		start := time.Now()
		next(c) // Call the next handler
		logger.Printf("Profiling duration for %s: %v", c.FullPath(), time.Since(start))
	}
}
