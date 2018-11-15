package httplib

import (
	. "github.com/smartystreets/goconvey/convey"
	"net/url"
	"strings"
	"testing"
)

func TestGet(t *testing.T) {
	Convey("测试GET", t, func() {
		var req *HttpRequest
		var err error

		req, err = Get("http://httpbin.org/get")
		So(err, ShouldBeNil)
		_, err = req.String()
		So(err, ShouldBeNil)
	})
}

func TestPost(t *testing.T) {
	Convey("测试POST", t, func() {
		var req *HttpRequest
		var err error
		var str string

		req, err = Post("http://httpbin.org/post")
		So(err, ShouldBeNil)
		str, err = req.FormBody(url.Values{
			"name": []string{"kimi"},
		}).String()
		So(err, ShouldBeNil)
		So(strings.Index(str, "kimi"), ShouldNotEqual, -1)
	})
}

func TestResponse(t *testing.T) {
	Convey("测试Response", t, func() {
		var req *HttpRequest
		var err error

		req, err = Post("http://httpbin.org/post")
		So(err, ShouldBeNil)
		_, err = req.Response()
		So(err, ShouldBeNil)
	})
}

func TestHeader(t *testing.T) {
	Convey("测试Header", t, func() {
		var req *HttpRequest
		var err error

		req, err = Get("http://httpbin.org/header")
		So(err, ShouldBeNil)
		_, err = req.Header("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.102 Safari/537.36").String()
		So(err, ShouldBeNil)
	})
}

func TestStringBody(t *testing.T) {
	Convey("测试string Body", t, func() {
		var req *HttpRequest
		var err error

		req, err = Post("http://httpbin.org/anything")
		So(err, ShouldBeNil)

		str := "kimi"
		resp, err := req.Body(str).String()
		So(err, ShouldBeNil)
		So(strings.Index(resp, str), ShouldNotEqual, -1)
	})
}

func TestBytsBody(t *testing.T) {
	Convey("测试byte Body", t, func() {
		var req *HttpRequest
		var err error

		req, err = Post("http://httpbin.org/anything")
		So(err, ShouldBeNil)

		str := "kimi"
		resp, err := req.Body([]byte(str)).String()
		So(err, ShouldBeNil)
		So(strings.Index(resp, str), ShouldNotEqual, -1)
	})
}

func TestParseJson(t *testing.T) {
	Convey("测试ParseJson", t, func() {
		var req *HttpRequest
		var err error

		req, err = Post("http://httpbin.org/anything")
		So(err, ShouldBeNil)

		type Name struct {
			Firstname string `json:"firstname"`
			Lastname  string `json:"lastname"`
		}

		var name Name
		body := Name{
			"Kimi", "Wang",
		}

		req, err = req.JsonBody(&body)
		So(err, ShouldBeNil)
		err = req.ParseJson(&name)
		t.Log(name)
		So(err, ShouldBeNil)
		//So(name.Firstname, ShouldEqual, "Kimi")
	})
}

func TestParseXml(t *testing.T) {
	Convey("测试ParseXml", t, func() {
		var req *HttpRequest
		var err error

		req, err = Post("http://httpbin.org/anything")
		So(err, ShouldBeNil)

		type Name struct {
			Firstname string `xml:"firstname"`
			Lastname  string `xml:"lastname"`
		}

		var name Name
		body := Name{
			"Kimi", "Wang",
		}

		req, err = req.XmlBody(&body)
		So(err, ShouldBeNil)
		err = req.ParseXml(&name)
		So(err, ShouldBeNil)
		So(name.Firstname, ShouldEqual, "Kimi")
	})
}
