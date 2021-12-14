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
go version go1.17.5 linux/amd64
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
BenchmarkStatic_R2-8                      104644             11522 ns/op               0 B/op          0 allocs/op
BenchmarkStatic_HttpRouter-8               91291             11384 ns/op               0 B/op          0 allocs/op
BenchmarkStatic_Chi-8                      13148             91956 ns/op           70355 B/op        471 allocs/op
BenchmarkStatic_GorillaMux-8                1638            729118 ns/op          158405 B/op       1421 allocs/op
BenchmarkGitHubAPI_R2-8                    26758             46273 ns/op            4417 B/op        184 allocs/op
BenchmarkGitHubAPI_HttpRouter-8            10000            110456 ns/op          101856 B/op        920 allocs/op
BenchmarkGitHubAPI_Chi-8                    8133            153753 ns/op           99036 B/op        663 allocs/op
BenchmarkGitHubAPI_GorillaMux-8              278           4322280 ns/op          281873 B/op       2173 allocs/op
BenchmarkGPlusAPI_R2-8                    579604              2187 ns/op             264 B/op         11 allocs/op
BenchmarkGPlusAPI_HttpRouter-8            197043              5971 ns/op            5832 B/op         55 allocs/op
BenchmarkGPlusAPI_Chi-8                   160880              7748 ns/op            5825 B/op         39 allocs/op
BenchmarkGPlusAPI_GorillaMux-8             32467             38134 ns/op           16533 B/op        128 allocs/op
BenchmarkParseAPI_R2-8                    383479              3195 ns/op             384 B/op         16 allocs/op
BenchmarkParseAPI_HttpRouter-8            138969              8385 ns/op            8192 B/op         80 allocs/op
BenchmarkParseAPI_Chi-8                    78828             14986 ns/op           11651 B/op         78 allocs/op
BenchmarkParseAPI_GorillaMux-8             16750             71020 ns/op           31130 B/op        250 allocs/op
PASS
ok      github.com/aofei/go-http-request-routing-benchmark      22.914s
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
