package error

import (
	"bytes"
	"github.com/goforgery/forgery2"
	. "github.com/ricallinson/simplebdd"
	"strings"
	"testing"
)

func TestCreate(t *testing.T) {

	Describe("Create()", func() {

		var app *f.Application
		var req *f.Request
		var res *f.Response
		var buf *bytes.Buffer

		BeforeEach(func() {
			app = f.CreateApp()
			req = f.CreateRequestMock(app)
			res, buf = f.CreateResponseMock(app, false)
		})

		It("should return [nil]", func() {
			app.Use(Create())
			app.Use(func(req *f.Request, res *f.Response, next func()) {
				res.End("nil")
			})
			app.Handle(req, res, 0)
			w := buf.String()
			AssertEqual(w, "nil")
		})

		It("should return [panic]", func() {
			app.Env = "prod"
			app.Use(Create())
			app.Use(func(req *f.Request, res *f.Response, next func()) {
				panic("panic")
			})
			app.Handle(req, res, 0)
			w := buf.String()
			AssertEqual(w, "panic")
		})

		It("should return [34]", func() {
			app.Env = "prod"
			app.Use(Create())
			app.Use(func(req *f.Request, res *f.Response, next func()) {
				panic("panic")
			})
			req.Header.Set("Accept", "text/html")
			app.Handle(req, res, 0)
			w := buf.String()
			AssertEqual(strings.Index(w, "<title>panic</title>"), 34)
		})

		It("should return [67]", func() {
			app.Env = "prod"
			app.Use(Create("Title"))
			app.Use(func(req *f.Request, res *f.Response, next func()) {
				panic("panic")
			})
			req.Header.Set("Accept", "text/html")
			app.Handle(req, res, 0)
			w := buf.String()
			AssertEqual(strings.Index(w, "<h1>Title</h1>"), 67)
		})

		It("should return [{\"code\":\"500\",\"error\":\"panic\"}]", func() {
			app.Env = "prod"
			app.Use(Create())
			app.Use(func(req *f.Request, res *f.Response, next func()) {
				panic("panic")
			})
			req.Header.Set("Accept", "application/json")
			app.Handle(req, res, 0)
			w := buf.String()
			AssertEqual(w, "{\"code\":\"500\",\"error\":\"panic\"}")
		})
	})

	Report(t)
}
