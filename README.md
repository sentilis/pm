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
  * `version` 👉 semver.org

## 📒 Getting started with PM

Table of contents

* [Install](#install)
* [Run your first command](#run-your-first-command)
* [Docs](#docs)
* [Issues](#issues)
* [Discussions](#discussions)

### Install

```bash
go install github.com/josehbez/pm/cmd/pm
```

### Run your first command

```bash
# Create an empty PM
pm init 

pm version # Show version based on semver.org
pm --major / --minor / --patch # Increase major/ minor / patch version
pm version --check 1.0.1-alfa.1+exp.sha.1 # check if version is based on semver.org


pm changelog # Show changelog
pm changelog --added "sub-command changelog" # for new feature
pm changelog --fixed "Show changelog by order desc " # for any bug fixes

pm license # Show license
pm license --list # list available licenses
pm license --fetch MIT # fetch license by identifier
pm license --save --fetch MIT # fetch & save license by identifier


```

### Docs

[https://github.com/josehbez/pm/wiki](https://github.com/josehbez/pm/wiki)

### Issues

[https://github.com/josehbez/pm/issues](https://github.com/josehbez/pm/issues)

### Discussions

[https://github.com/josehbez/pm/discussions](https://github.com/josehbez/pm/discussions)