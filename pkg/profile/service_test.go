//go:build unit

package profile_test

import (
	"context"
	"github.com/stretchr/testify/mock"
	"maker-checker/logger"
	mocks "maker-checker/pkg/mocks/profile"
	"maker-checker/pkg/profile"
	"reflect"
	"testing"
	"time"
)

func TestNewService(t *testing.T) {
	type args struct {
		log  logger.ContextLogger
		repo *mocks.Repository
	}
	tests := []struct {
		name string
		args args
		want *profile.Service
	}{
		{
			name: "should return new service",
			args: args{
				log:  logger.NoOpContextLogger{},
				repo: &mocks.Repository{},
			},
			want: profile.NewService(logger.NoOpContextLogger{}, &mocks.Repository{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := profile.NewService(tt.args.log, tt.args.repo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_CreateProfile(t *testing.T) {
	type fields struct {
		log  logger.ContextLogger
		repo *mocks.Repository
	}
	type args struct {
		ctx   context.Context
		input profile.Row
	}

	log := logger.NoOpContextLogger{}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "should return new profileID",
			fields: fields{
				log:  log,
				repo: &mocks.Repository{},
			},
			args: args{
				ctx: context.TODO(),
				input: profile.Row{
					FirstName: "John",
					LastName:  "Doe",
					Email:     "johndoe@example.com",
				},
			},
			want:    "123",
			wantErr: false,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := profile.NewService(tt.fields.log, tt.fields.repo)
			tt.fields.repo.On("Create", tt.args.ctx, mock.AnythingOfType("profile.Row")).Return("123", nil)
			got, err := s.CreateProfile(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateProfile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CreateProfile() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_FindByID(t *testing.T) {
	type fields struct {
		log  logger.ContextLogger
		repo *mocks.Repository
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    profile.Row
		wantErr bool
	}{
		{
			name: "should return profile",
			fields: fields{
				log:  logger.NoOpContextLogger{},
				repo: &mocks.Repository{},
			},
			args: args{
				ctx: context.TODO(),
				id:  "123",
			},
			want: profile.Row{
				ID:        "123",
				FirstName: "Su",
				LastName:  "Shi",
				Email:     "sushi@example.com",
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := profile.NewService(tt.fields.log, tt.fields.repo)
			tt.fields.repo.On("FindByID", tt.args.ctx, tt.args.id).Return(&tt.want, nil)
			got, err := s.FindByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}
