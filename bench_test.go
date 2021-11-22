package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aofei/r2"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/mux"
	"github.com/julienschmidt/httprouter"
)

type route struct {
	method string
	path   string
}

var (
	staticRoutes = []*route{
		{"GET", "/"},
		{"GET", "/Makefile"},
		{"GET", "/articles/"},
		{"GET", "/articles/go_command.html"},
		{"GET", "/articles/index.html"},
		{"GET", "/articles/wiki/"},
		{"GET", "/articles/wiki/Makefile"},
		{"GET", "/articles/wiki/edit.html"},
		{"GET", "/articles/wiki/final-noclosure.go"},
		{"GET", "/articles/wiki/final-noerror.go"},
		{"GET", "/articles/wiki/final-parsetemplate.go"},
		{"GET", "/articles/wiki/final-template.go"},
		{"GET", "/articles/wiki/final.go"},
		{"GET", "/articles/wiki/get.go"},
		{"GET", "/articles/wiki/http-sample.go"},
		{"GET", "/articles/wiki/index.html"},
		{"GET", "/articles/wiki/notemplate.go"},
		{"GET", "/articles/wiki/part1-noerror.go"},
		{"GET", "/articles/wiki/part1.go"},
		{"GET", "/articles/wiki/part2.go"},
		{"GET", "/articles/wiki/part3-errorhandling.go"},
		{"GET", "/articles/wiki/part3.go"},
		{"GET", "/articles/wiki/test.bash"},
		{"GET", "/articles/wiki/test_Test.txt.good"},
		{"GET", "/articles/wiki/test_edit.good"},
		{"GET", "/articles/wiki/test_view.good"},
		{"GET", "/articles/wiki/view.html"},
		{"GET", "/cmd.html"},
		{"GET", "/code.html"},
		{"GET", "/codewalk/"},
		{"GET", "/codewalk/codewalk.css"},
		{"GET", "/codewalk/codewalk.js"},
		{"GET", "/codewalk/codewalk.xml"},
		{"GET", "/codewalk/functions.xml"},
		{"GET", "/codewalk/markov.go"},
		{"GET", "/codewalk/markov.xml"},
		{"GET", "/codewalk/pig.go"},
		{"GET", "/codewalk/popout.png"},
		{"GET", "/codewalk/run"},
		{"GET", "/codewalk/sharemem.xml"},
		{"GET", "/codewalk/urlpoll.go"},
		{"GET", "/contrib.html"},
		{"GET", "/contribute.html"},
		{"GET", "/debugging_with_gdb.html"},
		{"GET", "/devel/"},
		{"GET", "/devel/release.html"},
		{"GET", "/devel/weekly.html"},
		{"GET", "/docs.html"},
		{"GET", "/effective_go.html"},
		{"GET", "/files.log"},
		{"GET", "/gccgo_contribute.html"},
		{"GET", "/gccgo_install.html"},
		{"GET", "/go-logo-black.png"},
		{"GET", "/go-logo-blue.png"},
		{"GET", "/go-logo-white.png"},
		{"GET", "/go1.1.html"},
		{"GET", "/go1.2.html"},
		{"GET", "/go1.html"},
		{"GET", "/go1compat.html"},
		{"GET", "/go_faq.html"},
		{"GET", "/go_mem.html"},
		{"GET", "/go_spec.html"},
		{"GET", "/gopher/"},
		{"GET", "/gopher/appenginegopher.jpg"},
		{"GET", "/gopher/appenginegophercolor.jpg"},
		{"GET", "/gopher/appenginelogo.gif"},
		{"GET", "/gopher/bumper.png"},
		{"GET", "/gopher/bumper192x108.png"},
		{"GET", "/gopher/bumper320x180.png"},
		{"GET", "/gopher/bumper480x270.png"},
		{"GET", "/gopher/bumper640x360.png"},
		{"GET", "/gopher/doc.png"},
		{"GET", "/gopher/frontpage.png"},
		{"GET", "/gopher/gopherbw.png"},
		{"GET", "/gopher/gophercolor.png"},
		{"GET", "/gopher/gophercolor16x16.png"},
		{"GET", "/gopher/help.png"},
		{"GET", "/gopher/pencil/"},
		{"GET", "/gopher/pencil/gopherhat.jpg"},
		{"GET", "/gopher/pencil/gopherhelmet.jpg"},
		{"GET", "/gopher/pencil/gophermega.jpg"},
		{"GET", "/gopher/pencil/gopherrunning.jpg"},
		{"GET", "/gopher/pencil/gopherswim.jpg"},
		{"GET", "/gopher/pencil/gopherswrench.jpg"},
		{"GET", "/gopher/pkg.png"},
		{"GET", "/gopher/project.png"},
		{"GET", "/gopher/ref.png"},
		{"GET", "/gopher/run.png"},
		{"GET", "/gopher/talks.png"},
		{"GET", "/help.html"},
		{"GET", "/ie.css"},
		{"GET", "/install-source.html"},
		{"GET", "/install.html"},
		{"GET", "/logo-153x55.png"},
		{"GET", "/play/"},
		{"GET", "/play/fib.go"},
		{"GET", "/play/hello.go"},
		{"GET", "/play/life.go"},
		{"GET", "/play/peano.go"},
		{"GET", "/play/pi.go"},
		{"GET", "/play/sieve.go"},
		{"GET", "/play/solitaire.go"},
		{"GET", "/play/tree.go"},
		{"GET", "/progs/"},
		{"GET", "/progs/cgo1.go"},
		{"GET", "/progs/cgo2.go"},
		{"GET", "/progs/cgo3.go"},
		{"GET", "/progs/cgo4.go"},
		{"GET", "/progs/defer.go"},
		{"GET", "/progs/defer.out"},
		{"GET", "/progs/defer2.go"},
		{"GET", "/progs/defer2.out"},
		{"GET", "/progs/eff_bytesize.go"},
		{"GET", "/progs/eff_bytesize.out"},
		{"GET", "/progs/eff_qr.go"},
		{"GET", "/progs/eff_sequence.go"},
		{"GET", "/progs/eff_sequence.out"},
		{"GET", "/progs/eff_unused1.go"},
		{"GET", "/progs/eff_unused2.go"},
		{"GET", "/progs/error.go"},
		{"GET", "/progs/error2.go"},
		{"GET", "/progs/error3.go"},
		{"GET", "/progs/error4.go"},
		{"GET", "/progs/go1.go"},
		{"GET", "/progs/gobs1.go"},
		{"GET", "/progs/gobs2.go"},
		{"GET", "/progs/image_draw.go"},
		{"GET", "/progs/image_package1.go"},
		{"GET", "/progs/image_package1.out"},
		{"GET", "/progs/image_package2.go"},
		{"GET", "/progs/image_package2.out"},
		{"GET", "/progs/image_package3.go"},
		{"GET", "/progs/image_package3.out"},
		{"GET", "/progs/image_package4.go"},
		{"GET", "/progs/image_package4.out"},
		{"GET", "/progs/image_package5.go"},
		{"GET", "/progs/image_package5.out"},
		{"GET", "/progs/image_package6.go"},
		{"GET", "/progs/image_package6.out"},
		{"GET", "/progs/interface.go"},
		{"GET", "/progs/interface2.go"},
		{"GET", "/progs/interface2.out"},
		{"GET", "/progs/json1.go"},
		{"GET", "/progs/json2.go"},
		{"GET", "/progs/json2.out"},
		{"GET", "/progs/json3.go"},
		{"GET", "/progs/json4.go"},
		{"GET", "/progs/json5.go"},
		{"GET", "/progs/run"},
		{"GET", "/progs/slices.go"},
		{"GET", "/progs/timeout1.go"},
		{"GET", "/progs/timeout2.go"},
		{"GET", "/progs/update.bash"},
		{"GET", "/root.html"},
		{"GET", "/share.png"},
		{"GET", "/sieve.gif"},
		{"GET", "/tos.html"},
	}

	githubAPIRoutes = []*route{
		// The following routes are disabled due to the lack of
		// functionality of HttpRouter.
		//
		// {"DELETE", "/repos/:owner/:repo/issues/comments/:id"},
		// {"GET", "/gists/public"},
		// {"GET", "/gists/starred"},
		// {"GET", "/repos/:owner/:repo/:archive_format/:ref"},
		// {"GET", "/repos/:owner/:repo/issues/comments"},
		// {"GET", "/repos/:owner/:repo/issues/comments/:id"},
		// {"GET", "/repos/:owner/:repo/issues/events"},
		// {"GET", "/repos/:owner/:repo/issues/events/:id"},
		// {"GET", "/repos/:owner/:repo/pulls/comments"},
		// {"GET", "/repos/:owner/:repo/pulls/comments/:number"},
		// {"PATCH", "/repos/:owner/:repo/issues/comments/:id"},
		// {"PATCH", "/repos/:owner/:repo/pulls/comments/:number"},

		{"DELETE", "/applications/:client_id/tokens"},
		{"DELETE", "/applications/:client_id/tokens/:access_token"},
		{"DELETE", "/authorizations/:id"},
		{"DELETE", "/gists/:id"},
		{"DELETE", "/gists/:id/star"},
		{"DELETE", "/notifications/threads/:id/subscription"},
		{"DELETE", "/orgs/:org/members/:user"},
		{"DELETE", "/orgs/:org/public_members/:user"},
		{"DELETE", "/repos/:owner/:repo"},
		{"DELETE", "/repos/:owner/:repo/collaborators/:user"},
		{"DELETE", "/repos/:owner/:repo/comments/:id"},
		{"DELETE", "/repos/:owner/:repo/downloads/:id"},
		{"DELETE", "/repos/:owner/:repo/hooks/:id"},
		{"DELETE", "/repos/:owner/:repo/issues/:number/labels"},
		{"DELETE", "/repos/:owner/:repo/issues/:number/labels/:name"},
		{"DELETE", "/repos/:owner/:repo/keys/:id"},
		{"DELETE", "/repos/:owner/:repo/labels/:name"},
		{"DELETE", "/repos/:owner/:repo/milestones/:number"},
		{"DELETE", "/repos/:owner/:repo/pulls/comments/:number"},
		{"DELETE", "/repos/:owner/:repo/releases/:id"},
		{"DELETE", "/repos/:owner/:repo/subscription"},
		{"DELETE", "/teams/:id"},
		{"DELETE", "/teams/:id/members/:user"},
		{"DELETE", "/teams/:id/repos/:owner/:repo"},
		{"DELETE", "/user/emails"},
		{"DELETE", "/user/following/:user"},
		{"DELETE", "/user/keys/:id"},
		{"DELETE", "/user/starred/:owner/:repo"},
		{"DELETE", "/user/subscriptions/:owner/:repo"},
		{"GET", "/applications/:client_id/tokens/:access_token"},
		{"GET", "/authorizations"},
		{"GET", "/authorizations/:id"},
		{"GET", "/emojis"},
		{"GET", "/events"},
		{"GET", "/feeds"},
		{"GET", "/gists"},
		{"GET", "/gists/:id"},
		{"GET", "/gists/:id/star"},
		{"GET", "/gitignore/templates"},
		{"GET", "/gitignore/templates/:name"},
		{"GET", "/issues"},
		{"GET", "/legacy/issues/search/:owner/:repo/:state/:keyword"},
		{"GET", "/legacy/repos/search/:keyword"},
		{"GET", "/legacy/user/email/:email"},
		{"GET", "/legacy/user/search/:keyword"},
		{"GET", "/meta"},
		{"GET", "/networks/:owner/:repo/events"},
		{"GET", "/notifications"},
		{"GET", "/notifications/threads/:id"},
		{"GET", "/notifications/threads/:id/subscription"},
		{"GET", "/orgs/:org"},
		{"GET", "/orgs/:org/events"},
		{"GET", "/orgs/:org/issues"},
		{"GET", "/orgs/:org/members"},
		{"GET", "/orgs/:org/members/:user"},
		{"GET", "/orgs/:org/public_members"},
		{"GET", "/orgs/:org/public_members/:user"},
		{"GET", "/orgs/:org/repos"},
		{"GET", "/orgs/:org/teams"},
		{"GET", "/rate_limit"},
		{"GET", "/repos/:owner/:repo"},
		{"GET", "/repos/:owner/:repo/assignees"},
		{"GET", "/repos/:owner/:repo/assignees/:assignee"},
		{"GET", "/repos/:owner/:repo/branches"},
		{"GET", "/repos/:owner/:repo/branches/:branch"},
		{"GET", "/repos/:owner/:repo/collaborators"},
		{"GET", "/repos/:owner/:repo/collaborators/:user"},
		{"GET", "/repos/:owner/:repo/comments"},
		{"GET", "/repos/:owner/:repo/comments/:id"},
		{"GET", "/repos/:owner/:repo/commits"},
		{"GET", "/repos/:owner/:repo/commits/:sha"},
		{"GET", "/repos/:owner/:repo/commits/:sha/comments"},
		{"GET", "/repos/:owner/:repo/contributors"},
		{"GET", "/repos/:owner/:repo/downloads"},
		{"GET", "/repos/:owner/:repo/downloads/:id"},
		{"GET", "/repos/:owner/:repo/events"},
		{"GET", "/repos/:owner/:repo/forks"},
		{"GET", "/repos/:owner/:repo/git/blobs/:sha"},
		{"GET", "/repos/:owner/:repo/git/commits/:sha"},
		{"GET", "/repos/:owner/:repo/git/refs"},
		{"GET", "/repos/:owner/:repo/git/tags/:sha"},
		{"GET", "/repos/:owner/:repo/git/trees/:sha"},
		{"GET", "/repos/:owner/:repo/hooks"},
		{"GET", "/repos/:owner/:repo/hooks/:id"},
		{"GET", "/repos/:owner/:repo/issues"},
		{"GET", "/repos/:owner/:repo/issues/:number"},
		{"GET", "/repos/:owner/:repo/issues/:number/comments"},
		{"GET", "/repos/:owner/:repo/issues/:number/events"},
		{"GET", "/repos/:owner/:repo/issues/:number/labels"},
		{"GET", "/repos/:owner/:repo/keys"},
		{"GET", "/repos/:owner/:repo/keys/:id"},
		{"GET", "/repos/:owner/:repo/labels"},
		{"GET", "/repos/:owner/:repo/labels/:name"},
		{"GET", "/repos/:owner/:repo/languages"},
		{"GET", "/repos/:owner/:repo/milestones"},
		{"GET", "/repos/:owner/:repo/milestones/:number"},
		{"GET", "/repos/:owner/:repo/milestones/:number/labels"},
		{"GET", "/repos/:owner/:repo/notifications"},
		{"GET", "/repos/:owner/:repo/pulls"},
		{"GET", "/repos/:owner/:repo/pulls/:number"},
		{"GET", "/repos/:owner/:repo/pulls/:number/comments"},
		{"GET", "/repos/:owner/:repo/pulls/:number/commits"},
		{"GET", "/repos/:owner/:repo/pulls/:number/files"},
		{"GET", "/repos/:owner/:repo/pulls/:number/merge"},
		{"GET", "/repos/:owner/:repo/readme"},
		{"GET", "/repos/:owner/:repo/releases"},
		{"GET", "/repos/:owner/:repo/releases/:id"},
		{"GET", "/repos/:owner/:repo/releases/:id/assets"},
		{"GET", "/repos/:owner/:repo/stargazers"},
		{"GET", "/repos/:owner/:repo/stats/code_frequency"},
		{"GET", "/repos/:owner/:repo/stats/commit_activity"},
		{"GET", "/repos/:owner/:repo/stats/contributors"},
		{"GET", "/repos/:owner/:repo/stats/participation"},
		{"GET", "/repos/:owner/:repo/stats/punch_card"},
		{"GET", "/repos/:owner/:repo/statuses/:ref"},
		{"GET", "/repos/:owner/:repo/subscribers"},
		{"GET", "/repos/:owner/:repo/subscription"},
		{"GET", "/repos/:owner/:repo/tags"},
		{"GET", "/repos/:owner/:repo/teams"},
		{"GET", "/repositories"},
		{"GET", "/search/code"},
		{"GET", "/search/issues"},
		{"GET", "/search/repositories"},
		{"GET", "/search/users"},
		{"GET", "/teams/:id"},
		{"GET", "/teams/:id/members"},
		{"GET", "/teams/:id/members/:user"},
		{"GET", "/teams/:id/repos"},
		{"GET", "/teams/:id/repos/:owner/:repo"},
		{"GET", "/user"},
		{"GET", "/user/emails"},
		{"GET", "/user/followers"},
		{"GET", "/user/following"},
		{"GET", "/user/following/:user"},
		{"GET", "/user/issues"},
		{"GET", "/user/keys"},
		{"GET", "/user/keys/:id"},
		{"GET", "/user/orgs"},
		{"GET", "/user/repos"},
		{"GET", "/user/starred"},
		{"GET", "/user/starred/:owner/:repo"},
		{"GET", "/user/subscriptions"},
		{"GET", "/user/subscriptions/:owner/:repo"},
		{"GET", "/user/teams"},
		{"GET", "/users"},
		{"GET", "/users/:user"},
		{"GET", "/users/:user/events"},
		{"GET", "/users/:user/events/orgs/:org"},
		{"GET", "/users/:user/events/public"},
		{"GET", "/users/:user/followers"},
		{"GET", "/users/:user/following"},
		{"GET", "/users/:user/following/:target_user"},
		{"GET", "/users/:user/gists"},
		{"GET", "/users/:user/keys"},
		{"GET", "/users/:user/orgs"},
		{"GET", "/users/:user/received_events"},
		{"GET", "/users/:user/received_events/public"},
		{"GET", "/users/:user/repos"},
		{"GET", "/users/:user/starred"},
		{"GET", "/users/:user/subscriptions"},
		{"PATCH", "/authorizations/:id"},
		{"PATCH", "/gists/:id"},
		{"PATCH", "/notifications/threads/:id"},
		{"PATCH", "/orgs/:org"},
		{"PATCH", "/repos/:owner/:repo"},
		{"PATCH", "/repos/:owner/:repo/comments/:id"},
		{"PATCH", "/repos/:owner/:repo/hooks/:id"},
		{"PATCH", "/repos/:owner/:repo/issues/:number"},
		{"PATCH", "/repos/:owner/:repo/keys/:id"},
		{"PATCH", "/repos/:owner/:repo/labels/:name"},
		{"PATCH", "/repos/:owner/:repo/milestones/:number"},
		{"PATCH", "/repos/:owner/:repo/pulls/:number"},
		{"PATCH", "/repos/:owner/:repo/releases/:id"},
		{"PATCH", "/teams/:id"},
		{"PATCH", "/user"},
		{"PATCH", "/user/keys/:id"},
		{"POST", "/authorizations"},
		{"POST", "/gists"},
		{"POST", "/gists/:id/forks"},
		{"POST", "/markdown"},
		{"POST", "/markdown/raw"},
		{"POST", "/orgs/:org/repos"},
		{"POST", "/orgs/:org/teams"},
		{"POST", "/repos/:owner/:repo/commits/:sha/comments"},
		{"POST", "/repos/:owner/:repo/forks"},
		{"POST", "/repos/:owner/:repo/git/blobs"},
		{"POST", "/repos/:owner/:repo/git/commits"},
		{"POST", "/repos/:owner/:repo/git/refs"},
		{"POST", "/repos/:owner/:repo/git/tags"},
		{"POST", "/repos/:owner/:repo/git/trees"},
		{"POST", "/repos/:owner/:repo/hooks"},
		{"POST", "/repos/:owner/:repo/hooks/:id/tests"},
		{"POST", "/repos/:owner/:repo/issues"},
		{"POST", "/repos/:owner/:repo/issues/:number/comments"},
		{"POST", "/repos/:owner/:repo/issues/:number/labels"},
		{"POST", "/repos/:owner/:repo/keys"},
		{"POST", "/repos/:owner/:repo/labels"},
		{"POST", "/repos/:owner/:repo/merges"},
		{"POST", "/repos/:owner/:repo/milestones"},
		{"POST", "/repos/:owner/:repo/pulls"},
		{"POST", "/repos/:owner/:repo/releases"},
		{"POST", "/repos/:owner/:repo/statuses/:ref"},
		{"POST", "/user/emails"},
		{"POST", "/user/keys"},
		{"POST", "/user/repos"},
		{"PUT", "/authorizations/clients/:client_id"},
		{"PUT", "/gists/:id/star"},
		{"PUT", "/notifications"},
		{"PUT", "/notifications/threads/:id/subscription"},
		{"PUT", "/orgs/:org/public_members/:user"},
		{"PUT", "/repos/:owner/:repo/collaborators/:user"},
		{"PUT", "/repos/:owner/:repo/issues/:number/labels"},
		{"PUT", "/repos/:owner/:repo/notifications"},
		{"PUT", "/repos/:owner/:repo/pulls/:number/comments"},
		{"PUT", "/repos/:owner/:repo/pulls/:number/merge"},
		{"PUT", "/repos/:owner/:repo/subscription"},
		{"PUT", "/teams/:id/members/:user"},
		{"PUT", "/teams/:id/repos/:owner/:repo"},
		{"PUT", "/user/following/:user"},
		{"PUT", "/user/starred/:owner/:repo"},
		{"PUT", "/user/subscriptions/:owner/:repo"},
	}

	gplusAPIRoutes = []*route{
		{"DELETE", "/moments/:id"},
		{"GET", "/activities"},
		{"GET", "/activities/:activityId"},
		{"GET", "/activities/:activityId/comments"},
		{"GET", "/activities/:activityId/people/:collection"},
		{"GET", "/comments/:commentId"},
		{"GET", "/people"},
		{"GET", "/people/:userId"},
		{"GET", "/people/:userId/activities/:collection"},
		{"GET", "/people/:userId/moments/:collection"},
		{"GET", "/people/:userId/openIdConnect"},
		{"GET", "/people/:userId/people/:collection"},
		{"POST", "/people/:userId/moments/:collection"},
	}

	parseAPIRoutes = []*route{
		{"DELETE", "/1/classes/:className/:objectId"},
		{"DELETE", "/1/installations/:objectId"},
		{"DELETE", "/1/roles/:objectId"},
		{"DELETE", "/1/users/:objectId"},
		{"GET", "/1/classes/:className"},
		{"GET", "/1/classes/:className/:objectId"},
		{"GET", "/1/installations"},
		{"GET", "/1/installations/:objectId"},
		{"GET", "/1/login"},
		{"GET", "/1/roles"},
		{"GET", "/1/roles/:objectId"},
		{"GET", "/1/users"},
		{"GET", "/1/users/:objectId"},
		{"POST", "/1/classes/:className"},
		{"POST", "/1/events/:eventName"},
		{"POST", "/1/files/:fileName"},
		{"POST", "/1/functions"},
		{"POST", "/1/installations"},
		{"POST", "/1/push"},
		{"POST", "/1/requestPasswordReset"},
		{"POST", "/1/roles"},
		{"POST", "/1/users"},
		{"PUT", "/1/classes/:className/:objectId"},
		{"PUT", "/1/installations/:objectId"},
		{"PUT", "/1/roles/:objectId"},
		{"PUT", "/1/users/:objectId"},
	}

	githubAPICurlyBracketRoutes = curlyBracketRoutes(githubAPIRoutes)
	gplusAPICurlyBracketRoutes  = curlyBracketRoutes(gplusAPIRoutes)
	parseAPICurlyBracketRoutes  = curlyBracketRoutes(parseAPIRoutes)

	staticR2         = r2Handler(staticRoutes)
	staticHttpRouter = httprouterHandler(staticRoutes)
	staticChi        = chiHandler(staticRoutes)
	staticGorillaMux = gorillamuxHandler(staticRoutes)

	githubAPIR2         = r2Handler(githubAPIRoutes)
	githubAPIHttpRouter = httprouterHandler(githubAPIRoutes)
	githubAPIChi        = chiHandler(githubAPICurlyBracketRoutes)
	githubAPIGorillaMux = gorillamuxHandler(githubAPICurlyBracketRoutes)

	gplusAPIR2         = r2Handler(gplusAPIRoutes)
	gplusAPIHttpRouter = httprouterHandler(gplusAPIRoutes)
	gplusAPIChi        = chiHandler(gplusAPICurlyBracketRoutes)
	gplusAPIGorillaMux = gorillamuxHandler(gplusAPICurlyBracketRoutes)

	parseAPIR2         = r2Handler(parseAPIRoutes)
	parseAPIHttpRouter = httprouterHandler(parseAPIRoutes)
	parseAPIChi        = chiHandler(parseAPICurlyBracketRoutes)
	parseAPIGorillaMux = gorillamuxHandler(parseAPICurlyBracketRoutes)
)

