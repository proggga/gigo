// Copyright (c) 2014 Sergey Romanov <xxsmotur@gmail.com>
// Copyright (c) 2018 Mikhail Fesenko <proggga@gmail.com>
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
        if source.Name() == "github.com" {
            srcdir2, err2 := ioutil.ReadDir(srcpath + source.Name())
            if err2 != nil {
                log.Fatal(err2)
            }
            for _, pack := range srcdir2 {
                if pack.IsDir() {
                    srcdir3, err3 := ioutil.ReadDir(srcpath + source.Name() + "/" + pack.Name())
                    if err3 != nil {
                        log.Fatal(err3)
                    }
                    for _, pack2 := range srcdir3 {
                        if pack.IsDir() {
                            fmt.Println("github.com/" + pack.Name() + "/" + pack2.Name())
                        }
                    }
                }
            }
        } else {
            title, err := ioutil.ReadDir(srcpath + source.Name())
            if err != nil {
                log.Print(err)
            }

            for _, pack := range title {
                if pack.IsDir() {
                    fmt.Println(source.Name() + "/" + pack.Name())
                }
            }
            fmt.Println("\n")
        }
	}
}
