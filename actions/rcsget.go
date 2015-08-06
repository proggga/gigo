package actions

import (
  "os"
  "os/exec"
  "path"
  "strings"
)

/** Fetch the specific
*/
func RcsGet(rcsurl string) {
  parts := strings.Split(rcsurl, ",")

  var src string
  var dest string

  if len(parts) > 2 {
    println("There can only be one comma on a line")
    os.Exit(3)
  }

  if len(parts) == 2 {
    src = parts[0]
    dest = path.Join("src", strings.TrimSpace(parts[1]))
  } else {
    src = rcsurl
    _dest := strings.LastIndex(src, "/")
    _git_index := strings.LastIndex(src, ".git")
    dest = path.Join("src", strings.TrimSpace(src[_dest:_git_index]))
  }

  // for the future, let's support many
  switch true {
    case strings.Index(src, "git@") != -1:
      gitGet(src, dest)
    case strings.Index(src, "git+ssh://") != -1:
      gitGet(src, dest)
    case strings.Index(src, ".git") != -1:
      gitGet(src, dest)
  }

}

/** Fetch the item from git.
First, we clone the repo, then we roll it to the revision */
func gitGet(srcurl string, dest string) {
  hash := strings.Index(srcurl, "#")
  _dest := strings.TrimSpace(dest)
  var c *exec.Cmd

  _, e := os.Stat(_dest)
  if e == nil {
    // FUTURE
    // println(_dest + " already exists, to upgrade pass in --upgrade")

    // NOW
    println(_dest + " already exists, continuing.")
    return
  }


  if hash == -1 {
    c = exec.Command("git", "clone", srcurl, _dest)
  } else {
    c = exec.Command("git", "clone", srcurl[:hash], _dest)
  }

  c.Stdout = os.Stdout
  c.Stderr = os.Stderr
  err := c.Run()
  if err != nil {
    os.Exit(1)
  }

  if hash == -1 {
    return
  }

  last_slash := strings.LastIndex(srcurl, "/")
  var reponame string
  if hash == -1 {
    reponame = srcurl[last_slash:]
  } else {
    reponame = srcurl[last_slash:hash]
  }

  // this is where we reset to the named hash
  here, wd_err := os.Getwd()
  if (wd_err != nil) {
    println("Error getting current working directory.")
    os.Exit(3)
  }
  os.Chdir(path.Join(_dest, reponame))
  c = exec.Command("git", "checkout", srcurl[hash+1:])

  hash_err := c.Run()
  if hash_err != nil {
    os.Exit(1)
  }
  os.Chdir(here)

}
