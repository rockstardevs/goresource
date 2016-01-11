package goresource_test

import (
	"github.com/golang/mock/gomock"
	"github.com/rockstardevs/goresource"
	"github.com/rockstardevs/goresource/mocks"

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
		It("works correctly.", func() {
			Expect(manager.GetName()).To(Equal("test"))
		})
	})

	Describe(".GetEntity", func() {
		It("works correctly.", func() {
			store.EXPECT().GetEntity("test", "foo")
			manager.GetEntity("foo", nil)
		})
	})

	Describe(".CreateEntity", func() {
		It("works correctly.", func() {
			e := &mocks.MockEntity{"fakeid"}
			store.EXPECT().CreateEntity("test", e)
			manager.CreateEntity(e, nil)
		})
	})
	Describe(".ListEntities", func() {
		It("works correctly.", func() {
			store.EXPECT().ListEntities("test", nil)
			manager.ListEntities(nil)
		})
	})
	Describe(".UpdateEntity", func() {
		It("works correctly.", func() {
			e := &mocks.MockEntity{"fakeid"}
			store.EXPECT().UpdateEntity("test", "fakeid", e)
			manager.UpdateEntity("fakeid", e, nil)
		})
	})
	Describe(".DeleteEntity", func() {
		It("works correctly.", func() {
			store.EXPECT().DeleteEntity("test", "bar")
			manager.DeleteEntity("bar", nil)
		})
	})
})
