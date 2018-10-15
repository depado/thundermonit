package storer

import (
	"reflect"
	"testing"

	"github.com/Depado/thundermonit/domain"
	"github.com/Depado/thundermonit/uc"
)

func TestNewNoopStorer(t *testing.T) {
	tests := []struct {
		name string
		want uc.Storer
	}{
		{"should return a new uc.Storer", NewNoopStorer()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNoopStorer(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNoopStorer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_noopStorer_GetAllServices(t *testing.T) {
	tests := []struct {
		name    string
		n       noopStorer
		want    []*domain.Service
		wantErr bool
	}{
		{"should always return nil", noopStorer{}, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := noopStorer{}
			got, err := n.GetAllServices()
			if (err != nil) != tt.wantErr {
				t.Errorf("noopStorer.GetAllServices() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("noopStorer.GetAllServices() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_noopStorer_GetService(t *testing.T) {
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		n       noopStorer
		args    args
		want    *domain.Service
		wantErr bool
	}{
		{"should always return nil", noopStorer{}, args{1}, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := noopStorer{}
			got, err := n.GetService(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("noopStorer.GetService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("noopStorer.GetService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_noopStorer_GetRepo(t *testing.T) {
	type args struct {
		in0 *domain.Service
	}
	tests := []struct {
		name    string
		n       noopStorer
		args    args
		want    *domain.Repo
		wantErr bool
	}{
		{"should always return nil", noopStorer{}, args{&domain.Service{}}, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := noopStorer{}
			got, err := n.GetRepo(tt.args.in0)
			if (err != nil) != tt.wantErr {
				t.Errorf("noopStorer.GetRepo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("noopStorer.GetRepo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_noopStorer_GetAllRepos(t *testing.T) {
	tests := []struct {
		name    string
		n       noopStorer
		want    []*domain.Repo
		wantErr bool
	}{
		{"should always return nil", noopStorer{}, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := noopStorer{}
			got, err := n.GetAllRepos()
			if (err != nil) != tt.wantErr {
				t.Errorf("noopStorer.GetAllRepos() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("noopStorer.GetAllRepos() = %v, want %v", got, tt.want)
			}
		})
	}
}
