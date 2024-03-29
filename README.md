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
go version go1.20.3 darwin/arm64
$ go list -m all | tail -n +2
github.com/aofei/r2 v0.3.3
github.com/go-chi/chi/v5 v5.0.8
github.com/gorilla/mux v1.8.0
github.com/julienschmidt/httprouter v1.3.0
$ go test -bench .
goos: darwin
goarch: arm64
pkg: github.com/aofei/go-http-request-routing-benchmark
BenchmarkStatic_R2-8                      236452              5084 ns/op               0 B/op          0 allocs/op
BenchmarkStatic_HttpRouter-8              185234              5618 ns/op               0 B/op          0 allocs/op
BenchmarkStatic_Chi-8                      35596             33717 ns/op           47741 B/op        314 allocs/op
BenchmarkStatic_GorillaMux-8                3904            306781 ns/op          113171 B/op       1107 allocs/op
BenchmarkGitHubAPI_R2-8                    54349             22111 ns/op            4417 B/op        184 allocs/op
BenchmarkGitHubAPI_HttpRouter-8            28010             43961 ns/op           75360 B/op        736 allocs/op
BenchmarkGitHubAPI_Chi-8                   20300             58942 ns/op           67203 B/op        442 allocs/op
BenchmarkGitHubAPI_GorillaMux-8              562           2118916 ns/op          218188 B/op       1731 allocs/op
BenchmarkGPlusAPI_R2-8                   1000000              1023 ns/op             264 B/op         11 allocs/op
BenchmarkGPlusAPI_HttpRouter-8            546333              2195 ns/op            4248 B/op         44 allocs/op
BenchmarkGPlusAPI_Chi-8                   432822              2780 ns/op            3953 B/op         26 allocs/op
BenchmarkGPlusAPI_GorillaMux-8             74625             16112 ns/op           12788 B/op        102 allocs/op
BenchmarkParseAPI_R2-8                    764920              1496 ns/op             384 B/op         16 allocs/op
BenchmarkParseAPI_HttpRouter-8            373254              3322 ns/op            5888 B/op         64 allocs/op
BenchmarkParseAPI_Chi-8                   226174              5345 ns/op            7906 B/op         52 allocs/op
BenchmarkParseAPI_GorillaMux-8             41613             28922 ns/op           23639 B/op        198 allocs/op
PASS
ok      github.com/aofei/go-http-request-routing-benchmark      22.602s
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
