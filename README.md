
<p align="center">  
  <a href="#">
    <img src="https://img.shields.io/github/go-mod/go-version/josehbez/semver">
  </a>
  <a href="LICENSE">
    <img src="https://img.shields.io/github/license/josehbez/semver?style=flat-square" />
  </a>
  <a href="semv.toml">
    <img src="https://img.shields.io/badge/semv-1.0.3-green">
  </a>
</p>
<p align="center">  
  <a href="https://github.com/josehbez/semv">
    <img src="https://img.shields.io/badge/semver.org-by semv-green">
  </a>
</p>

## Semantic Versioning

`semv` It is a tool that facilitates the manipulation of the semantic-versioning of a software project. Upon initialization, it generates a `semv.toml` file that contains the semantic-versioning data. 



## Our goals

* `semv.toml` be a standard file that stores the software semantic-versioning
* Facilitate the manipulation of semantic-versioning regardless of the software project
* Based on spec [semver.org](http://semver.org)


## Use badge
```markdown

# HEADER:  Copy and replace VERSION for current version 
[![](https://img.shields.io/badge/semv-VERSION-green)](semv.toml)

[![](https://img.shields.io/badge/semver.org-by semv-green)](https://github.com/josehbez/semv)

```

## Download and install

if you're interested  in getting the source code, or hacking on `semv`, you can intall via `go get`. 

```bash
go get -u github.com/josehbez/semv
```

If you're interested use mode production 
```bash

go install github.com/josehbez/semv/cmd/semv

```


