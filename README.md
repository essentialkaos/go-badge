<p align="center"><a href="#readme"><img src=".github/images/card.svg"/></a></p>

<p align="center">
  <a href="https://kaos.sh/g/go-badge"><img src=".github/images/godoc.svg"/></a>
  <a href="https://kaos.sh/w/go-badge/ci"><img src="https://kaos.sh/w/go-badge/ci.svg" alt="GitHub Actions CI Status" /></a>
  <a href="https://kaos.sh/r/go-badge"><img src="https://kaos.sh/r/go-badge.svg" alt="GoReportCard" /></a>
  <a href="https://kaos.sh/c/go-badge"><img src="https://kaos.sh/c/go-badge.svg" alt="Coverage Status" /></a><br/>
  <a href="https://kaos.sh/b/go-badge"><img src="https://kaos.sh/b/0cbded00-9dfb-458b-bdf8-27b2c70ede9b.svg" alt="Codebeat badge" /></a>
  <a href="https://kaos.sh/y/go-badge"><img src="https://kaos.sh/y/a090e4e0d8e14e58bc9c7a5458c2803e.svg" alt="Codacy badge" /></a>
  <a href="https://kaos.sh/w/go-badge/codeql"><img src="https://kaos.sh/w/go-badge/codeql.svg" alt="GitHub Actions CodeQL Status" /></a>
  <a href="#license"><img src=".github/images/license.svg"/></a>
</p>

<p align="center"><a href="#usage-example">Usage example</a> • <a href="#examples">Examples</a> • <a href="#ci-status">CI Status</a> • <a href="#contributing">Contributing</a> • <a href="#thanks">Thanks</a> • <a href="#license">License</a></p>

<br/>

`badge` is a Go package for generating SVG badges.

### Usage example

```go
package main

// ////////////////////////////////////////////////////////////////////////// //

import (
  "fmt"
  "github.com/essentialkaos/go-badge"
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
| ![flat](.github/images/flat_brightgreen.svg) | ![square](.github/images/square_brightgreen.svg) | ![plastic](.github/images/plastic_brightgreen.svg) | ![flat-simple](.github/images/flat_simple_brightgreen.svg) | ![square-simple](.github/images/square_simple_brightgreen.svg) | ![plastic-simple](.github/images/plastic_simple_brightgreen.svg) |
| ![flat](.github/images/flat_green.svg) | ![square](.github/images/square_green.svg) | ![plastic](.github/images/plastic_green.svg) | ![flat-simple](.github/images/flat_simple_green.svg) | ![square-simple](.github/images/square_simple_green.svg) | ![plastic-simple](.github/images/plastic_simple_green.svg) |
| ![flat](.github/images/flat_yellowgreen.svg) | ![square](.github/images/square_yellowgreen.svg) | ![plastic](.github/images/plastic_yellowgreen.svg) | ![flat-simple](.github/images/flat_simple_yellowgreen.svg) | ![square-simple](.github/images/square_simple_yellowgreen.svg) | ![plastic-simple](.github/images/plastic_simple_yellowgreen.svg) |
| ![flat](.github/images/flat_yellow.svg) | ![square](.github/images/square_yellow.svg) | ![plastic](.github/images/plastic_yellow.svg) | ![flat-simple](.github/images/flat_simple_yellow.svg) | ![square-simple](.github/images/square_simple_yellow.svg) | ![plastic-simple](.github/images/plastic_simple_yellow.svg) |
| ![flat](.github/images/flat_orange.svg) | ![square](.github/images/square_orange.svg) | ![plastic](.github/images/plastic_orange.svg) | ![flat-simple](.github/images/flat_simple_orange.svg) | ![square-simple](.github/images/square_simple_orange.svg) | ![plastic-simple](.github/images/plastic_simple_orange.svg) |
| ![flat](.github/images/flat_red.svg) | ![square](.github/images/square_red.svg) | ![plastic](.github/images/plastic_red.svg) | ![flat-simple](.github/images/flat_simple_red.svg) | ![square-simple](.github/images/square_simple_red.svg) | ![plastic-simple](.github/images/plastic_simple_red.svg) |
| ![flat](.github/images/flat_blue.svg) | ![square](.github/images/square_blue.svg) | ![plastic](.github/images/plastic_blue.svg) | ![flat-simple](.github/images/flat_simple_blue.svg) | ![square-simple](.github/images/square_simple_blue.svg) | ![plastic-simple](.github/images/plastic_simple_blue.svg) |
| ![flat](.github/images/flat_lightgrey.svg) | ![square](.github/images/square_lightgrey.svg) | ![plastic](.github/images/plastic_lightgrey.svg) | ![flat-simple](.github/images/flat_simple_lightgrey.svg) | ![square-simple](.github/images/square_simple_lightgrey.svg) | ![plastic-simple](.github/images/plastic_simple_lightgrey.svg) |
| ![flat](.github/images/flat_success.svg) | ![square](.github/images/square_success.svg) | ![plastic](.github/images/plastic_success.svg) | ![flat-simple](.github/images/flat_simple_success.svg) | ![square-simple](.github/images/square_simple_success.svg) | ![plastic-simple](.github/images/plastic_simple_success.svg) |
| ![flat](.github/images/flat_important.svg) | ![square](.github/images/square_important.svg) | ![plastic](.github/images/plastic_important.svg) | ![flat-simple](.github/images/flat_simple_important.svg) | ![square-simple](.github/images/square_simple_important.svg) | ![plastic-simple](.github/images/plastic_simple_important.svg) |
| ![flat](.github/images/flat_critical.svg) | ![square](.github/images/square_critical.svg) | ![plastic](.github/images/plastic_critical.svg) | ![flat-simple](.github/images/flat_simple_critical.svg) | ![square-simple](.github/images/square_simple_critical.svg) | ![plastic-simple](.github/images/plastic_simple_critical.svg) |
| ![flat](.github/images/flat_informational.svg) | ![square](.github/images/square_informational.svg) | ![plastic](.github/images/plastic_informational.svg) | ![flat-simple](.github/images/flat_simple_informational.svg) | ![square-simple](.github/images/square_simple_informational.svg) | ![plastic-simple](.github/images/plastic_simple_informational.svg) |
| ![flat](.github/images/flat_inactive.svg) | ![square](.github/images/square_inactive.svg) | ![plastic](.github/images/plastic_inactive.svg) | ![flat-simple](.github/images/flat_simple_inactive.svg) | ![square-simple](.github/images/square_simple_inactive.svg) | ![plastic-simple](.github/images/plastic_simple_inactive.svg) |
| ![flat](.github/images/flat_custom.svg) | ![square](.github/images/square_custom.svg) | ![plastic](.github/images/plastic_custom.svg) | ![flat-simple](.github/images/flat_simple_custom.svg) | ![square-simple](.github/images/square_simple_custom.svg) | ![plastic-simple](.github/images/plastic_simple_custom.svg) |
| ![flat](.github/images/flat_japanese.svg) | ![square](.github/images/square_japanese.svg) | ![plastic](.github/images/plastic_japanese.svg) | ![flat-simple](.github/images/flat_simple_japanese.svg) | ![square-simple](.github/images/square_simple_japanese.svg) | ![plastic-simple](.github/images/plastic_simple_japanese.svg) |

_All badges are generated with the latest version of the package._

### CI Status

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
