package main
import (
	"github.com/wuyan94zl/api/routes"
)
func main() {
	router := routes.Register()
	router.Run(":8888")
}