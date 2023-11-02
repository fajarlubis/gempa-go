# gempa-go

[![Run Tests](https://github.com/fajarlubis/gempa-go/actions/workflows/ci.yml/badge.svg)](https://github.com/fajarlubis/gempa-go/actions/workflows/ci.yml)
[![GoDoc](https://godoc.org/github.com/fajarlubis/gempa-go?status.svg)](https://godoc.org/github.com/fajarlubis/gempa-go)

Go wrapper for BMKG (Badan Meteorologi, Klimatologi dan Geofisika). All the data provided here https://data.bmkg.go.id/gempabumi

### Why I build this package?

Indonesia is one of many countries that lies on top of the ring of fire with almost more than 10,000 volcanic activity in 2021. This package is mainly built for IoT purposes for interacting with Indonesia government agency (BMKG) web API that informs every tectonic activity. Feel free to use and contribute to this package. I will maintain this package as long as I can do and adding more functionality like custom shakemap, geolocation, etc.

I use this package for integrate with IoT devices like ESP32, Arduino, etc. And for the end-user product will be an app for smartwatch.

## Usage

### Start using it

Download and install it:

```sh
go get -u github.com/fajarlubis/gempa-go
```

Import it in your code:

```go
import "github.com/fajarlubis/gempa-go"
```
