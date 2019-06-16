package controllers

import (
	"encoding/json"
	"gin-learning/core"
	"gin-learning/models"
	"gin-learning/tests"
	. "github.com/bouk/monkey"
	"github.com/gin-gonic/gin"
	. "github.com/smartystreets/goconvey/convey"
	"reflect"
	"testing"
)

func TestDemo(t *testing.T) {
	Convey("Test Demo", t, func() {
		r := gin.New()
		uri := "/demo"
		r.POST(uri, core.Handle(Demo))
		defer UnpatchAll()

		Convey("Test 缺少参数", func() {
			req := map[string]interface{}{
				"id":     "id",
				"name":   "name",
				"age":    18,
				"school": "",
			}
			body := tests.PostJson(uri, req, r)
			var resp tests.Response
			err := json.Unmarshal(body, &resp)
			So(err, ShouldBeNil)
			So(resp.Retcode, ShouldEqual, 10001)
		})

		Convey("Test NO", func() {
			var pupil *models.Pupil
			PatchInstanceMethod(reflect.TypeOf(pupil), "HasACar", func(_ *models.Pupil, _ string) bool {
				return false
			})
			req := map[string]interface{}{
				"id":     "id",
				"name":   "name",
				"age":    18,
				"school": "school",
			}
			body := tests.PostJson(uri, req, r)
			var resp tests.Response
			err := json.Unmarshal(body, &resp)
			So(err, ShouldBeNil)
			So(resp.Retcode, ShouldEqual, 10002)
		})

		Convey("Test success", func() {
			var pupil *models.Pupil
			PatchInstanceMethod(reflect.TypeOf(pupil), "HasACar", func(_ *models.Pupil, _ string) bool {
				return true
			})
			req := map[string]interface{}{
				"id":     "id",
				"name":   "name",
				"age":    18,
				"school": "school",
			}
			body := tests.PostJson(uri, req, r)
			var resp tests.Response
			err := json.Unmarshal(body, &resp)
			So(err, ShouldBeNil)
			So(resp.Retcode, ShouldEqual, 200)
		})
	})
}
