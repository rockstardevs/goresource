package util_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/rockstardevs/goresource/util"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Util", func() {

	Describe("WriteJSON", func() {
		var rw *httptest.ResponseRecorder

		BeforeEach(func() {
			rw = httptest.NewRecorder()
		})

		Context("given an empty string", func() {
			It("responds correctly", func() {
				util.WriteJSON("", rw)
				got, err := ioutil.ReadAll(rw.Body)
				Expect(err).To(BeNil())
				Expect(got).To(Equal([]byte("\"\"")))
				Expect(rw.Code).To(Equal(http.StatusOK))
				Expect(rw.Header().Get("Content-Type")).To(Equal("application/json"))
			})
		})
		Context("given nil", func() {
			It("responds correctly", func() {
				util.WriteJSON(nil, rw)
				got, err := ioutil.ReadAll(rw.Body)
				Expect(err).To(BeNil())
				Expect(got).To(Equal([]byte("null")))
				Expect(rw.Code).To(Equal(http.StatusOK))
				Expect(rw.Header().Get("Content-Type")).To(Equal("application/json"))
			})
		})
		Context("given an empty slice", func() {
			It("responds correctly", func() {
				util.WriteJSON(make([]int, 0), rw)
				got, err := ioutil.ReadAll(rw.Body)
				Expect(err).To(BeNil())
				Expect(got).To(Equal([]byte("[]")))
				Expect(rw.Code).To(Equal(http.StatusOK))
				Expect(rw.Header().Get("Content-Type")).To(Equal("application/json"))
			})
		})
		Context("given an empty map", func() {
			It("responds correctly", func() {
				util.WriteJSON(make(map[string]string, 0), rw)
				got, err := ioutil.ReadAll(rw.Body)
				Expect(err).To(BeNil())
				Expect(got).To(Equal([]byte("{}")))
				Expect(rw.Code).To(Equal(http.StatusOK))
				Expect(rw.Header().Get("Content-Type")).To(Equal("application/json"))
			})
		})
		Context("given a struct", func() {
			It("responds correctly", func() {
				data := struct {
					name  string
					Value string `json:"baz"`
					Alpha string
				}{name: "foo", Value: "bar", Alpha: "beta"}
				util.WriteJSON(data, rw)
				got, err := ioutil.ReadAll(rw.Body)
				Expect(err).To(BeNil())
				Expect(got).To(Equal([]byte(`{"baz":"bar","Alpha":"beta"}`)))
				Expect(rw.Code).To(Equal(http.StatusOK))
				Expect(rw.Header().Get("Content-Type")).To(Equal("application/json"))
			})
		})
		Context("given a nested struct", func() {
			It("responds correctly", func() {
				type A struct {
					Test string
				}
				data := struct {
					A
					name  string
					Value string `json:"baz"`
					Alpha string
				}{A{"thing"}, "foo", "bar", "beta"}
				util.WriteJSON(data, rw)
				got, err := ioutil.ReadAll(rw.Body)
				Expect(err).To(BeNil())
				Expect(got).To(Equal([]byte(`{"Test":"thing","baz":"bar","Alpha":"beta"}`)))
				Expect(rw.Code).To(Equal(http.StatusOK))
				Expect(rw.Header().Get("Content-Type")).To(Equal("application/json"))
			})
		})
		Context("given invalid json", func() {
			It("responds with an error", func() {
				data := struct {
					Invalid map[int]string
				}{make(map[int]string)}
				util.WriteJSON(data, rw)
				got, err := ioutil.ReadAll(rw.Body)
				Expect(err).To(BeNil())
				Expect(got).To(Equal([]byte("json: unsupported type: map[int]string\n")))
				Expect(rw.Code).To(Equal(http.StatusInternalServerError))
				Expect(rw.Header().Get("Content-Type")).To(Equal("text/plain; charset=utf-8"))
			})
		})
	})

})
