# clsgo

<p align="center">
    <a href="https://pkg.go.dev/github.com/lovelacelee/clsgo" title="Go API Reference" rel="nofollow"><img src="https://img.shields.io/badge/go-documentation-blue.svg?style=flat" alt="Go API Reference"></a>
    <a href="https://goreportcard.com/report/github.com/lovelacelee/clsgo"><img src="https://goreportcard.com/badge/github.com/lovelacelee/clsgo" alt="Code Status" /></a>
    <a href="https://github.com/lovelacelee/clsgo/actions/workflows/static_analysis.yml"><img src="https://github.com/lovelacelee/clsgo/actions/workflows/static_analysis.yml/badge.svg" alt="Static Analysis"/></a>
    <a href="https://github.com/lovelacelee/clsgo/actions/workflows/build.yml"><img src="https://github.com/lovelacelee/clsgo/actions/workflows/build.yml/badge.svg" alt="Build Result"/></a>
</p>

[CLSGO](https://pkg.go.dev/github.com/lovelacelee/clsgo) packages is a collection of excellent open source components for Golang learning and development, a step-by-step sample library for the learning phase, and a project template for rapid development for engineering development. Most implement ideas comes from [Viper](https://github.com/spf13/viper), [GFrame](https://github.com/gogf/gf) etc.

## Multi-Platform

[![Ubuntu](https://img.shields.io/badge/Ubuntu-E95420?logo=ubuntu\&logoColor=white)](https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_idruns-on) [![Mac OS](https://img.shields.io/badge/mac%20os-000000?logo=macos\&logoColor=F0F0F0)](https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_idruns-on) [![Windows](https://img.shields.io/badge/Windows-0078D6?logo=windows\&logoColor=white)](https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_idruns-on)

## Fetures

* 🟩Protobuf support
* ✅MQ support, MQTT client/server, rabbitmq publish and consume routine
* ✅HTTP static file server
* ✅HTTP RESTFUL API server
* ✅TCP protocol plugin
* ✅TCP client/server
* ✅File rotate logger
* ✅Terminal logger
* ✅Storage, Cache support, such as gorm, gredis
* ✅Utils library package
* ✅Viper config support, config manager

## Quick start

```shell

# Install latest clsmt from https://github.com/lovelacelee/mirror-tools/releases/latest
go install github.com/fzipp/gocyclo/cmd/gocyclo@latest
go install honnef.co/go/tools/cmd/staticcheck@latest
# Test
clsmt run check test
# Import 
go get github.com/lovelacelee/clsgo
```


<!-- [![Star History Chart](https://api.star-history.com/svg?repos=lovelacelee/clsgo&type=Date)](https://star-history.com/#lovelacelee/clsgo&Date) -->


## Credits

 * [Connard](https://github.com/lovelacelee)

## License

The MIT License (MIT) - see [`LICENSE`](https://github.com/lovelacelee/clsgo/blob/master/LICENSE) for more details.