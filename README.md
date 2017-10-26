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
