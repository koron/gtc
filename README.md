# Go Tool Catalog

[![GoDoc](https://godoc.org/github.com/koron/gtc?status.svg)](https://godoc.org/github.com/koron/gtc)
[![Actions/Go](https://github.com/koron/gtc/workflows/Go/badge.svg)](https://github.com/koron/gtc/actions?query=workflow%3AGo)
[![Go Report Card](https://goreportcard.com/badge/github.com/koron/gtc)](https://goreportcard.com/report/github.com/koron/gtc)

Gtc installs and updates well known tools written by golang.

## Install gtc

```console
$ go get -u github.com/koron/gtc
```

## Usages


```console
# List tools installed
$ go list

# List tools not-installed
$ go list -filter notinstalled

# List tools unknown (for gtc)
$ go list -filter unknown

# Install a tool
$ go install jvgrep

# Install multiple tools
$ go install goimports golint jvgrep

# Update a tool
$ go update jvgrep

# Update multiple tools
$ go update goimports golint jvgrep

# Update all tools which has been installed and not touched over 5 days
$ go update -all
```

## Customize with your own catalog


Create a new repository `github.com/{YOURNAME}/mygtc` with main.go like this:

```go
package main

import (
	"fmt"
	"os"

	"github.com/koron/gtc/gtcore"
)

func main() {
	err := gtcore.DefaultCatalog.Merge([]gtcore.Tool{
		{
			Path: "github.com/{YOURNAME}/mygtc",
			Desc: "My own go tools catalog",
		},
		// FIXME: add your favorite tools at here!
	}...).Run(os.Args)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
```

And push it to github, then install it:

```console
$ go get -u github.com/{YOURNAME}/mygtc
```

Now you can run `mygtc` instead of `gtc`.

## Customize with JSON

To load your own catalog from a file, prepare a JSON file like below.
And set its filename to `GTC_CATALOG_FILE` environment variable.
It will be merged with the default catalog.
But tools in the default catalog overrides same name tools in your catalog.

If you consider to manage your own catalog with git,
you should manage it in [golang - main.go](#customize-with-your-own-catalog)
instead of JSON.
It can override the default catalog entirely.

```json
[
  {
    "path": "github.com/foo/foo",
    "desc": "your favorite foo"
  },
  {
    "path": "github.com/foo/bar",
    "desc": "your favorite bar"
  },
  ...(other your favorite tools)...
]
```
