package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/atotto/clipboard"
	"github.com/urfave/cli/v2"
)

var subCmdList = []*cli.Command{}

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
	imageURL, err := getImageURL(cCtx.String("tag"))
	if err != nil {
		return fmt.Errorf("failed to get image url: %w", err)
	}

	clipboard.WriteAll(formatForMarkdown(imageURL))

	return nil
}

func getImageURL(tag string) (string, error) {
	apiKey := os.Getenv("GIPHY_API_KEY")
	if apiKey == "" {
		return "", errors.New("empty API Key")
	}

	resp, err := http.Get(fmt.Sprintf("https://api.giphy.com/v1/gifs/random?api_key=%s&tag=%s&rating=g", apiKey, tag))
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

func formatForMarkdown(url string) string {
	return fmt.Sprintf("![LGTM](%s)", url)
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
