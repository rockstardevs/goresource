package goresource_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/rockstardevs/goresource"
	"github.com/rockstardevs/goresource/mocks"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Resource", func() {
	var (
		ctrl    *gomock.Controller
		manager *mocks.MockResourceManager
		router  *mux.Router
		r       *goresource.Resource
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		manager = mocks.NewMockResourceManager(ctrl)
		router = mux.NewRouter()
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("at initialization", func() {
		It("sets up routes correctly", func() {
			manager.EXPECT().GetName().Times(2).Return("foo")
			r = goresource.NewResource(manager, router)
		})
	})
})

var _ = Describe("Resource.Get", func() {
	var (
		ctrl    *gomock.Controller
		manager *mocks.MockResourceManager
		router  *mux.Router
		r       *goresource.Resource
		rw      *httptest.ResponseRecorder
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		manager = mocks.NewMockResourceManager(ctrl)
		router = mux.NewRouter().PathPrefix("/api").Subrouter()
		rw = httptest.NewRecorder()
		manager.EXPECT().GetName().AnyTimes().Return("test")
		r = goresource.NewResource(manager, router)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("given a valid id", func() {
		It("responds with the corresponding entity.", func() {
			req, _ := http.NewRequest("GET", "/api/test/fakeid", nil)
			manager.EXPECT().GetEntity("fakeid", req.URL.Query()).Return("fake-entity", nil)
			router.ServeHTTP(rw, req)
			Expect(rw.Code).To(Equal(http.StatusOK))
			Expect(rw.Header().Get("Content-Type")).To(Equal("application/json"))
			Expect(rw.Body.String()).To(Equal(`"fake-entity"`))
		})
		It("responds with an error, if one occurs.", func() {
			req, _ := http.NewRequest("GET", "/api/test/fakeid", nil)
			manager.EXPECT().GetEntity("fakeid", req.URL.Query()).Return(nil, fmt.Errorf("Test Error"))
			router.ServeHTTP(rw, req)
			Expect(rw.Code).To(Equal(http.StatusInternalServerError))
			Expect(rw.Header().Get("Content-Type")).To(Equal("text/plain; charset=utf-8"))
			Expect(rw.Body.String()).To(Equal("Test Error\n"))
		})
	})
	Context("not given an id", func() {
		It("responds with all entities.", func() {
			req, _ := http.NewRequest("GET", "/api/test", nil)
			manager.EXPECT().ListEntities(req.URL.Query()).Return([]interface{}{"entities"}, nil)
			router.ServeHTTP(rw, req)
			Expect(rw.Code).To(Equal(http.StatusOK))
			Expect(rw.Header().Get("Content-Type")).To(Equal("application/json"))
			Expect(rw.Body.String()).To(Equal(`["entities"]`))
		})
		It("responds with an error, if one occurs.", func() {
			req, _ := http.NewRequest("GET", "/api/test", nil)
			manager.EXPECT().ListEntities(req.URL.Query()).Return(nil, fmt.Errorf("Test Error"))
			router.ServeHTTP(rw, req)
			Expect(rw.Code).To(Equal(http.StatusInternalServerError))
			Expect(rw.Header().Get("Content-Type")).To(Equal("text/plain; charset=utf-8"))
			Expect(rw.Body.String()).To(Equal("Test Error\n"))
		})
	})
})

var _ = Describe("Resource.Head", func() {
	var (
		ctrl    *gomock.Controller
		manager *mocks.MockResourceManager
		router  *mux.Router
		r       *goresource.Resource
		rw      *httptest.ResponseRecorder
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		manager = mocks.NewMockResourceManager(ctrl)
		router = mux.NewRouter().PathPrefix("/api").Subrouter()
		rw = httptest.NewRecorder()
		manager.EXPECT().GetName().AnyTimes().Return("test")
		r = goresource.NewResource(manager, router)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("given a valid id", func() {
		It("responds with the correct headers.", func() {
			req, _ := http.NewRequest("HEAD", "/api/test/fakeid", nil)
			manager.EXPECT().GetEntity("fakeid", req.URL.Query()).Return("fake-entity", nil)
			router.ServeHTTP(rw, req)
			Expect(rw.Code).To(Equal(http.StatusOK))
			Expect(rw.Header().Get("Content-Type")).To(Equal("application/json"))
			Expect(rw.Body.String()).To(Equal(""))
		})
		It("responds with an error, if one occurs.", func() {
			req, _ := http.NewRequest("HEAD", "/api/test/fakeid", nil)
			manager.EXPECT().GetEntity("fakeid", req.URL.Query()).Return(nil, fmt.Errorf("Test Error"))
			router.ServeHTTP(rw, req)
			Expect(rw.Code).To(Equal(http.StatusInternalServerError))
			Expect(rw.Header().Get("Content-Type")).To(Equal("text/plain; charset=utf-8"))
			Expect(rw.Body.String()).To(Equal("Test Error\n"))
		})
	})
	Context("not given an id", func() {
		It("responds with all entities.", func() {
			req, _ := http.NewRequest("HEAD", "/api/test", nil)
			manager.EXPECT().ListEntities(req.URL.Query()).Return([]interface{}{"entities"}, nil)
			router.ServeHTTP(rw, req)
			Expect(rw.Code).To(Equal(http.StatusOK))
			Expect(rw.Header().Get("Content-Type")).To(Equal("application/json"))
			Expect(rw.Body.String()).To(Equal(""))
		})
		It("responds with an error, if one occurs.", func() {
			req, _ := http.NewRequest("HEAD", "/api/test", nil)
			manager.EXPECT().ListEntities(req.URL.Query()).Return(nil, fmt.Errorf("Test Error"))
			router.ServeHTTP(rw, req)
			Expect(rw.Code).To(Equal(http.StatusInternalServerError))
			Expect(rw.Header().Get("Content-Type")).To(Equal("text/plain; charset=utf-8"))
			Expect(rw.Body.String()).To(Equal("Test Error\n"))
		})
	})
})

var _ = Describe("Resource.PostOrPut", func() {
	var (
		ctrl    *gomock.Controller
		manager *mocks.MockResourceManager
		router  *mux.Router
		r       *goresource.Resource
		rw      *httptest.ResponseRecorder
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		manager = mocks.NewMockResourceManager(ctrl)
		router = mux.NewRouter().PathPrefix("/api").Subrouter()
		rw = httptest.NewRecorder()
		manager.EXPECT().GetName().AnyTimes().Return("test")
		r = goresource.NewResource(manager, router)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("given an id in the URI", func() {
		It("updates the given entity, requesting via PUT.", func() {
			body := ioutil.NopCloser(strings.NewReader("fake-content"))
			e := &mocks.MockEntity{"fake-entity"}
			req, _ := http.NewRequest("PUT", "/api/test/fakeid", body)
			manager.EXPECT().ParseJSON(body).Return(e, nil)
			manager.EXPECT().UpdateEntity("fakeid", e, req.URL.Query()).Return("fake-entity", nil)
			router.ServeHTTP(rw, req)
			Expect(rw.Code).To(Equal(http.StatusOK))
			Expect(rw.Header().Get("Content-Type")).To(Equal("application/json"))
			Expect(rw.Body.String()).To(Equal(`"fake-entity"`))
		})
		It("updates the given entity, requesting via POST.", func() {
			body := ioutil.NopCloser(strings.NewReader("fake-content"))
			e := &mocks.MockEntity{"fakeid"}
			req, _ := http.NewRequest("POST", "/api/test", body)
			manager.EXPECT().ParseJSON(body).Return(e, nil)
			manager.EXPECT().UpdateEntity("fakeid", e, req.URL.Query()).Return("fake-entity", nil)
			router.ServeHTTP(rw, req)
			Expect(rw.Code).To(Equal(http.StatusOK))
			Expect(rw.Header().Get("Content-Type")).To(Equal("application/json"))
			Expect(rw.Body.String()).To(Equal(`"fake-entity"`))
		})
		It("creates the given entity, requesting via POST.", func() {
			body := ioutil.NopCloser(strings.NewReader("fake-content"))
			e := &mocks.MockEntity{}
			req, _ := http.NewRequest("POST", "/api/test", body)
			manager.EXPECT().ParseJSON(body).Return(e, nil)
			manager.EXPECT().CreateEntity(e, req.URL.Query()).Return("fake-entity", nil)
			router.ServeHTTP(rw, req)
			Expect(rw.Code).To(Equal(http.StatusOK))
			Expect(rw.Header().Get("Content-Type")).To(Equal("application/json"))
			Expect(rw.Body.String()).To(Equal(`"fake-entity"`))
		})
		It("creates the given entity, requesting via PUT.", func() {
			body := ioutil.NopCloser(strings.NewReader("fake-content"))
			e := &mocks.MockEntity{}
			req, _ := http.NewRequest("PUT", "/api/test", body)
			manager.EXPECT().ParseJSON(body).Return(e, nil)
			manager.EXPECT().CreateEntity(e, req.URL.Query()).Return("fake-entity", nil)
			router.ServeHTTP(rw, req)
			Expect(rw.Code).To(Equal(http.StatusOK))
			Expect(rw.Header().Get("Content-Type")).To(Equal("application/json"))
			Expect(rw.Body.String()).To(Equal(`"fake-entity"`))
		})
		It("responds with an error, if given invalid json via PUT.", func() {
			body := ioutil.NopCloser(strings.NewReader("fake-content"))
			req, _ := http.NewRequest("PUT", "/api/test/fakeid", body)
			// TODO: figure out why this doesn't work without matching Any().
			manager.EXPECT().ParseJSON(body).Return(nil, fmt.Errorf("test error"))
			router.ServeHTTP(rw, req)
			Expect(rw.Code).To(Equal(http.StatusBadRequest))
			Expect(rw.Header().Get("Content-Type")).To(Equal("text/plain; charset=utf-8"))
			Expect(rw.Body.String()).To(Equal("test error\n"))
		})
		It("responds with an error, if given invalid json via POST.", func() {
			body := ioutil.NopCloser(strings.NewReader("fake-content"))
			req, _ := http.NewRequest("POST", "/api/test", body)
			manager.EXPECT().ParseJSON(body).Return(nil, fmt.Errorf("test error"))
			router.ServeHTTP(rw, req)
			Expect(rw.Code).To(Equal(http.StatusBadRequest))
			Expect(rw.Header().Get("Content-Type")).To(Equal("text/plain; charset=utf-8"))
			Expect(rw.Body.String()).To(Equal("test error\n"))
		})
		It("responds with an error, if one occurs, via PUT.", func() {
			body := ioutil.NopCloser(strings.NewReader("fake-content"))
			e := &mocks.MockEntity{"fake-entity"}
			req, _ := http.NewRequest("PUT", "/api/test/fakeid", body)
			manager.EXPECT().ParseJSON(body).Return(e, nil)
			manager.EXPECT().UpdateEntity("fakeid", e, req.URL.Query()).Return(nil, fmt.Errorf("test error"))
			router.ServeHTTP(rw, req)
			Expect(rw.Code).To(Equal(http.StatusInternalServerError))
			Expect(rw.Header().Get("Content-Type")).To(Equal("text/plain; charset=utf-8"))
			Expect(rw.Body.String()).To(Equal("test error\n"))
		})
		It("responds with an error, if one occurs, via POST.", func() {
			body := ioutil.NopCloser(strings.NewReader("fake-content"))
			e := &mocks.MockEntity{"fakeid"}
			req, _ := http.NewRequest("POST", "/api/test", body)
			manager.EXPECT().ParseJSON(body).Return(e, nil)
			manager.EXPECT().UpdateEntity("fakeid", e, req.URL.Query()).Return(nil, fmt.Errorf("test error"))
			router.ServeHTTP(rw, req)
			Expect(rw.Code).To(Equal(http.StatusInternalServerError))
			Expect(rw.Header().Get("Content-Type")).To(Equal("text/plain; charset=utf-8"))
			Expect(rw.Body.String()).To(Equal("test error\n"))
		})
	})
})

var _ = Describe("Resource.Delete", func() {
	var (
		ctrl    *gomock.Controller
		manager *mocks.MockResourceManager
		router  *mux.Router
		r       *goresource.Resource
		rw      *httptest.ResponseRecorder
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		manager = mocks.NewMockResourceManager(ctrl)
		router = mux.NewRouter().PathPrefix("/api").Subrouter()
		rw = httptest.NewRecorder()
		manager.EXPECT().GetName().AnyTimes().Return("test")
		r = goresource.NewResource(manager, router)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("given a valid id", func() {
		It("deletes the corresponding entity.", func() {
			req, _ := http.NewRequest("DELETE", "/api/test/fakeid", nil)
			manager.EXPECT().DeleteEntity("fakeid", req.URL.Query()).Return(nil)
			router.ServeHTTP(rw, req)
			Expect(rw.Code).To(Equal(http.StatusNoContent))
			Expect(rw.Body.String()).To(Equal(""))
		})
		It("responds with an error, if one occurs.", func() {
			req, _ := http.NewRequest("DELETE", "/api/test/fakeid", nil)
			manager.EXPECT().DeleteEntity("fakeid", req.URL.Query()).Return(fmt.Errorf("Test Error"))
			router.ServeHTTP(rw, req)
			Expect(rw.Code).To(Equal(http.StatusInternalServerError))
			Expect(rw.Header().Get("Content-Type")).To(Equal("text/plain; charset=utf-8"))
			Expect(rw.Body.String()).To(Equal("Test Error\n"))
		})
	})
	Context("not given an id", func() {
		It("responds with an error", func() {
			req, _ := http.NewRequest("DELETE", "/api/test", nil)
			router.ServeHTTP(rw, req)
			Expect(rw.Code).To(Equal(http.StatusBadRequest))
			Expect(rw.Header().Get("Content-Type")).To(Equal("text/plain; charset=utf-8"))
			Expect(rw.Body.String()).To(Equal("Invalid Id\n"))
		})
	})
})

var _ = Describe("Resource.UnsupportedMethod", func() {
	var (
		ctrl    *gomock.Controller
		manager *mocks.MockResourceManager
		router  *mux.Router
		r       *goresource.Resource
		rw      *httptest.ResponseRecorder
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		manager = mocks.NewMockResourceManager(ctrl)
		router = mux.NewRouter().PathPrefix("/api").Subrouter()
		rw = httptest.NewRecorder()
		manager.EXPECT().GetName().AnyTimes().Return("test")
		r = goresource.NewResource(manager, router)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	It("responds with an error", func() {
		req, _ := http.NewRequest("FAKE", "/api/test", nil)
		router.ServeHTTP(rw, req)
		Expect(rw.Code).To(Equal(http.StatusNotImplemented))
		Expect(rw.Header().Get("Content-Type")).To(Equal("text/plain; charset=utf-8"))
		Expect(rw.Body.String()).To(Equal("Method Not Supported\n"))
	})
})
