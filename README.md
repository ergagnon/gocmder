# Go CMDER
Go Cmder combine [Cobra](https://github.com/spf13/cobra) and [Viper](https://github.com/spf13/viper) to auto generate command-line tool where the user can control configuration in a precedence order: **flags > environment variables > configuration files and the defaults defined by the config**.

This library was inspired by [Sting of the viper](https://carolynvanslyck.com/blog/2020/08/sting-of-the-viper/) article from [Carolyn Van Slyck](https://carolynvanslyck.com/).


[![Build](https://github.com/ergagnon/gocmder/actions/workflows/gocmder.yml/badge.svg)](https://github.com/ergagnon/gocmder/actions/workflows/gocmder.yml) [![codecov](https://codecov.io/github/ergagnon/gocmder/branch/main/graph/badge.svg?token=R9SPKS0Y7B)](https://codecov.io/github/ergagnon/gocmder) [![Go Report Card](https://goreportcard.com/badge/github.com/ergagnon/gocmder)](https://goreportcard.com/report/github.com/ergagnon/gocmder) [![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://github.com/ergagnon/gocmder/blob/main/LICENCE)


## Install
```
go get github.com/ergagnon/gocmder
```

## How to use

### Create the config struct with tags

Supported tags:
1. `desc`: used for the Flag description.
2. `default`: default value used in Flag and Viper config. Supported value type: `string`, `int`, `float32`, `bool`
3. `required`: set a required Flag. 
4. `hidden`: don't create a Flag when hidden is true.

Example:
``` go
type AppConfig struct {
    Directory string `desc:"Directory to browse" default:"."`
    Server ServerConfig // support multi-level config          
}

type ServerConfig struct {
    Url string `desc:"Server url" default:"localhost"`
    Port int `desc:"Username" default:"8080"`
}
```

### Create a Go CMDER with the config and options
Call the function
``` go
func NewCmder(cfg any, onFinalize OnFinalizeFunc, opts ...CmderOption) (*Cmder, error)
```

```go
cli, err := gocmder.NewCmder(AppConfig{}, func(cfg any) error {
    // At this step, the config object has been filled
    // with all the values. 
    // You can use it in your application.
    return nil
})
```

This will auto-generate 

**Flags**:
```
Usage:
   [flags]

Flags:
      --directory string    Directory to browse (default ".")
  -h, --help                help for this command
      --server-port int     Username (default 8080)
      --server-url string   Server url (default "localhost")
```

**Environment Variable**
```
DIRECTORY
SERVER_PORT
SERVER_URL
```

**Config file (optional)**  
Configuration will use [Viper](https://github.com/spf13/viper) defaults.
> reading from JSON, TOML, YAML, HCL, envfile and Java properties config files

Yaml example:

```yaml
directory: "my_dir"
server:
  url: "127.0.0.1"
  port: 8080
```