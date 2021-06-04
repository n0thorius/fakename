package cmd

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

const (
	imageUrl = "https://thispersondoesnotexist.com/image"
	prefix   = "image"
	ext      = ".jpeg"
)

var imageCmd = &cobra.Command{
	Use:   "image",
	Short: "get new image",
	Run: func(cmd *cobra.Command, args []string) {
		image, err := getImage()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error fetching image: %v\n", err)
		}
		fmt.Printf("Image '%s' saved successfully\n", image)
	},
}

func init() {
	RootCmd.AddCommand(imageCmd)
}

func getImage() (string, error) {
	resp, err := http.Get(imageUrl)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	filename, err := generateFilename()
	if err != nil {
		return "", err
	}

	f, err := os.Create(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	if _, err := io.Copy(f, resp.Body); err != nil {
		return "", err
	}

	return filename, nil
}

func generateFilename() (string, error) {
	randBytes := make([]byte, 8)
	rand.Read(randBytes)

	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return filepath.Join(dir, prefix+hex.EncodeToString(randBytes)+ext), nil
}
