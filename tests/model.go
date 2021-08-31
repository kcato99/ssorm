package tests

import (
	"time"

	"cloud.google.com/go/spanner"
)

type Singers struct {
	SingerId        int64  `ssorm_key:"primary"`
	FirstName       string `spanner:"FirstName"`
	LastName        spanner.NullString
	TestTime        spanner.NullTime     `spanner:"TestTime"` //NULL を許容する場合必ず、spanner.NullTimeを指定すること
	TestSpannerTime spanner.NullTime     `spanner:"TestSpannerTime"`
	TagIDs          []spanner.NullString `spanner:"TagIds"`
	Numbers         []int64              `spanner:"Numbers"`
	DeleteTime      spanner.NullTime     `spanner:"DeleteTime" ssorm_key:"delete_time"`
	CreateTime      time.Time            `spanner:"CreateTime" ssorm_key:"create_time"`
	UpdateTime      time.Time            `spanner:"UpdateTime" ssorm_key:"update_time"`
}

type Singer struct {
	SingerId   int64 `key:"primary"`
	FirstName  string
	LastName   string
	Albums     []*Albums
	Concerts   []*Concerts
	DeleteTime spanner.NullTime `spanner:"DeleteTime" ssorm_key:"delete_time"`
}

type Albums struct {
	SingerId   int64
	AlbumId    int64
	Title      string
	DeleteTime spanner.NullTime `spanner:"DeleteTime" ssorm_key:"delete_time"`
}
type Concerts struct {
	SingerId   int64
	ConcertId  int64
	Price      int64
	DeleteTime spanner.NullTime `spanner:"DeleteTime" ssorm_key:"delete_time"`
}
