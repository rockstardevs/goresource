package store_test

import (
	"goresource/store"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type TestItem struct {
	ID   bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string        `bson:"name"`
	Tag  string        `bson:"tag"`
}

var _ = Describe("MongoStore", func() {
	var (
		session  *mgo.Session
		database *mgo.Database
		//collection *mgo.Collection
		//query      *mgo.Query
		testdbhost = "127.0.0.1"
		testdbname = "goresourcetestdatabase"
		testcoll   = "testitems"
	)

	BeforeEach(func() {
		session, _ = mgo.Dial(testdbhost)
		database = session.DB(testdbname)
		database.DropDatabase()
	})

	Describe("NewMongoStore", func() {
		It("returns an initialized database.", func() {
			s, err := store.NewMongoStore(testdbhost, testdbname, 5*time.Second)
			defer s.Close()
			Expect(err).To(BeNil())
			Expect(s).ToNot(BeNil())
		})
		It("passes errors through.", func() {
			s, err := store.NewMongoStore(testdbhost, testdbname, 1*time.Nanosecond)
			Expect(err).ToNot(BeNil())
			Expect(s).To(BeNil())
		})
	})

	Describe("ListEntities", func() {
		var (
			s   store.Store
			err error
		)

		BeforeEach(func() {
			s, err = store.NewMongoStore(testdbhost, testdbname, 5*time.Second)
		})

		AfterEach(func() {
			s.Close()
		})

		Context("if entities exist in the database.", func() {
			BeforeEach(func() {
				database.C(testcoll).Insert(
					TestItem{Name: "item1", Tag: ""},
					TestItem{Name: "item2", Tag: "imp"},
					TestItem{Name: "item3", Tag: "imp"},
					TestItem{Name: "item4", Tag: ""})
			})
			It("fetches all entities not given a filter", func() {
				var items []TestItem
				err := s.ListEntities(testcoll, nil, &items)
				Expect(err).To(BeNil())
				Expect(len(items)).To(Equal(4))
				Expect(items[0].Name).To(Equal("item1"))
				Expect(items[1].Name).To(Equal("item2"))
				Expect(items[2].Name).To(Equal("item3"))
				Expect(items[3].Name).To(Equal("item4"))
			})
			It("fetches all entities given a filter", func() {
				var items []TestItem
				err := s.ListEntities(testcoll, bson.M{"tag": "imp"}, &items)
				Expect(err).To(BeNil())
				Expect(len(items)).To(Equal(2))
				Expect(items[0].Name).To(Equal("item2"))
				Expect(items[1].Name).To(Equal("item3"))
			})
		})

		Context("if no entities exist in the database.", func() {
			It("returns an empty slice not given a filter", func() {
				var items []TestItem
				err := s.ListEntities(testcoll, nil, &items)
				Expect(err).To(BeNil())
				Expect(len(items)).To(Equal(0))
			})
			It("returns an empty slice given a filter", func() {
				var items []TestItem
				err := s.ListEntities(testcoll, bson.M{"tag": "imp"}, &items)
				Expect(err).To(BeNil())
				Expect(len(items)).To(Equal(0))
			})
		})
	})

	Describe("GetEntity", func() {
		var (
			s   store.Store
			err error
		)

		BeforeEach(func() {
			s, err = store.NewMongoStore(testdbhost, testdbname, 5*time.Second)
		})

		AfterEach(func() {
			s.Close()
		})

		Context("given a valid id", func() {
			It("fetches the entity.", func() {
				var result TestItem
				source := TestItem{Name: "foo", Tag: "bar", ID: bson.NewObjectId()}
				if err := database.C(testcoll).Insert(source); err != nil {
					Fail(err.Error())
				}
				err := s.GetEntity(testcoll, source.ID.Hex(), &result)
				Expect(err).To(BeNil())
				Expect(result.ID).To(Equal(source.ID))
				Expect(result.Name).To(Equal(source.Name))
				Expect(result.Tag).To(Equal(source.Tag))
			})
		})

		Context("given an invalid id", func() {
			It("returns an error.", func() {
				var result TestItem
				source := TestItem{Name: "foo", Tag: "bar", ID: bson.NewObjectId()}
				err := s.GetEntity(testcoll, source.ID.Hex(), &result)
				Expect(err).ToNot(BeNil())
			})
		})
	})

	Describe("CreateEntity", func() {
		var (
			s   store.Store
			err error
		)

		BeforeEach(func() {
			s, err = store.NewMongoStore(testdbhost, testdbname, 5*time.Second)
		})

		AfterEach(func() {
			s.Close()
		})

		Context("given a valid entity.", func() {
			It("persists it in the database.", func() {
				var result TestItem
				item := TestItem{Name: "foo", Tag: "bar"}
				err := s.CreateEntity(testcoll, item, &result)
				Expect(err).To(BeNil())
				Expect(result.ID.Valid()).To(BeTrue())
			})
		})

		Context("given an invalid entity.", func() {
			It("returns an error.", func() {
				var result TestItem
				err := s.CreateEntity(testcoll, nil, &result)
				Expect(err).ToNot(BeNil())
				Expect(result.ID.Valid()).To(BeFalse())
			})
		})
	})

	Describe("UpdateEntity", func() {
		var (
			s   store.Store
			err error
		)

		BeforeEach(func() {
			s, err = store.NewMongoStore(testdbhost, testdbname, 5*time.Second)
		})

		AfterEach(func() {
			s.Close()
		})

		Context("given a valid entity.", func() {
			It("updates it in the database.", func() {
				var result TestItem
				source := TestItem{ID: bson.NewObjectId(), Name: "foo", Tag: ""}
				changed := TestItem{Name: "bar", Tag: "baz"}
				if err := database.C(testcoll).Insert(source); err != nil {
					Fail(err.Error())
				}
				err := s.UpdateEntity(testcoll, source.ID.Hex(), changed, &result)
				Expect(err).To(BeNil())
				Expect(result.ID.Valid()).To(BeTrue())
				Expect(result.ID).To(Equal(source.ID))
				Expect(result.Name).To(Equal("bar"))
				Expect(result.Tag).To(Equal("baz"))
			})
		})

		Context("given a non existent entity.", func() {
			It("returns an error.", func() {
				var result TestItem
				changed := TestItem{Name: "foo", Tag: "bar"}
				id := bson.NewObjectId().Hex()
				err := s.UpdateEntity(testcoll, id, changed, &result)
				Expect(err).ToNot(BeNil())
				Expect(result.ID.Valid()).To(BeFalse())
			})
		})
	})

	Describe("DeleteEntity", func() {
		var (
			s   store.Store
			err error
		)

		BeforeEach(func() {
			s, err = store.NewMongoStore(testdbhost, testdbname, 5*time.Second)
		})

		AfterEach(func() {
			s.Close()
		})

		Context("given a valid entity.", func() {
			It("removes it from the database.", func() {
				var result TestItem
				source := TestItem{ID: bson.NewObjectId(), Name: "foo", Tag: ""}
				if err := database.C(testcoll).Insert(source); err != nil {
					Fail(err.Error())
				}
				err := s.DeleteEntity(testcoll, source.ID.Hex())
				Expect(err).To(BeNil())
				err = database.C(testcoll).FindId(source.ID).One(&result)
				Expect(err).ToNot(BeNil())
			})
		})

		Context("given a non existent entity.", func() {
			It("returns an error.", func() {
				source := TestItem{ID: bson.NewObjectId(), Name: "foo", Tag: ""}
				err := s.DeleteEntity(testcoll, source.ID.Hex())
				Expect(err).ToNot(BeNil())
			})
		})
	})

})