func curlyBracketRoutes(routes []*route) []*route {
	cbrs := make([]*route, len(routes))
	for i, route := range routes {
		route := *route
		cbrs[i] = &route
	}

	for _, route := range cbrs {
		for i, l := 0, len(route.path); i < l; i++ {
			if route.path[i] == ':' {
				j := i + 1
				for ; i < l && route.path[i] != '/'; i++ {
				}

				paramName := route.path[j:i]
				route.path = route.path[:j-1] +
					"{" + paramName + "}" +
					route.path[i:]

				i, l = j, len(route.path)
			}
		}
	}

	return cbrs
}

var handler = http.HandlerFunc(func(_ http.ResponseWriter, _ *http.Request) {})

func r2Handler(routes []*route) http.Handler {
	router := &r2.Router{}
	for _, route := range routes {
		router.Handle(route.method, route.path, handler)
	}

	return router
}

func httprouterHandler(routes []*route) http.Handler {
	router := httprouter.New()
	for _, route := range routes {
		router.Handler(route.method, route.path, handler)
	}

	return router
}

func chiHandler(routes []*route) http.Handler {
	router := chi.NewRouter()
	for _, route := range routes {
		router.Method(route.method, route.path, handler)
	}

	return router
}

func gorillamuxHandler(routes []*route) http.Handler {
	router := mux.NewRouter()
	for _, route := range routes {
		router.Handle(route.path, handler).Methods(route.method)
	}

	return router
}

