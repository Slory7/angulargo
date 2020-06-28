package repositories

import (
	"github.com/slory7/angulargo/src/infrastructure/data/db"
)

type IRepositoryReadOnly interface {
	List(slicedest interface{}, queryBoolColumns string, queryObject ...interface{}) error
	ListBy(slicePtr interface{}, query string, params ...interface{}) error
	Get(dest interface{}) (bool, error)
	GetByID(ID interface{}, dest interface{}) (bool, error)
	GetByOrder(dest interface{}, orderby string, isdecending bool) (bool, error)
	Query(slicePtr interface{}, selectquery string, params ...interface{}) error
	Count(dest interface{}, query string, params ...interface{}) (int64, error)
	Exists(dest interface{}, query string, params ...interface{}) (bool, error)
	DB() *db.Database
}

type RepositoryReadOnly struct {
	db *db.Database
}

var _ IRepositoryReadOnly = (*RepositoryReadOnly)(nil)

func NewRepositoryReadOnly(db *db.Database) *RepositoryReadOnly {
	return &RepositoryReadOnly{db}
}

func (r *RepositoryReadOnly) DB() *db.Database {
	return r.db
}

func (r *RepositoryReadOnly) List(slicedest interface{}, queryBoolColumns string, queryObject ...interface{}) error {
	err := r.db.List(slicedest, queryBoolColumns, queryObject...)
	return err
}

func (r *RepositoryReadOnly) ListBy(slicePtr interface{}, query string, params ...interface{}) error {
	err := r.db.ListBy(slicePtr, query, params...)
	return err
}

func (r *RepositoryReadOnly) Get(dest interface{}) (bool, error) {
	b, err := r.db.Get(dest)
	return b, err
}

func (r *RepositoryReadOnly) GetByID(ID interface{}, dest interface{}) (bool, error) {
	b, err := r.db.GetByID(ID, dest)
	return b, err
}

func (r *RepositoryReadOnly) GetByOrder(dest interface{}, orderby string, isdecending bool) (bool, error) {
	b, err := r.db.GetByOrder(dest, orderby, isdecending)
	return b, err
}

func (r *RepositoryReadOnly) Query(slicePtr interface{}, selectquery string, params ...interface{}) error {
	err := r.db.Query(slicePtr, selectquery, params...)
	return err
}

func (r *RepositoryReadOnly) Exists(dest interface{}, query string, params ...interface{}) (bool, error) {
	b, err := r.db.Exists(dest, query, params...)
	return b, err
}

func (r *RepositoryReadOnly) Count(dest interface{}, query string, params ...interface{}) (int64, error) {
	n, err := r.db.Count(dest, query, params...)
	return n, err
}
