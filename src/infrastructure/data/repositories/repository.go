package repositories

import (
	"github.com/slory7/angulargo/src/infrastructure/data/db"
)

type IRepository interface {
	IRepositoryReadOnly
	Add(data interface{}) (affect int64, err error)
	Update(data interface{}, isAllCols bool, queryObject ...interface{}) (affect int64, err error)
	UpdateByID(ID interface{}, data interface{}, isAllCols bool) (affect int64, err error)
	AddOrUpdate(ID interface{}, data interface{}) (affect int64, err error)
	Delete(queryObject interface{}) (affect int64, err error)

	New(db *db.Database) IRepository
}

type Repository struct {
	RepositoryReadOnly
	db *db.Database
}

var _ IRepository = (*Repository)(nil)

func NewRepository(db *db.Database) *Repository {
	return &Repository{RepositoryReadOnly: RepositoryReadOnly{db}, db: db}
}

func (r *Repository) New(db *db.Database) IRepository {
	return &Repository{RepositoryReadOnly: RepositoryReadOnly{db}, db: db}
}

func (r *Repository) DB() *db.Database {
	return r.db
}

func (r *Repository) List(slicedest interface{}, boolColumn string, queryObject ...interface{}) error {
	err := r.db.List(slicedest, boolColumn, queryObject...)
	return err
}

func (r *Repository) GetByID(ID interface{}, dest interface{}) (bool, error) {
	b, err := r.db.GetByID(ID, dest)
	return b, err
}

func (r *Repository) Add(data interface{}) (affect int64, err error) {
	affect, err = r.db.Create(data)
	return
}

func (r *Repository) Update(data interface{}, isAllCols bool, queryObject ...interface{}) (affect int64, err error) {
	affect, err = r.db.Update(data, isAllCols, queryObject...)
	return affect, err
}

func (r *Repository) UpdateByID(ID interface{}, data interface{}, isAllCols bool) (affect int64, err error) {
	affect, err = r.db.UpdateByID(ID, data, isAllCols)
	return affect, err
}

func (r *Repository) AddOrUpdate(ID interface{}, data interface{}) (affect int64, err error) {
	affect, err = r.db.AddOrUpdate(ID, data)
	return
}

func (r *Repository) Delete(queryObject interface{}) (affect int64, err error) {
	affect, err = r.db.Delete(queryObject)
	return affect, err
}
