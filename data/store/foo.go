package store

import (
	rethink "github.com/gorethink/gorethink"
	db "github.com/javinc/mango/database/rethink"

	"github.com/javinc/graham/model"
)

const tableName = "foo"

func init() {
	db.Connect(db.Config{
		Host:    "localhost:28015",
		Db:      "graham",
		MaxOpen: 100,
	})

	db.CreateTable("foo")
}

func (x *store) FindFoo() ([]*model.Foo, error) {
	l := []*model.Foo{}

	db.Find(rethink.Table(tableName), &l)

	return l, nil
}

func (x *store) GetFoo(id string) (*model.Foo, error) {
	r := new(model.Foo)

	db.FindOne(rethink.Table(tableName), r)

	return r, nil
}

func (x *store) CreateFoo(r *model.Foo) (*model.Foo, error) {
	// assuming its inserted to database
	r.ID = "200"

	return r, nil
}
