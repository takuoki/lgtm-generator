package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/urfave/cli/v2"
)

var subCmdList = []cli.Command{}

func main() {
	app := &cli.App{
		Name:  "lgtm-generator",
		Usage: "CLI for lgtm image",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "tag",
				Value: "lgtm",
				Usage: "search keyword",
			},
		},
		Action: giphyRandom,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func giphyRandom(cCtx *cli.Context) error {
	fmt.Println(cCtx.String("tag"))
	getImageURL(cCtx.String("tag"))
	return nil
}

func getImageURL(tag string) (string, error) {
	api_key := os.Getenv("GIPHY_API_KEY")
	resp, err := http.Get(fmt.Sprintf("https://api.giphy.com/v1/gifs/random?api_key=%s&tag=%s&rating=g", api_key, tag))
	if err != nil {
		return "", fmt.Errorf("failed to get image: %w", err)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	giphyResponse := GiphyResponse{}
	if err := json.Unmarshal(respBody, &giphyResponse); err != nil {
		return "", fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return giphyResponse.Data.Images.Original.URL, nil
}

type GiphyResponse struct {
	Data struct {
		Images struct {
			Original struct {
				URL string `json:"url"`
			} `json:"original"`
		} `json:"images"`
	} `json:"data"`
}
