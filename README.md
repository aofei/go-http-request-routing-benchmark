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
go version go1.17.3 linux/amd64
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
BenchmarkStatic_R2-8                      105404             11325 ns/op               0 B/op          0 allocs/op
BenchmarkStatic_HttpRouter-8               95414             11473 ns/op               0 B/op          0 allocs/op
BenchmarkStatic_Chi-8                      12068             99207 ns/op           70354 B/op        471 allocs/op
BenchmarkStatic_GorillaMux-8                1610            739387 ns/op          158403 B/op       1421 allocs/op
BenchmarkGitHubAPI_R2-8                    26295             45770 ns/op            4417 B/op        184 allocs/op
BenchmarkGitHubAPI_HttpRouter-8            10000            121240 ns/op          101856 B/op        920 allocs/op
BenchmarkGitHubAPI_Chi-8                    6400            165740 ns/op           99036 B/op        663 allocs/op
BenchmarkGitHubAPI_GorillaMux-8              277           4309490 ns/op          281877 B/op       2173 allocs/op
BenchmarkGPlusAPI_R2-8                    574969              2227 ns/op             264 B/op         11 allocs/op
BenchmarkGPlusAPI_HttpRouter-8            176149              6657 ns/op            5832 B/op         55 allocs/op
BenchmarkGPlusAPI_Chi-8                   142299              8447 ns/op            5825 B/op         39 allocs/op
BenchmarkGPlusAPI_GorillaMux-8             31264             39432 ns/op           16533 B/op        128 allocs/op
BenchmarkParseAPI_R2-8                    356826              3243 ns/op             384 B/op         16 allocs/op
BenchmarkParseAPI_HttpRouter-8            119695              9337 ns/op            8192 B/op         80 allocs/op
BenchmarkParseAPI_Chi-8                    78052             16110 ns/op           11651 B/op         78 allocs/op
BenchmarkParseAPI_GorillaMux-8             16302             73312 ns/op           31129 B/op        250 allocs/op
PASS
ok      github.com/aofei/go-http-request-routing-benchmark      22.935s
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
