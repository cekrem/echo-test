package main

import (
	"fmt"
	"golang.org/x/crypto/acme/autocert"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func getAddress() string {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "1337"
	}

	return fmt.Sprintf(":%s", port)
}

func main() {
	// Echo instance
	e := echo.New()

	// TLS
	e.AutoTLSManager.Cache = autocert.DirCache(".cache")

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)

	// Start server
	e.Logger.Fatal(e.StartAutoTLS(getAddress()))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
