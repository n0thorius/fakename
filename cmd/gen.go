package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/n0thorius/fakename/models"

	"github.com/spf13/cobra"
)

const (
	generatorUrl = "https://www.fakenamegenerator.com/"
)

var (
	regexes map[string]*regexp.Regexp
	results map[string]string
)

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "generate fake identity",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Fprintf(os.Stderr, "You need to provide name set and country code\n")
			os.Exit(1)
		}

		if err := checkVals(args[0], args[1]); err != nil {
			fmt.Fprintf(os.Stderr, "Wrong codes: %v\n", err)
			os.Exit(1)
		}

		regexes = make(map[string]*regexp.Regexp)
		results = make(map[string]string)

		addRegexes()

		url := fmt.Sprintf("%sgen-random-%s-%s.php", generatorUrl, args[0], args[1])

		client := &http.Client{
			Timeout: 3 * time.Second,
		}

		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Add("User-Agent", "curl/7.64.1")

		resp, err := client.Do(req)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error fetching data: %v\n", err)
			os.Exit(1)
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading response: %v\n", err)
			os.Exit(1)
		}

		for key, re := range regexes {
			match := re.FindStringSubmatch(string(body))
			if len(match) > 0 {
				result := match[1]
				if key == "addr" {
					result = strings.TrimLeft(fmt.Sprintf("%s, %s", match[1], match[2]), " ")
				}
				results[key] = result
			}
		}

		user := models.FakeUser{
			Name:     results["name"],
			Address:  results["addr"],
			SSN:      results["ssn"],
			Phone:    results["phone"],
			Birthday: results["birthday"],
			Email:    results["email"],
			Username: results["username"],
			Password: results["password"],
			Height:   results["height"],
			Weight:   results["weight"],
		}

		user.Save()

		dl, _ := cmd.Flags().GetBool("download")
		if dl {
			filename, err := getImage()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error downloading image: %v\n", err)
				os.Exit(1)
			}
			fmt.Printf("\n\nImage '%s' successfully downloaded\n", filename)
		}
	},
}

func init() {
	genCmd.Flags().BoolP("download", "d", false, "download image")
	RootCmd.AddCommand(genCmd)
}

func addRegexes() {
	regexes["name"] = regexp.MustCompile(`class="address">\n.*<h3>(.*)</h3>`)
	regexes["addr"] = regexp.MustCompile(`class="adr">\n(.*)<br.*/>(.*)</div>`)
	regexes["ssn"] = regexp.MustCompile(`<dt>SSN</dt><dd>(.*).<div`)
	regexes["phone"] = regexp.MustCompile(`<dt>Phone</dt>\n.*<dd>(.*)</dd>`)
	regexes["birthday"] = regexp.MustCompile(`<dt>Birthday</dt>\n.*<dd>(.*)</dd>`)
	regexes["email"] = regexp.MustCompile(`<dt>Email Address</dt>\n\n.*<dd>(.*)<div`)
	regexes["username"] = regexp.MustCompile(`<dt>Username</dt>\n.*<dd>(.*)</dd>`)
	regexes["password"] = regexp.MustCompile(`<dt>Password</dt>\n.*<dd>(.*)</dd>`)
	regexes["height"] = regexp.MustCompile(`<dt>Height</dt>\n.*<dd>(.*)</dd>`)
	regexes["weight"] = regexp.MustCompile(`<dt>Weight</dt>\n.*<dd>(.*)</dd>`)
}
