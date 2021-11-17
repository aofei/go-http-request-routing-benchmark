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
github.com/aofei/r2 v0.2.0
github.com/go-chi/chi/v5 v5.0.6
github.com/gorilla/mux v1.8.0
github.com/julienschmidt/httprouter v1.3.0
$ go test -bench .
goos: linux
goarch: amd64
pkg: github.com/aofei/go-http-request-routing-benchmark
cpu: AMD EPYC 7B12
BenchmarkStatic_R2-8                        2470            483046 ns/op          771571 B/op       1727 allocs/op
BenchmarkStatic_HttpRouter-8                2392            496707 ns/op          771562 B/op       1727 allocs/op
BenchmarkStatic_Chi-8                       1798            601870 ns/op          842219 B/op       2199 allocs/op
BenchmarkStatic_GorillaMux-8                 916           1285054 ns/op          930373 B/op       3150 allocs/op
BenchmarkGitHubAPI_R2-8                     1434            896290 ns/op         1185408 B/op       3353 allocs/op
BenchmarkGitHubAPI_HttpRouter-8             1460            838313 ns/op         1191237 B/op       3351 allocs/op
BenchmarkGitHubAPI_Chi-8                    1233            976479 ns/op         1198183 B/op       3285 allocs/op
BenchmarkGitHubAPI_GorillaMux-8              205           5365106 ns/op         1381143 B/op       4794 allocs/op
BenchmarkGPlusAPI_R2-8                     24177             49746 ns/op           69814 B/op        198 allocs/op
BenchmarkGPlusAPI_HttpRouter-8             25266             48744 ns/op           69906 B/op        198 allocs/op
BenchmarkGPlusAPI_Chi-8                    21542             53183 ns/op           70456 B/op        193 allocs/op
BenchmarkGPlusAPI_GorillaMux-8             13753             88778 ns/op           81167 B/op        282 allocs/op
BenchmarkParseAPI_R2-8                     13122             92461 ns/op          136153 B/op        366 allocs/op
BenchmarkParseAPI_HttpRouter-8             14160             87058 ns/op          135988 B/op        366 allocs/op
BenchmarkParseAPI_Chi-8                    10000            103210 ns/op          140071 B/op        380 allocs/op
BenchmarkParseAPI_GorillaMux-8              7861            172459 ns/op          159557 B/op        552 allocs/op
PASS
ok      github.com/aofei/go-http-request-routing-benchmark      24.568s
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
