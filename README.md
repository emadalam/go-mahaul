# go-mahaul

Parse and validate your environment variables easily in Go with the following benefits.

* Compile time access guarantees
* Parsed values with accurate type-safe Go data types
* Validation of required values before app boot
* Environment specific defaults and fallbacks
* Code autocompletion for environment variables

[Read more](#why-this-package) for understanding why to use this package and its benefits.

## Installation

1. Create a configuration file `internal/env/env.config.json` 
```json
{
  "DEPLOYMENT_ENV": {
    "type": "string",
    "choices": ["staging", "live"]
  },
  "PORT": {
    "type": "port",
    "default": "8080"
  },
  "HOST": {
    "type": "host",
    "default": "//example.com"
  },
  "API_URL": {
    "type": "uri",
    "default": "https://www.example.com/api"
  }
}
```

2. Create the package file `internal/env/env.go` with the following contents
```go
package env

//go:generate go run github.com/emadalam/go-mahaul/cmd/gen@latest
```

Alternatively if you prefer to strictly pin the version of this tool in your `go.mod` alongside your other dependencies, create a `tools.go` file with a build tag:

```go
//go:build tools
// +build tools

package tools

import (
	_ "github.com/emadalam/go-mahaul/cmd/gen"
)
```
Run `go mod tidy` to track the version, and then use this go:generate directive:

```go
package env

//go:generate go run github.com/emadalam/go-mahaul/cmd/gen
```

3. Finally run the following as part of your build process which will generate the `env` package
```sh
go generate ./...
```

## Usage

Once the `env` package is generated, you can import and make use of it anywhere in your app. You must invoke `env.MustValidate()` (or `env.WarnValidate` in dev/test mode) before booting up your app so the validation kicks in and your app boots only when the required environment variables are configured.

```go
package main

import (
	"flag"
	"fmt"

	"github.com/emadalam/some-app/internal/env"
)

func main() {
	isDev := flag.Bool("dev", false, "Runs the app in dev mode")
	flag.Parse()

	if *isDev {
		env.WarnValidate()
	} else {
		env.MustValidate()
	}

	fmt.Println(env.Value.ApiUrl())
	fmt.Println(env.Value.DeploymentEnv())
	fmt.Println(env.Value.Host())
	fmt.Println(env.Value.Port())
}
```

## Supported types

The following type configurations are supported.

| Type      | Go Type    | Valid Environment variables value                              |
| :-------- | :--------- | :------------------------------------------------------------- |
| `:string` | `string`   | Any string                                                     |
| `:num`    | `float64`  | Any string that can be parsed as float                         |
| `:int`    | `int`      | Any string that can be parsed as integer                       |
| `:bool`   | `bool`     | 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False        |
| `:port`   | `uint16`   | Any valid port value between "1" to "65535" (casted as uint16) |
| `:host`   | `*url.URL` | Any valid host name (no scheme or hostname allowed)              |
| `:uri`    | `*url.URL` | Any valid uris (must include scheme and hostname)              |

## Setting Defaults

Any defaults and fallback values can be set globally using the `default` configuration option. Make sure to **use the string values** same as we set in the actual system environment, as it will be parsed depending upon the provided `type` configuration.

## Setting choices list

You can further restrict the parsed values to a predefined list by setting the `choices` option with list of allowed values. Note that the values are parsed first and then matched against the provided list.

## Generator CLI flags

| Flag     | Default          | Description                          |
| :------- | :--------------- | :----------------------------------- |
| -config  | env.config.json  | Path to the JSON config file.        |
| -out     | env_generated.go | Name/path of the generated output.   |
| -pkgName | env              | Package name for the generated file. |


## Why this package

`go-mahaul` accomplishes the following functionalities for streamlining the environment variables requirements for an elixir app.

#### Compile time access guarantees

Using the code generation setup, `go-mahaul` creates a package with compile time methods for accessing the environment variables. This guarantees that there are no accidental typos during the access of the environment variables from the code. Also as an added bonus, if you are using the [Go Language Server](https://go.dev/gopls/) for your development environment, you'd get code autocompletion.

#### Parsed values with accurate data types

Depending upon the configuration, the access to the predefined environment variables string values are parsed and the correct Go data types are returned. `go-mahaul` supports [a wide range](#supported-types) of commonly set environment variable types. It also supports the `choices` options to limit the allowed values for an environment variable.

#### Validation of required values before app boot

Often times we release new versions of the app accessing new environment variables, but we forget to set those for one of our app deployments. This creates nasty bugs that are only discovered when certain parts of the app behaves erratically or fails. With `go-mahaul`, you can pre-validate the existence of the required environment variables with correct values before booting the app (ideally in `main.go`). This ensures that your application will fail to boot unless you have set those environment variables with correct values. This works really well with any cloud deployments that makes new version of your app active and available only after ensuring the new deployment had a successful boot.

#### Defaults and fallbacks

You can [set default values](#setting-defaults) for the production or development environment of your app while configuring `go-mahaul`. This comes handy when you want some defaults for dev/test environment to let other contributors of your app quickly start the dev environment of your app without worrying to set some needed environment variables. Or have some sensible defaults for production version of your app with flexibility to change the values by setting an environment variable.

## Contributing

Contributions are welcome. Please follow the commit guidelines from https://www.conventionalcommits.org.

### Setup

Clone the repo and fetch its dependencies:

```sh
git clone https://github.com/emadalam/go-mahaul.git
cd go-mahaul
go mod download
```

## LICENSE

See [LICENSE](https://github.com/emadalam/go-mahaul/blob/main/LICENSE.txt)
