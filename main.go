package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.App{
		Name:    "gopl",
		Usage:   "translation on command",
		Version: "0.0.1",
	}

	app.Action = func(ctx *cli.Context) error {
		fmt.Print("Welcome goeepL")
		fmt.Print(ctx.Args().Get(0))
		return nil
	}

	app.Run(os.Args)
}

//func getApiKey() string {
//	key := os.Getenv("DEEPL_API_KEY")
//	if key == "" {
//		return nil
//	}
//}
