package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Handler function for Echo
func handler2(c echo.Context) error {
	time.Sleep(2100 * time.Millisecond) // Simulating delay
	return c.String(http.StatusOK, "Request processed successfully")
}

func startServer2() {
	e := echo.New()

	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: 2 * time.Second, // 2 second timeout for handler
	}))

	e.GET("/", handler2)

	server := &http.Server{
		Addr:              ":8080",
		ReadHeaderTimeout: 500 * time.Millisecond,
		ReadTimeout:       500 * time.Millisecond, // entire request
	}

	fmt.Println("Starting server on :8080...")
	if err := e.StartServer(server); err != nil {
		fmt.Println("Server error:", err)
	}
}