func benchmark(b *testing.B, h http.Handler, req *http.Request, rs []*route) {
	b.ReportAllocs()

	if req == nil {
		req = httptest.NewRequest("", "/", nil)
	}

	rec := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		for _, r := range rs {
			req.Method = r.method
			req.RequestURI = r.path
			req.URL.Path = r.path
			req.URL.RawPath = r.path
			h.ServeHTTP(rec, req)
		}
	}
}

func BenchmarkStatic_R2(b *testing.B) {
	req := httptest.NewRequest("", "/", nil).WithContext(r2.Context())
	benchmark(b, staticR2, req, staticRoutes)
}

func BenchmarkStatic_HttpRouter(b *testing.B) {
	benchmark(b, staticHttpRouter, nil, staticRoutes)
}

func BenchmarkStatic_Chi(b *testing.B) {
	benchmark(b, staticChi, nil, staticRoutes)
}

func BenchmarkStatic_GorillaMux(b *testing.B) {
	benchmark(b, staticGorillaMux, nil, staticRoutes)
}

func BenchmarkGitHubAPI_R2(b *testing.B) {
	req := httptest.NewRequest("", "/", nil).WithContext(r2.Context())
	benchmark(b, githubAPIR2, req, githubAPIRoutes)
}

func BenchmarkGitHubAPI_HttpRouter(b *testing.B) {
	benchmark(b, githubAPIHttpRouter, nil, githubAPIRoutes)
}

