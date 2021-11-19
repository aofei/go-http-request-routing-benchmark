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
github.com/aofei/r2 v0.2.2
github.com/go-chi/chi/v5 v5.0.7
github.com/gorilla/mux v1.8.0
github.com/julienschmidt/httprouter v1.3.0
$ go test -bench .
goos: linux
goarch: amd64
pkg: github.com/aofei/go-http-request-routing-benchmark
cpu: AMD EPYC 7B12
BenchmarkStatic_R2-8                        2522            445565 ns/op          771571 B/op       1727 allocs/op
BenchmarkStatic_HttpRouter-8                2258            478018 ns/op          771567 B/op       1727 allocs/op
BenchmarkStatic_Chi-8                       2144            536540 ns/op          842228 B/op       2199 allocs/op
BenchmarkStatic_GorillaMux-8                 960           1250253 ns/op          930394 B/op       3150 allocs/op
BenchmarkGitHubAPI_R2-8                     1348            811915 ns/op         1185405 B/op       3353 allocs/op
BenchmarkGitHubAPI_HttpRouter-8             1480            772285 ns/op         1191232 B/op       3352 allocs/op
BenchmarkGitHubAPI_Chi-8                    1441            903787 ns/op         1198201 B/op       3285 allocs/op
BenchmarkGitHubAPI_GorillaMux-8              229           5167329 ns/op         1381182 B/op       4794 allocs/op
BenchmarkGPlusAPI_R2-8                     28180             45249 ns/op           69815 B/op        198 allocs/op
BenchmarkGPlusAPI_HttpRouter-8             26870             42398 ns/op           69907 B/op        198 allocs/op
BenchmarkGPlusAPI_Chi-8                    23659             48228 ns/op           70457 B/op        193 allocs/op
BenchmarkGPlusAPI_GorillaMux-8             14595             87087 ns/op           81168 B/op        282 allocs/op
BenchmarkParseAPI_R2-8                     13683             81912 ns/op          136155 B/op        366 allocs/op
BenchmarkParseAPI_HttpRouter-8             14487             80600 ns/op          135988 B/op        366 allocs/op
BenchmarkParseAPI_Chi-8                    12061             94948 ns/op          140074 B/op        380 allocs/op
BenchmarkParseAPI_GorillaMux-8              8124            172667 ns/op          159558 B/op        552 allocs/op
PASS
ok      github.com/aofei/go-http-request-routing-benchmark      25.071s
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
