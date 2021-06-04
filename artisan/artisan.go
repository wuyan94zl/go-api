package main

import (
	"fmt"
	"github.com/wuyan94zl/go-api/artisan/cli"
	"os"
)

func main() {
	artisan := cli.NewArtisan(os.Args)
	if err := artisan.Run(); err != nil {
		fmt.Println("error:", err)
	}
}
