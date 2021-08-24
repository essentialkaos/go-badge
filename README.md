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

Make sure you have a working Go 1.16+ workspace (_[instructions](https://golang.org/doc/install)_), then:

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

| Flat | Flat Square | Plastic | Flat (Simple) | Flat Square (Simple) | Plastic (Simple) |
|------|-------------|---------|---------------|----------------------|------------------|
| ![flat](.github/images/flat_flat_brightgreen.svg) | ![square](.github/images/square_flat_brightgreen.svg) | ![plastic](.github/images/plastic_flat_brightgreen.svg) | ![flat-simple](.github/images/flat_simple_flat_brightgreen.svg) | ![square-simple](.github/images/square_simple_flat_brightgreen.svg) | ![plastic-simple](.github/images/plastic_simple_flat_brightgreen.svg) |
| ![flat](.github/images/flat_flat_green.svg) | ![square](.github/images/square_flat_green.svg) | ![plastic](.github/images/plastic_flat_green.svg) | ![flat-simple](.github/images/flat_simple_flat_green.svg) | ![square-simple](.github/images/square_simple_flat_green.svg) | ![plastic-simple](.github/images/plastic_simple_flat_green.svg) |
| ![flat](.github/images/flat_flat_yellowgreen.svg) | ![square](.github/images/square_flat_yellowgreen.svg) | ![plastic](.github/images/plastic_flat_yellowgreen.svg) | ![flat-simple](.github/images/flat_simple_flat_yellowgreen.svg) | ![square-simple](.github/images/square_simple_flat_yellowgreen.svg) | ![plastic-simple](.github/images/plastic_simple_flat_yellowgreen.svg) |
| ![flat](.github/images/flat_flat_yellow.svg) | ![square](.github/images/square_flat_yellow.svg) | ![plastic](.github/images/plastic_flat_yellow.svg) | ![flat-simple](.github/images/flat_simple_flat_yellow.svg) | ![square-simple](.github/images/square_simple_flat_yellow.svg) | ![plastic-simple](.github/images/plastic_simple_flat_yellow.svg) |
| ![flat](.github/images/flat_flat_orange.svg) | ![square](.github/images/square_flat_orange.svg) | ![plastic](.github/images/plastic_flat_orange.svg) | ![flat-simple](.github/images/flat_simple_flat_orange.svg) | ![square-simple](.github/images/square_simple_flat_orange.svg) | ![plastic-simple](.github/images/plastic_simple_flat_orange.svg) |
| ![flat](.github/images/flat_flat_red.svg) | ![square](.github/images/square_flat_red.svg) | ![plastic](.github/images/plastic_flat_red.svg) | ![flat-simple](.github/images/flat_simple_flat_red.svg) | ![square-simple](.github/images/square_simple_flat_red.svg) | ![plastic-simple](.github/images/plastic_simple_flat_red.svg) |
| ![flat](.github/images/flat_flat_blue.svg) | ![square](.github/images/square_flat_blue.svg) | ![plastic](.github/images/plastic_flat_blue.svg) | ![flat-simple](.github/images/flat_simple_flat_blue.svg) | ![square-simple](.github/images/square_simple_flat_blue.svg) | ![plastic-simple](.github/images/plastic_simple_flat_blue.svg) |
| ![flat](.github/images/flat_flat_lightgrey.svg) | ![square](.github/images/square_flat_lightgrey.svg) | ![plastic](.github/images/plastic_flat_lightgrey.svg) | ![flat-simple](.github/images/flat_simple_flat_lightgrey.svg) | ![square-simple](.github/images/square_simple_flat_lightgrey.svg) | ![plastic-simple](.github/images/plastic_simple_flat_lightgrey.svg) |
| ![flat](.github/images/flat_flat_success.svg) | ![square](.github/images/square_flat_success.svg) | ![plastic](.github/images/plastic_flat_success.svg) | ![flat-simple](.github/images/flat_simple_flat_success.svg) | ![square-simple](.github/images/square_simple_flat_success.svg) | ![plastic-simple](.github/images/plastic_simple_flat_success.svg) |
| ![flat](.github/images/flat_flat_important.svg) | ![square](.github/images/square_flat_important.svg) | ![plastic](.github/images/plastic_flat_important.svg) | ![flat-simple](.github/images/flat_simple_flat_important.svg) | ![square-simple](.github/images/square_simple_flat_important.svg) | ![plastic-simple](.github/images/plastic_simple_flat_important.svg) |
| ![flat](.github/images/flat_flat_critical.svg) | ![square](.github/images/square_flat_critical.svg) | ![plastic](.github/images/plastic_flat_critical.svg) | ![flat-simple](.github/images/flat_simple_flat_critical.svg) | ![square-simple](.github/images/square_simple_flat_critical.svg) | ![plastic-simple](.github/images/plastic_simple_flat_critical.svg) |
| ![flat](.github/images/flat_flat_informational.svg) | ![square](.github/images/square_flat_informational.svg) | ![plastic](.github/images/plastic_flat_informational.svg) | ![flat-simple](.github/images/flat_simple_flat_informational.svg) | ![square-simple](.github/images/square_simple_flat_informational.svg) | ![plastic-simple](.github/images/plastic_simple_flat_informational.svg) |
| ![flat](.github/images/flat_flat_inactive.svg) | ![square](.github/images/square_flat_inactive.svg) | ![plastic](.github/images/plastic_flat_inactive.svg) | ![flat-simple](.github/images/flat_simple_flat_inactive.svg) | ![square-simple](.github/images/square_simple_flat_inactive.svg) | ![plastic-simple](.github/images/plastic_simple_flat_inactive.svg) |
| ![flat](.github/images/flat_flat_custom.svg) | ![square](.github/images/square_flat_custom.svg) | ![plastic](.github/images/plastic_flat_custom.svg) | ![flat-simple](.github/images/flat_simple_flat_custom.svg) | ![square-simple](.github/images/square_simple_flat_custom.svg) | ![plastic-simple](.github/images/plastic_simple_flat_custom.svg) |
| ![flat](.github/images/flat_flat_japanese.svg) | ![square](.github/images/square_flat_japanese.svg) | ![plastic](.github/images/plastic_flat_japanese.svg) | ![flat-simple](.github/images/flat_simple_flat_japanese.svg) | ![square-simple](.github/images/square_simple_flat_japanese.svg) | ![plastic-simple](.github/images/plastic_simple_flat_japanese.svg) |

_All badges are generated with the latest version of the package._

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
