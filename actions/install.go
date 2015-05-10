package actions

import (
	"github.com/LyricalSecurity/gigo/helpers"
	"github.com/codegangsta/cli"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func Install(c *cli.Context) {
	reqs := c.String("r")

	if reqs == "" {
		installPackages(c.Args())
		return
	}

	_, e := os.Stat(reqs)
	if os.IsNotExist(e) {
		println(reqs + " does not exist, exiting.")
		os.Exit(3)
	}

	file, err := ioutil.ReadFile(reqs)
	if err != nil {
		println(err)
		os.Exit(3)
	}

	contents := strings.Split(string(file), "\n")
	installPackages(contents)

}

// this is why GOPATH has to bet set in main.go
func installFromGoGet(url string) error {
	c := exec.Command("go", "get", url)
	err := c.Run()
	if err != nil {
		println(err)
		os.Exit(3)
	}

	return nil
}

func installPackages(packages []string) {

	for _, elem := range packages {
		if elem == "" || elem == "\n" {
			continue
		}

		println("Installing " + elem)
		if helpers.IsGoGettable(elem) {
			installFromGoGet(elem)
		} else {
			RcsGet(elem)
		}

	}

}
