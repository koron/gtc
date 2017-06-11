# Go Tool Catalog

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
$ go list -not

# Install a tool
$ go install jvgrep

# Install multiple tools
$ go install goimports golint jvgrep

# Update a tool
$ go update jvgrep

# Update multiple tools
$ go update goimports golint jvgrep
```
