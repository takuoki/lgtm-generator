package main_test

import (
	"bytes"
	"os"
	"testing"

	main "github.com/takuoki/lgtm-generator/cmd/lgtm-generator"
)

func TestGiphyRandomRun(t *testing.T) {

	// TODO: giphy API をモックサーバー化して、認証通さなくても良いようにする
	os.Setenv("GIPHY_API_KEY", "dummy")

	buf := new(bytes.Buffer)
	err := main.Run(buf, "lgtm")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if buf.String() != "copied!\n" {
		t.Fatalf("unexpected output: %s", buf.String())
	}
}
