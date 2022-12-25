package main

import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

type GoeepLResponse struct {
	Translations []Translated
}

type Translated struct {
	Text string `json:"text"`
}

func main() {
	app := cli.App{
		Name:    "gopl",
		Usage:   "translation on command",
		Version: "0.0.1",
	}

	app.Action = func(ctx *cli.Context) error {
		if len(ctx.Args()) == 0 {
			return fmt.Errorf("USAGE: %s", "goeepL [words]")
		}
		err := doTranslate(ctx.Args().Get(0))
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

// 引数に受けとった値を翻訳する
func doTranslate(str string) error {
	key := os.Getenv("DEEPL_TOKEN")
	if len(key) == 0 {
		return fmt.Errorf("Error: %s", "DEEPL_TOKEN not found ")
	}

	resp, err := doApiRequest(key, str)

	if err != nil {
		return fmt.Errorf("Error: %s", "resp")
	}

	result, error := parse(resp)
	if error != nil {
		return fmt.Errorf("Error: %s", "result")
	}

	fmt.Println(result.Translations[0].Text)
	return nil
}

func doApiRequest(key string, str string) (*http.Response, error) {
	path := "https://api-free.deepl.com/v2/translate"
	params := url.Values{}
	params.Add("auth_key", key)
	params.Add("source_lang", "EN")
	params.Add("target_lang", "JA")
	params.Add("text", str)

	resp, err := http.PostForm(path, params)

	if err != nil {
		return nil, fmt.Errorf("Error: %s", "resp")
	}

	return resp, nil
}

func parse(resp *http.Response) (GoeepLResponse, error) {
	var responseJson GoeepLResponse
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		err := fmt.Errorf("error")
		return responseJson, err
	}

	err = json.Unmarshal(body, &responseJson)
	if err != nil {
		err := fmt.Errorf("error")
		return responseJson, err
	}
	return responseJson, nil
}
