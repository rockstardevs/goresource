package goresource_test

import (
	"fmt"
	"goresource"
	"goresource/mocks"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DefaultManager", func() {
	var (
		ctrl    *gomock.Controller
		manager goresource.DefaultManager
		store   *mocks.MockStore
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		store = mocks.NewMockStore(ctrl)
		manager = goresource.NewDefaultManager("test", store)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Describe(".GetName", func() {
		It("returns the manager name.", func() {
			Expect(manager.GetName()).To(Equal("test"))
		})
	})

	Describe(".GetEntity", func() {
		It("returns the fetched entity from the store.", func() {
			want := map[string]interface{}{"bar": "baz"}
			store.EXPECT().GetEntity("test", "foo", gomock.Any()).Times(1).SetArg(2, want).Return(nil)
			got, err := manager.GetEntity("foo", nil)
			Expect(err).To(BeNil())
			Expect(got.(map[string]interface{})["bar"]).To(Equal("baz"))
		})
		It("passes through any errors from the store.", func() {
			e := fmt.Errorf("test error")
			store.EXPECT().GetEntity("test", "foo", gomock.Any()).Times(1).Return(e)
			got, err := manager.GetEntity("foo", nil)
			Expect(err).ToNot(BeNil())
			Expect(err.Error()).To(Equal("test error"))
			Expect(got).To(BeNil())
		})
	})

	Describe(".CreateEntity", func() {
		It("creates a database entity and returns the created entity.", func() {
			want := map[string]interface{}{"bar": "baz"}
			e := &mocks.MockEntity{"fakeid"}
			store.EXPECT().CreateEntity("test", e, gomock.Any()).Times(1).SetArg(2, want).Return(nil)
			got, err := manager.CreateEntity(e, nil)
			Expect(err).To(BeNil())
			Expect(got.(map[string]interface{})["bar"]).To(Equal("baz"))
		})
		It("passes through any errors from the store.", func() {
			e := fmt.Errorf("test error")
			store.EXPECT().CreateEntity("test", nil, gomock.Any()).Times(1).Return(e)
			got, err := manager.CreateEntity(nil, nil)
			Expect(err).ToNot(BeNil())
			Expect(err.Error()).To(Equal("test error"))
			Expect(got).To(BeNil())
		})
	})
	Describe(".ListEntities", func() {
		It("returns the fetched entities from the store.", func() {
			want := []map[string]interface{}{{"item1": "value1"}, {"item2": "value2"}}
			store.EXPECT().ListEntities("test", nil, gomock.Any()).Times(1).SetArg(2, want).Return(nil)
			result, err := manager.ListEntities(nil)
			Expect(err).To(BeNil())
			got, ok := result.([]map[string]interface{})
			Expect(ok).To(BeTrue())
			Expect(len(got)).To(Equal(2))
			Expect(got[0]["item1"]).To(Equal("value1"))
			Expect(got[1]["item2"]).To(Equal("value2"))
		})
		It("passes through any errors from the store.", func() {
			e := fmt.Errorf("test error")
			store.EXPECT().ListEntities("test", nil, gomock.Any()).Times(1).Return(e)
			got, err := manager.ListEntities(nil)
			Expect(err.Error()).To(Equal("test error"))
			Expect(got).To(BeNil())
		})
	})
	Describe(".UpdateEntity", func() {
		It("updates the database entity and returns it.", func() {
			want := map[string]interface{}{"bar": "baz"}
			e := &mocks.MockEntity{"fakeid"}
			store.EXPECT().UpdateEntity("test", "fakeid", e, gomock.Any()).Times(1).SetArg(3, want).Return(nil)
			got, err := manager.UpdateEntity("fakeid", e, nil)
			Expect(err).To(BeNil())
			Expect(got).To(BeEquivalentTo(want))
		})
		It("passes through any errors from the store.", func() {
			e := &mocks.MockEntity{"fakeid"}
			er := fmt.Errorf("test error")
			store.EXPECT().UpdateEntity("test", "fakeid", e, gomock.Any()).Times(1).Return(er)
			got, err := manager.UpdateEntity("fakeid", e, nil)
			Expect(err.Error()).To(Equal("test error"))
			Expect(got).To(BeNil())
		})
	})
	Describe(".DeleteEntity", func() {
		It("deletes the corresponding entity from the store.", func() {
			store.EXPECT().DeleteEntity("test", "bar").Times(1).Return(nil)
			err := manager.DeleteEntity("bar", nil)
			Expect(err).To(BeNil())
		})
		It("passes through any errors from the store.", func() {
			e := fmt.Errorf("test error")
			store.EXPECT().DeleteEntity("test", "bar").Times(1).Return(e)
			err := manager.DeleteEntity("bar", nil)
			Expect(err).ToNot(BeNil())
			Expect(err.Error()).To(Equal("test error"))
		})
	})
})
