package main

import (
	"fmt"

	"github.com/alexflint/go-arg"
	"github.com/gin-gonic/gin"
	"github.com/nathan-wien/qkly/internal/companion"
	"github.com/nathan-wien/qkly/internal/filesystem"
	"github.com/nathan-wien/qkly/internal/workspace"
)

func main() {
	fmt.Println("qkly - Competive Programing Utilities")

	var args struct {
		Mode string `arg:"positional"`
	}

	arg.MustParse(&args)

	if args.Mode == "" {
		fmt.Println("Usage: qkly <mode>. Please run `qkly -h` for more info")
	} else if args.Mode == "fetch" {
		// Load configuration
		conf, err := workspace.NewConfig(filesystem.Mgr().Fs)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("Config file found")
		fmt.Printf("Listening on port %s\n", conf.Port())

		gin.SetMode(gin.ReleaseMode)
		engine := gin.Default()
		engine.SetTrustedProxies(nil)
		engine.POST("/", companion.Handle)
		engine.Run(fmt.Sprintf("localhost:%s", conf.Port()))
	}
}
