package main
 
import (
	"github.com/labstack/echo"
	"net/http"
)
 
func main() {
	e := echo.New()
	e.GET("/", getEmps)
   e.Logger.Fatal(e.Start(":80"))
}
func getEmps(ctx echo.Context) error {

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"Employees": "hellow world",
	})
}
