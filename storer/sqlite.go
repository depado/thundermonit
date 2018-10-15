package storer

import (
	"github.com/Depado/thundermonit/domain"
	"github.com/Depado/thundermonit/uc"
	"github.com/jinzhu/gorm"
)

type sqliteStorer struct {
	db *gorm.DB
}

// NewSQLiteStorer will return a new Storer which will use an SQLite database
// to store/create/modify objects
func NewSQLiteStorer(db *gorm.DB) uc.Storer {
	return sqliteStorer{
		db: db,
	}
}
func (sqliteStorer) GetAllServices() ([]*domain.Service, error) {
	panic("not implemented")
}

func (sqliteStorer) GetService(id int64) (*domain.Service, error) {
	panic("not implemented")
}

func (sqliteStorer) GetRepo(*domain.Service) (*domain.Repo, error) {
	panic("not implemented")
}

func (sqliteStorer) GetAllRepos() ([]*domain.Repo, error) {
	panic("not implemented")
}
