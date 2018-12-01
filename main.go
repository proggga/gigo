package main

import (
  "github.com/LyricalSecurity/gigo/actions"
  "github.com/codegangsta/cli"
  "log"
  "os"
)

func main() {

  gopath := os.Getenv("GOPATH")
  if gopath == "" {
    log.Fatal("GOPATH environment variable is not set.")
  }

  app := cli.NewApp()
  app.Name = "gigo"
  app.Version = "0.3.1"

  app.Commands = []cli.Command {
    {
      Name: "install",
      Usage: "Install packages",

      // command-line sub-flags for instal
      Flags: []cli.Flag{
        cli.StringFlag{
          Name: "r",
          Usage: "Path to the file listing golang packages to fetch",
        },
      },
      Action: actions.Install,
    },
    {
      Name: "uninstall",
      Usage: "Uninstall packages",
      Action: actions.Uninstall,
    },
    {
      Name: "list",
      Aliases: []string{"freeze"},
      Usage: "List of installed packages",
      Action: actions.List,
    },
  }

  app.Run(os.Args)

}
