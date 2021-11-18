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
github.com/aofei/r2 v0.2.1
github.com/go-chi/chi/v5 v5.0.6
github.com/gorilla/mux v1.8.0
github.com/julienschmidt/httprouter v1.3.0
$ go test -bench .
goos: linux
goarch: amd64
pkg: github.com/aofei/go-http-request-routing-benchmark
cpu: AMD EPYC 7B12
BenchmarkStatic_R2-8                        2263            467981 ns/op          771567 B/op       1727 allocs/op
BenchmarkStatic_HttpRouter-8                2138            502016 ns/op          771567 B/op       1727 allocs/op
BenchmarkStatic_Chi-8                       2270            595992 ns/op          842220 B/op       2199 allocs/op
BenchmarkStatic_GorillaMux-8                 926           1299767 ns/op          930387 B/op       3150 allocs/op
BenchmarkGitHubAPI_R2-8                     1404            843452 ns/op         1185418 B/op       3353 allocs/op
BenchmarkGitHubAPI_HttpRouter-8             1477            779844 ns/op         1191230 B/op       3351 allocs/op
BenchmarkGitHubAPI_Chi-8                    1430            957265 ns/op         1198195 B/op       3285 allocs/op
BenchmarkGitHubAPI_GorillaMux-8              225           5302534 ns/op         1381175 B/op       4794 allocs/op
BenchmarkGPlusAPI_R2-8                     24033             47864 ns/op           69815 B/op        198 allocs/op
BenchmarkGPlusAPI_HttpRouter-8             24322             46063 ns/op           69907 B/op        198 allocs/op
BenchmarkGPlusAPI_Chi-8                    24082             49573 ns/op           70457 B/op        193 allocs/op
BenchmarkGPlusAPI_GorillaMux-8             14078             86476 ns/op           81168 B/op        282 allocs/op
BenchmarkParseAPI_R2-8                     13323             86371 ns/op          136155 B/op        366 allocs/op
BenchmarkParseAPI_HttpRouter-8             13741             85138 ns/op          135988 B/op        366 allocs/op
BenchmarkParseAPI_Chi-8                    12650             99532 ns/op          140073 B/op        380 allocs/op
BenchmarkParseAPI_GorillaMux-8              6945            165118 ns/op          159560 B/op        552 allocs/op
PASS
ok      github.com/aofei/go-http-request-routing-benchmark      25.313s
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
