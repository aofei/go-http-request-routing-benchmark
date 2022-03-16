# Go HTTP Request Routing Benchmark

An HTTP request routing benchmark suite for Go.

## Routers

* [R2](https://github.com/aofei/r2)
* [HttpRouter](https://github.com/julienschmidt/httprouter)
* [Chi](https://github.com/go-chi/chi)
* [gorilla/mux](https://github.com/gorilla/mux)

A new router is only eligible to be added if it:

* Is designed to replace [http.ServeMux](https://pkg.go.dev/net/http#ServeMux).
* Supports path parameters.
* Can be fully functional without any [`http.Handler`](https://pkg.go.dev/net/http#Handler) variant.

## Results

```bash
$ go version
go version go1.18 linux/amd64
$ go list -m all | tail -n +2
github.com/aofei/r2 v0.3.0
github.com/go-chi/chi/v5 v5.0.7
github.com/gorilla/mux v1.8.0
github.com/julienschmidt/httprouter v1.3.0
$ go test -bench .
goos: linux
goarch: amd64
pkg: github.com/aofei/go-http-request-routing-benchmark
cpu: AMD EPYC 7B12
BenchmarkStatic_R2-8                      113211             10635 ns/op               0 B/op          0 allocs/op
BenchmarkStatic_HttpRouter-8              107781             11074 ns/op               0 B/op          0 allocs/op
BenchmarkStatic_Chi-8                      13286             89292 ns/op           70355 B/op        471 allocs/op
BenchmarkStatic_GorillaMux-8                1695            703530 ns/op          158402 B/op       1421 allocs/op
BenchmarkGitHubAPI_R2-8                    25297             47480 ns/op            4417 B/op        184 allocs/op
BenchmarkGitHubAPI_HttpRouter-8            10000            106909 ns/op          101856 B/op        920 allocs/op
BenchmarkGitHubAPI_Chi-8                    8036            150422 ns/op           99036 B/op        663 allocs/op
BenchmarkGitHubAPI_GorillaMux-8              283           4214569 ns/op          281874 B/op       2173 allocs/op
BenchmarkGPlusAPI_R2-8                    505563              2356 ns/op             264 B/op         11 allocs/op
BenchmarkGPlusAPI_HttpRouter-8            209221              5781 ns/op            5832 B/op         55 allocs/op
BenchmarkGPlusAPI_Chi-8                   158458              7484 ns/op            5825 B/op         39 allocs/op
BenchmarkGPlusAPI_GorillaMux-8             33201             36852 ns/op           16533 B/op        128 allocs/op
BenchmarkParseAPI_R2-8                    359286              3317 ns/op             384 B/op         16 allocs/op
BenchmarkParseAPI_HttpRouter-8            135511              8346 ns/op            8192 B/op         80 allocs/op
BenchmarkParseAPI_Chi-8                    86196             14122 ns/op           11651 B/op         78 allocs/op
BenchmarkParseAPI_GorillaMux-8             16698             70745 ns/op           31129 B/op        250 allocs/op
PASS
ok      github.com/aofei/go-http-request-routing-benchmark      22.717s
```

## Community

If you want to discuss Go HTTP Request Routing Benchmark, or ask questions about
it, simply post questions or ideas
[here](https://github.com/aofei/go-http-request-routing-benchmark/issues).

## Contributing

If you want to help build Go HTTP Request Routing Benchmark, simply follow
[this](https://github.com/aofei/go-http-request-routing-benchmark/wiki/Contributing)
to send pull requests
[here](https://github.com/aofei/go-http-request-routing-benchmark/pulls).

## License

This project is licensed under the MIT License.

License can be found [here](LICENSE).
