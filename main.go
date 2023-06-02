package main

import (
	"fmt"
	"os"

	"github.com/heroku/docker-registry-client/registry"
	"github.com/opencontainers/go-digest"
)

func main() {

	url := "https://ghcr.io/"
	username := "" // anonymous
	password := "" // anonymous
	image := "splunkdlt/ethlogger"
	tag := "3.5.8"
	registry, err := registry.New(url, username, password)

	if err != nil {
		fmt.Printf("Error getting registry")
		os.Exit(1)
	}

	manifest, err := registry.ManifestV2(image, tag)

	if err != nil {
		fmt.Println("Error getting manifest")
		os.Exit(1)
	}

	digest := digest.NewDigestFromHex("sha256", manifest.Layers[0].Digest.Hex())

	registry.DownloadBlob(image, digest)
}
