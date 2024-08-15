package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/atotto/clipboard"
	"github.com/urfave/cli/v2"
)

func init() {
	subCmdList = append(subCmdList, &cli.Command{
		Name:  "giphy-random",
		Usage: "get randomly giphy image url",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "tag",
				Value: "lgtm",
				Usage: "search keyword",
			},
		},
		Action: (&giphyRandomCmd{}).action,
	})
}

type giphyRandomCmd struct{}

func (c *giphyRandomCmd) action(cCtx *cli.Context) error {
	return c.run(os.Stdout, cCtx.String("tag"))
}

func (c *giphyRandomCmd) run(out io.Writer, tag string) error {
	imageURL, err := c.getImageURL(tag)
	if err != nil {
		return fmt.Errorf("failed to get image url: %w", err)
	}

	clipboard.WriteAll(c.formatForMarkdown(imageURL))

	if _, err := out.Write([]byte("copied!\n")); err != nil {
		return fmt.Errorf("failed to write output: %w", err)
	}

	return nil
}

func (c *giphyRandomCmd) getImageURL(tag string) (string, error) {
	apiKey := os.Getenv("GIPHY_API_KEY")
	if apiKey == "" {
		return "", errors.New("empty API Key")
	}

	v := url.Values{}
	v.Add("api_key", apiKey)
	v.Add("tag", tag)
	v.Add("rating", "g")

	resp, err := http.Get(fmt.Sprintf("https://api.giphy.com/v1/gifs/random?%s", v.Encode()))
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

func (c *giphyRandomCmd) formatForMarkdown(url string) string {
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
