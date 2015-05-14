package actions

import (
  "io/ioutil"
  "os"
  "path"
  "path/filepath"
  "github.com/codegangsta/cli"
)

// In this context, uninstall just means
func Uninstall(c *cli.Context) {
  for _, elem := range c.Args() {

    where := path.Join("src", elem)
    // check if path exist and die
    _, e := os.Stat(where)
    if os.IsNotExist(e) {
      println(where + " is not installed, exiting.")
      os.Exit(3)
    }

    err := os.RemoveAll(where)
    if err != nil {
      println("Could not remove package", where)
      os.Exit(3)
    }
    removeParent(where)
  }
}

/** Given a directory, look to see if the parent is empty, if it is, remove that puppy
to */
func removeParent(where string) {
  parent := path.Join(where, "..")
  absolute, err := filepath.Abs(parent)
  if (err != nil) {
    return
  }

  _, err = ioutil.ReadDir(absolute)
  if (err == nil) {
    os.RemoveAll(absolute)
  }
}
