/* hello.go*/
package main
import (
	
	"net/http"

	"github.com/labstack/echo/v4"

)

func hello (c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func main() {
	e := echo.New()
	e.GET("/hello", hello)
	e.Start(":8080")
}