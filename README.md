# DESCRIPTION

Fetching packages in golang can be difficult, especially when juggling multiple packages, and private repositories. GIGO (Gigo Installer for Go) tries to solve that problem, effectively being the golang equivalent of [python's pip](https://pip.pypa.io).

The direct result of using gigo, is that a *requirements.txt* file can now be used for golang projects, making it easy to componentize golang projects.  Through using gigo, components can now have their own life cycles, and be managed, even within private repositories.

---------------------


## How Gigo Works

For all packages that can be fetched using the existing *go get* calls, Gigo acts as a wrapper.  For elements in private repositories, gigo calls out to the appropriate revision control tool, like *git*.

## Usage

Packages can be installed in several ways. You can use a go getable path, a go getable path and a commit number for versioned code, or the path, optional version, and a URL to pull the code from. If a url is used, gigo will copy the repository found at that url into the path specified.

Make sure you've set your GOPATH and then:

    gigo install github.com/BurntSushi/toml
    gigo install github.com/BurntSushi/toml github.com/kr/fs
    gigo install -r requirements.txt
    gigo uninstall github.com/BurntSushi/toml
    gigo --help

## Building Gigo

First, make sure your GOPATH has been set. On unixes, simply run *sh build.sh*. On Windows, run build.bat.

Outputs are stored in dist/gigo and dist/gigo.exe respectively.  At that point, you'll want to copy the binary produced, into a known location.

## Requirements File Format

The gigo requirements.txt file follows [python's pip](https://pip.pypa.io) as inspiration. Each line in the requirements file should represent a *gigo fetchable* item. In the case of private github repositories, you can specify a hash or tag, as well as a destination.

Hashes or tags can be specified following a hashmark. Similarly, a destination is specified by adding a comma. Below is an example of a valid requirements.txt file, using some random repositories.

Example:

    github.com/BurntSushi/toml
    git@github.com:kr/fs#2788f0
    git+ssh@bitbucket.org:liamstask/goose#8488cc4, bitbucket.org/liamstask/goose

## Future

Ideally, rcsget.go would be extended, to add support for mercurial, subversion, and pretty much anything else under the sun. We'd also like to make it easy to upgrade packages.

--------------------

# LICENSE

The MIT License (MIT)

Copyright (c) 2015 Lyrical Security Ltd

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

