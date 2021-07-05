<p align="center">  
    <a href="#">
        <img src="https://img.shields.io/github/go-mod/go-version/josehbez/pm">
    </a>
    <a href="LICENSE">
        <img src="https://img.shields.io/github/license/josehbez/pm?style=flat-square" />
    </a>
    <a href=".pm/version.yml">
        <img src="https://img.shields.io/badge/dynamic/yaml?color=green&label=version&query=version.*&url=https://raw.githubusercontent.com/josehbez/pm/master/.pm/version.yml">
    </a>
    <a href=".pm/version.yml">
        <img src="https://img.shields.io/badge/dynamic/yaml?color=green&label=prerelease&query=prerelease.*&url=https://raw.githubusercontent.com/josehbez/pm/master/.pm/version.yml"/>
    </a>
    <a href=".pm/version.yml">
        <img src="https://img.shields.io/badge/dynamic/yaml?color=green&label=build&query=build.*&url=https://raw.githubusercontent.com/josehbez/pm/master/.pm/version.yml"/>
    </a>
</p>

## 📦 Project Metadata Management

What's PM?

It is a command line tool that initializes a folder called `.pm`,
whose objective is to store all the metadata files of the project(license, authors, changelog, sponsors, semantic-versioning ... ).
  
## 🎯 Our goals

* The `.pm` folder in a standard location to store project metadata
* Facilitate the manipulation metadata, based on spec such as:
  * `license` 👉 https://spdx.org/licenses
  * `changelog` 👉 https://keepachangelog.com
  * `version` 👉 https://semver.org

## 📒 Getting started with PM

Table of contents

* [Install](#install)
* [Quick help](#quick-help)
* [Run your first command](#run-your-first-command)
* [Docs](#docs)
* [Issues](#issues)
* [Discussions](#discussions)

### Install

```bash
go get -u github.com/josehbez/pm/cmd/pm
go install github.com/josehbez/pm/cmd/pm
```

### Quick Help
```bash
Usage:
  pm [command]

Available Commands:
  author      Show & add authors
  changelog   Show & add changelog
  help        Help about any command
  init        Create an empty PM
  license     Show & add license
  maintainer  Show & add maintainers
  version     Semantic versioning management

Flags:
  -h, --help   help for pm

Use "pm [command] --help" for more information about a command.
```

### Run your first command

```bash
# Create an empty PM
pm init 

pm version # Show version based on semver.org
pm version --major / --minor / --patch # Increase major/ minor / patch version
pm version --check 1.0.1-alfa.1+exp.sha.1 # check if version is based on semver.org

pm version prerelease # Show version-prerelease
pm version prerelease --label alfa # Add label pre-release
pm version prerelease --major # Increase major pre-realease

pm version build # Show version+build
pm version build --label exp.sha # Add label build
pm version build --major # Increase major build

pm changelog # Show changelog
pm changelog --added "sub-command changelog" # for new feature
pm changelog --fixed "Show changelog by order desc " # for any bug fixes

pm license # Show license
pm license --list # list available licenses
pm license --fetch MIT # fetch license by identifier
pm license --fetch MIT --save # fetch & save license by identifier

pm author # Show  authors
pm author --add "Jose Hbez" https://github.com/josehbez

pm maintainer # Show  maintainers
pm maintainer --add "Jose Hbez" https://github.com/josehbez

```

### Docs

[https://github.com/josehbez/pm/wiki](https://github.com/josehbez/pm/wiki)

### Issues

[https://github.com/josehbez/pm/issues](https://github.com/josehbez/pm/issues)

### Discussions

[https://github.com/josehbez/pm/discussions](https://github.com/josehbez/pm/discussions)


### Troubleshooting

#### Cannot find package
Pm installation seems to fail a lot lately with the following (or a similar) error:
```
package github.com/hashicorp/hcl/hcl/printer: cannot find package "github.com/hashicorp/hcl/hcl/printer" in any of:
        /usr/lib/go-1.13/src/github.com/hashicorp/hcl/hcl/printer (from $GOROOT)
        /home/user/go/src/github.com/hashicorp/hcl/hcl/printer (from $GOPATH)
```
The solution is easy: switch to using Go Modules. Please refer to the wiki on how to do that.

`export GO111MODULE=on`

