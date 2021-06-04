package main

import (
	"fmt"
	"os"
	"github.com/wuyan94zl/go-api/artisan/cli"
)

func main() {
	artisan := cli.NewArtisan(os.Args)
	if err := artisan.Run(); err != nil {
		fmt.Println("error:", err)
	}
}
