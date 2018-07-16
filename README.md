[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

# relay42-go

Unofficial GoLang SDK for the [Relay42](https://relay42.com) API

# Install

`go get github.com/ddevcap/relay42-go`

# Getting started

Import into your Go project or library.

```
import (
    "github.com/ddevcap/relay42-go"
)
```

Create an API client in order to interact with the Relay42 API endpoints.

```
username := "your-username"
password := "your-password"
client := relay42.NewClient(username, password)
```

#### SiteID
You need to set the Rely42 site id. 

```
siteID := 0
client.Site(siteID)
```

#### Debug mode
In debug mode, all outgoing http requests are printed nicely in the form of curl command so that you can easly drop into your command line to debug specific request.

```
client.Debug = true
```

# Documentation/References

### GoLang
[Effective Go](https://golang.org/doc/effective_go.html)

## Contributing
[WIP]

## License
MIT