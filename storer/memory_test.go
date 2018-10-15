package storer

import (
	"reflect"
	"testing"

	"github.com/Depado/thundermonit/domain"
	"github.com/Depado/thundermonit/uc"
)

var fixturesSlice = []*domain.Service{
	{
		ID: 0, Name: "goploader", URL: "https://gpldr.in", CI: &domain.CI{
			API: "https://drone.depado.eu",
			URL: "https://drone.depado.eu/Depado/goploader",
		}, Repo: &domain.Repo{Type: "github"},
	}, {
		ID: 1, Name: "gomonit", URL: "https://monit.depado.eu", CI: &domain.CI{
			API: "https://drone.depado.eu",
			URL: "https://drone.depado.eu/Depado/gomonit",
		},
	},
}

var fixturesMap = map[int64]*domain.Service{
	0: fixturesSlice[0],
	1: fixturesSlice[1],
}

func TestNewMemoryStorer(t *testing.T) {
	type args struct {
		ss []*domain.Service
	}
	tests := []struct {
		name string
		args args
		want uc.Storer
	}{
		{"should return storer", args{fixturesSlice}, memoryStorer{s: fixturesSlice, m: fixturesMap}},
		{"should return empty storer", args{nil}, memoryStorer{s: nil, m: make(map[int64]*domain.Service)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMemoryStorer(tt.args.ss...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMemoryStorer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_memoryStorer_GetAllServices(t *testing.T) {
	type fields struct {
		s []*domain.Service
		m map[int64]*domain.Service
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*domain.Service
		wantErr bool
	}{
		{"should return everything", fields{fixturesSlice, fixturesMap}, fixturesSlice, false},
		{"should return nil", fields{nil, nil}, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := memoryStorer{
				s: tt.fields.s,
				m: tt.fields.m,
			}
			got, err := m.GetAllServices()
			if (err != nil) != tt.wantErr {
				t.Errorf("memoryStorer.GetAllServices() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("memoryStorer.GetAllServices() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_memoryStorer_GetService(t *testing.T) {
	type fields struct {
		s []*domain.Service
		m map[int64]*domain.Service
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
		{"should find item", fields{fixturesSlice, fixturesMap}, args{0}, fixturesMap[0], false},
		{"should find item", fields{fixturesSlice, fixturesMap}, args{1}, fixturesMap[1], false},
		{"should not find item", fields{fixturesSlice, fixturesMap}, args{2}, nil, false},
		{"should not find item", fields{fixturesSlice, fixturesMap}, args{-1}, nil, false},
		{"should not find item", fields{fixturesSlice, fixturesMap}, args{1000}, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := memoryStorer{
				s: tt.fields.s,
				m: tt.fields.m,
			}
			got, err := m.GetService(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("memoryStorer.GetService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("memoryStorer.GetService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_memoryStorer_GetRepo(t *testing.T) {
	type fields struct {
		s []*domain.Service
		m map[int64]*domain.Service
	}
	type args struct {
		s *domain.Service
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.Repo
		wantErr bool
	}{
		{"should find and return a repo", fields{fixturesSlice, fixturesMap}, args{fixturesSlice[0]}, fixturesSlice[0].Repo, false},
		{"should return nil", fields{fixturesSlice, fixturesMap}, args{fixturesSlice[1]}, fixturesSlice[1].Repo, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := memoryStorer{
				s: tt.fields.s,
				m: tt.fields.m,
			}
			got, err := m.GetRepo(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("memoryStorer.GetRepo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("memoryStorer.GetRepo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_memoryStorer_GetAllRepos(t *testing.T) {
	type fields struct {
		s []*domain.Service
		m map[int64]*domain.Service
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*domain.Repo
		wantErr bool
	}{
		{"should return a slice of repo", fields{fixturesSlice, fixturesMap}, []*domain.Repo{fixturesSlice[0].Repo}, false},
		{"should return nil", fields{nil, nil}, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := memoryStorer{
				s: tt.fields.s,
				m: tt.fields.m,
			}
			got, err := m.GetAllRepos()
			if (err != nil) != tt.wantErr {
				t.Errorf("memoryStorer.GetAllRepos() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("memoryStorer.GetAllRepos() = %v, want %v", got, tt.want)
			}
		})
	}
}
