package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	app := cli.App{
		Name:    "gopl",
		Usage:   "translation on command",
		Version: "0.0.1",
	}

	app.Action = func(ctx *cli.Context) error {
		err := getApiKey()
		if err != nil {
			return err
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func getApiKey() error {
	key := os.Getenv("DEEPL_TOKEN")
	if len(key) == 0 {
		return fmt.Errorf("Error: %s", "DEEPL_TOKEN not found ")
	}
	fmt.Print(key)
	return nil
}
