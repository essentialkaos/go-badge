<p align="center"><a href="#readme"><img src="https://gh.kaos.st/go-badge.svg"/></a></p>

<p align="center">
  <a href="https://kaos.sh/g/go-badge.v1"><img src="https://gh.kaos.st/godoc.svg" alt="PkgGoDev" /></a>
  <a href="https://kaos.sh/w/go-badge/ci"><img src="https://kaos.sh/w/go-badge/ci.svg" alt="GitHub Actions CI Status" /></a>
  <a href="https://kaos.sh/r/go-badge"><img src="https://kaos.sh/r/go-badge.svg" alt="GoReportCard" /></a>
  <a href="https://kaos.sh/c/go-badge"><img src="https://kaos.sh/c/go-badge.svg" alt="Coverage Status" /></a>
  <a href="https://kaos.sh/b/go-badge"><img src="https://kaos.sh/b/0cbded00-9dfb-458b-bdf8-27b2c70ede9b.svg" alt="Codebeat badge" /></a>
  <a href="https://kaos.sh/w/go-badge/codeql"><img src="https://kaos.sh/w/go-badge/codeql.svg" alt="GitHub Actions CodeQL Status" /></a>
  <a href="#license"><img src="https://gh.kaos.st/apache2.svg"></a>
</p>

<p align="center"><a href="#installation">Installation</a> • <a href="#usage-example">Usage example</a> • <a href="#examples">Examples</a> • <a href="#build-status">Build Status</a> • <a href="#contributing">Contributing</a> • <a href="#thanks">Thanks</a> • <a href="#license">License</a></p>

<br/>

`badge` is a Go package for generating SVG badges.

### Installation

Make sure you have a working Go 1.15+ workspace (_[instructions](https://golang.org/doc/install)_), then:

````bash
go get -d pkg.re/essentialkaos/go-badge.v1
````

For update to latest stable release, do:

```bash
go get -d -u pkg.re/essentialkaos/go-badge.v1
```

### Usage example

```go
package main

// ////////////////////////////////////////////////////////////////////////// //

import (
  "fmt"
  "pkg.re/essentialkaos/go-badge.v1"
)

// ////////////////////////////////////////////////////////////////////////// //

func main() {
  g, err := badge.NewGenerator("Verdana.ttf", 11)

  if err != nil {
    panic(err)
  }

  fmt.Println(string(g.GeneratePlastic("status", "ok", "#97ca00")))
}
```

### Examples

| Flat | Flat Square | Plastic |
|------|-------------|---------|
| [flat](.github/images/flat_brightgreen.svg) | [square](.github/images/square_brightgreen.svg) | [plastic](.github/images/plastic_brightgreen.svg) |
| [flat](.github/images/flat_green.svg) | [square](.github/images/square_green.svg) | [plastic](.github/images/plastic_green.svg) |
| [flat](.github/images/flat_yellowgreen.svg) | [square](.github/images/square_yellowgreen.svg) | [plastic](.github/images/plastic_yellowgreen.svg) |
| [flat](.github/images/flat_yellow.svg) | [square](.github/images/square_yellow.svg) | [plastic](.github/images/plastic_yellow.svg) |
| [flat](.github/images/flat_orange.svg) | [square](.github/images/square_orange.svg) | [plastic](.github/images/plastic_orange.svg) |
| [flat](.github/images/flat_red.svg) | [square](.github/images/square_red.svg) | [plastic](.github/images/plastic_red.svg) |
| [flat](.github/images/flat_blue.svg) | [square](.github/images/square_blue.svg) | [plastic](.github/images/plastic_blue.svg) |
| [flat](.github/images/flat_lightgrey.svg) | [square](.github/images/square_lightgrey.svg) | [plastic](.github/images/plastic_lightgrey.svg) |
| [flat](.github/images/flat_success.svg) | [square](.github/images/square_success.svg) | [plastic](.github/images/plastic_success.svg) |
| [flat](.github/images/flat_important.svg) | [square](.github/images/square_important.svg) | [plastic](.github/images/plastic_important.svg) |
| [flat](.github/images/flat_critical.svg) | [square](.github/images/square_critical.svg) | [plastic](.github/images/plastic_critical.svg) |
| [flat](.github/images/flat_informational.svg) | [square](.github/images/square_informational.svg) | [plastic](.github/images/plastic_informational.svg) |
| [flat](.github/images/flat_inactive.svg) | [square](.github/images/square_inactive.svg) | [plastic](.github/images/plastic_inactive.svg) |
| [flat](.github/images/flat_custom.svg) | [square](.github/images/square_custom.svg) | [plastic](.github/images/plastic_custom.svg) |

### Build Status

| Branch | Status |
|--------|----------|
| `master` | [![CI](https://kaos.sh/w/go-badge/ci.svg?branch=master)](https://kaos.sh/w/go-badge/ci?query=branch:master) |
| `develop` | [![CI](https://kaos.sh/w/go-badge/ci.svg?branch=develop)](https://kaos.sh/w/go-badge/ci?query=branch:develop) |

### Contributing

Before contributing to this project please read our [Contributing Guidelines](https://github.com/essentialkaos/contributing-guidelines#contributing-guidelines).

### Thanks

We would like to thank:

* All authors and [contributors](https://github.com/badges/shields/graphs/contributors) of [shields.io](https://shields.io) service;
* All authors of [`freetype`](https://github.com/golang/freetype/blob/master/AUTHORS) package;
* [@narqo](https://github.com/narqo) for [`go-badge`](https://github.com/narqo/go-badge) package.

### License

[Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0)

<p align="center"><a href="https://essentialkaos.com"><img src="https://gh.kaos.st/ekgh.svg"/></a></p>