func BenchmarkGitHubAPI_Chi(b *testing.B) {
	benchmark(b, githubAPIChi, nil, githubAPICurlyBracketRoutes)
}

func BenchmarkGitHubAPI_GorillaMux(b *testing.B) {
	benchmark(b, githubAPIGorillaMux, nil, githubAPICurlyBracketRoutes)
}

func BenchmarkGPlusAPI_R2(b *testing.B) {
	req := httptest.NewRequest("", "/", nil).WithContext(r2.Context())
	benchmark(b, gplusAPIR2, req, gplusAPIRoutes)
}

func BenchmarkGPlusAPI_HttpRouter(b *testing.B) {
	benchmark(b, gplusAPIHttpRouter, nil, gplusAPIRoutes)
}

func BenchmarkGPlusAPI_Chi(b *testing.B) {
	benchmark(b, gplusAPIChi, nil, gplusAPICurlyBracketRoutes)
}

func BenchmarkGPlusAPI_GorillaMux(b *testing.B) {
	benchmark(b, gplusAPIGorillaMux, nil, gplusAPICurlyBracketRoutes)
}

func BenchmarkParseAPI_R2(b *testing.B) {
	req := httptest.NewRequest("", "/", nil).WithContext(r2.Context())
	benchmark(b, parseAPIR2, req, parseAPIRoutes)
}

func BenchmarkParseAPI_HttpRouter(b *testing.B) {
	benchmark(b, parseAPIHttpRouter, nil, parseAPIRoutes)
}

func BenchmarkParseAPI_Chi(b *testing.B) {
	benchmark(b, parseAPIChi, nil, parseAPICurlyBracketRoutes)
}

func BenchmarkParseAPI_GorillaMux(b *testing.B) {
	benchmark(b, parseAPIGorillaMux, nil, parseAPICurlyBracketRoutes)
}
