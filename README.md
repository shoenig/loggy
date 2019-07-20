loggy
=====

Go's worst logging package.

[![Go Report Card](https://goreportcard.com/badge/gophers.dev/pkgs/loggy)](https://goreportcard.com/report/gophers.dev/pkgs/loggy)
[![Build Status](https://travis-ci.com/shoenig/loggy.svg?branch=master)](https://travis-ci.com/shoenig/loggy)
[![GoDoc](https://godoc.org/gophers.dev/pkgs/loggy?status.svg)](https://godoc.org/gophers.dev/pkgs/loggy)
![NetflixOSS Lifecycle](https://img.shields.io/osslifecycle/shoenig/loggy.svg)
![GitHub](https://img.shields.io/github/license/shoenig/loggy.svg)

# Project Overview

Package `loggy` provides a simple logging library. It is focused on providing
the bare minimum implementation of a leveled-logger for Go projects.

# Getting Started

The `loggy` package can be installed by running:
```
$ go get gophers.dev/pkgs/loggy
```

#### Example Usage
A source file with this block
```golang
log := New("my-thing")
log.Tracef("this is trace %d", 1)
log.Infof("this is info %d", 2)
log.Warnf("this is warn %d", 3)
log.Errorf("this is error %d", 4)
```

Will produce output
```Text
2019/01/21 09:36:32 TRACE [my-thing] this is trace 1
2019/01/21 09:36:32 INFO  [my-thing] this is info 2
2019/01/21 09:36:32 WARN  [my-thing] this is warn 3
2019/01/21 09:36:32 ERROR [my-thing] this is error 4
```

# Contributing

Package `loggy` is pretty minimalist, and so it's basically feature complete.
Bug fixes and good ideas though are always welcome, please just file an issue.

# License

The `gophers.dev/pkgs/loggy` module is open source under the [BSD-3-Clause](LICENSE) license.
