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

<p align="center"><a href="#installation">Installation</a> • <a href="#usage-example">Usage example</a> • <a href="#build-status">Build Status</a> • <a href="#contributing">Contributing</a> • <a href="#license">License</a></p>

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
  g, err := badge.NewGenerator("Verdana.ttf")

  if err != nil {
    panic(err)
  }

  fmt.Println(g.GeneratePlastic("status", "ok", "#97ca00"))
}
```

### Build Status

| Branch | Status |
|--------|----------|
| `master` | [![CI](https://kaos.sh/w/go-badge/ci.svg?branch=master)](https://kaos.sh/w/go-badge/ci?query=branch:master) |
| `develop` | [![CI](https://kaos.sh/w/go-badge/ci.svg?branch=develop)](https://kaos.sh/w/go-badge/ci?query=branch:develop) |

### Contributing

Before contributing to this project please read our [Contributing Guidelines](https://github.com/essentialkaos/contributing-guidelines#contributing-guidelines).

### License

[Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0)

<p align="center"><a href="https://essentialkaos.com"><img src="https://gh.kaos.st/ekgh.svg"/></a></p>
