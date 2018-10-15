package storer

import (
	"reflect"
	"testing"

	"github.com/Depado/thundermonit/domain"
	"github.com/Depado/thundermonit/uc"
	"github.com/jinzhu/gorm"
)

func TestNewSQLiteStorer(t *testing.T) {
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want uc.Storer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSQLiteStorer(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSQLiteStorer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqliteStorer_GetAllServices(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*domain.Service
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := sqliteStorer{
				db: tt.fields.db,
			}
			got, err := s.GetAllServices()
			if (err != nil) != tt.wantErr {
				t.Errorf("sqliteStorer.GetAllServices() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sqliteStorer.GetAllServices() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqliteStorer_GetService(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.Service
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := sqliteStorer{
				db: tt.fields.db,
			}
			got, err := s.GetService(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("sqliteStorer.GetService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sqliteStorer.GetService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqliteStorer_GetRepo(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		in0 *domain.Service
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.Repo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := sqliteStorer{
				db: tt.fields.db,
			}
			got, err := s.GetRepo(tt.args.in0)
			if (err != nil) != tt.wantErr {
				t.Errorf("sqliteStorer.GetRepo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sqliteStorer.GetRepo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqliteStorer_GetAllRepos(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*domain.Repo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := sqliteStorer{
				db: tt.fields.db,
			}
			got, err := s.GetAllRepos()
			if (err != nil) != tt.wantErr {
				t.Errorf("sqliteStorer.GetAllRepos() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sqliteStorer.GetAllRepos() = %v, want %v", got, tt.want)
			}
		})
	}
}
