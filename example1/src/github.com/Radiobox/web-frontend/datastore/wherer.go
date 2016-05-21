package datastore

import (
	query_interfaces "github.com/nelsam/gorp_queries/interfaces"
	"github.com/stretchr/goweb/context"
)

type Queryer interface {
	Query(context.Context) (query_interfaces.SelectQuery, error)
}
