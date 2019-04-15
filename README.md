# Go Tool Catalog

Gtc installs and updates well known tools written by golang.

## Install gtc

```console
$ go get -u github.com/koron/gtc
```

## Usages


```console
# List tools installed
$ gtc list

# List tools not-installed
$ gtc list -filter notinstalled

# List tools unknown (for gtc)
$ gtc list -filter unknown

# Install a tool
$ gtc install jvgrep

# Install multiple tools
$ gtc install goimports golint jvgrep

# Update a tool
$ gtc update jvgrep

# Update multiple tools
$ gtc update goimports golint jvgrep

# Update all tools which has been installed and not touched over 5 days
$ gtc update -all
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
