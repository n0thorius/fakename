package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var postCmd = &cobra.Command{
	Use:   "post",
	Short: "post to ghostbin and pastebin",
	Run: func(cmd *cobra.Command, args []string) {
		if len(os.Args) < 1 {
			fmt.Fprintf(os.Stderr, "You need to provide filename")
			os.Exit(1)
		}

		url, err := saveGhostbin(args[0])
		if err != nil {
			panic(err)
		}
		fmt.Println(url)
	},
}

func init() {
	RootCmd.AddCommand(postCmd)
}

func saveGhostbin(filename string) (string, error) {
	content, _ := ioutil.ReadFile(filename)

	val := url.Values{
		"text":     {string(content)},
		"title":    {""},
		"password": {""},
	}

	req, _ := http.NewRequest("POST", "https://ghostbin.com/paste/new", strings.NewReader(val.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return errors.New("Redirect")
	}
	resp, err := client.Do(req)
	if err != nil {
		if resp.StatusCode == 302 {
			return resp.Header.Get("Location"), nil
		}
	}
	resp.Body.Close()
	return "", err
}
