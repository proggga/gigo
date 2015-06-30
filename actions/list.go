package actions

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"github.com/codegangsta/cli"
)

//List provides list of packages in $GOPATH/src
func List(c *cli.Context) {

	srcpath := fmt.Sprintf("%s/src/", os.Getenv("GOPATH"))
	srcdir, err := ioutil.ReadDir(srcpath)
	if err != nil {
		log.Fatal(err)
	}
	if len(srcdir) == 0 {
		log.Fatal(fmt.Sprintf("%s is empty", srcpath))
	}
	for _, source := range srcdir {
		fmt.Println("Packages from " + source.Name())
		title, err := ioutil.ReadDir(srcpath + source.Name())
		if err != nil {
			log.Print(err)
		}

		for _, pack := range title {
			if pack.IsDir() {
				fmt.Println(pack.Name())
			}
		}
		fmt.Println("\n")
	}
}
