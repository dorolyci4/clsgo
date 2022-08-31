# clsgo

<p align="center">
    <a href="https://pkg.go.dev/github.com/lovelacelee/clsgo" title="Go API Reference" rel="nofollow"><img src="https://img.shields.io/badge/go-documentation-blue.svg?style=flat" alt="Go API Reference"></a>
    <a href="https://goreportcard.com/report/github.com/lovelacelee/clsgo"><img src="https://goreportcard.com/badge/github.com/lovelacelee/clsgo" alt="Code Status" /></a>
    <a href="https://github.com/lovelacelee/clsgo/actions/workflows/static_analysis.yml"><img src="https://github.com/lovelacelee/clsgo/actions/workflows/static_analysis.yml/badge.svg" alt="Static Analysis"/></a>
</p>

CLS packages

## Requirments

* [Viper](https://github.com/spf13/viper) [Doc](https://pkg.go.dev/github.com/spf13/viper)
* [GFrame](https://github.com/gogf/gf) [Doc](https://pkg.go.dev/github.com/gogf/gf/v2)

## Todo list

* MQTT client/server support
* Protobuf support
* ✅HTTP static file server
* ✅HTTP RESTFUL API server
* ✅TCP protocol plugin
* ✅TCP client/server
* ✅File logger
* ✅Terminal logger
* ✅Json config loader
* ✅Viper config support

## Quick start

### Running examples

```shell
python example
```

### Running tests

```shell

go install github.com/fzipp/gocyclo/cmd/gocyclo@latest
go install honnef.co/go/tools/cmd/staticcheck@latest

go test ./... -v -bench=\"Bench*\" -count=1
staticcheck ./...
gocyclo -over 50 .
```

[![Star History Chart](https://api.star-history.com/svg?repos=lovelacelee/clsgo&type=Date)](https://star-history.com/#lovelacelee/clsgo&Date)

