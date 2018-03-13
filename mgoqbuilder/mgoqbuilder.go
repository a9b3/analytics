package mgoqbuilder

import (
	"net/url"
	"strconv"
	"strings"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func valueToBson(s string) interface{} {
	return s
}

func mgoQFromURLQ(q url.Values) map[string]interface{} {
	mgoQuery := bson.M{}
	for k, _ := range q {
		switch k {
		case
			"sort",
			"limit":
			continue
		}
		mgoQuery[k] = valueToBson(q.Get(k))
	}
	return mgoQuery
}

// BuildQuery will return a mgo.Query after parsing the url querystring for some
// special keys. "sort" and "limit"
// ex. ?sort=date,-age&limit=50
func BuildQuery(c *mgo.Collection, q url.Values) *mgo.Query {
	query := c.Find(mgoQFromURLQ(q))

	sortBys := strings.Split(q.Get("sort"), ",")
	if len(sortBys) > 0 && sortBys[0] != "" {
		query = query.Sort(sortBys...)
	}

	limit := q.Get("limit")
	limitInt, err := strconv.Atoi(limit)
	if err == nil {
		query = query.Limit(limitInt)
	}

	return query
}
