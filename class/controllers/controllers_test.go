//go:generate mockgen -package=mock_controllers -destination=./mocks/controllers_mock.go  -source=controllers.go github.com/JimySheepman/class/controllers

package controllers

import (
	"context"
	"errors"
	"testing"

	mock_controllers "github.com/JimySheepman/class/controllers/mocks"
	"github.com/JimySheepman/class/dto"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type localAdminServicesMocks struct {
	mockAdminServices *mock_controllers.MockadminServices
}

func _setupAdminControllerTest_(t *testing.T) (*localAdminServicesMocks, *adminController) {
	ctx := context.Background()
	ctrl, ctx := gomock.WithContext(ctx, t)

	mocks := &localAdminServicesMocks{
		mockAdminServices: mock_controllers.NewMockadminServices(ctrl),
	}

	repo := NewAdminControllers(mocks.mockAdminServices)

	return mocks, repo
}

func TestAdminController_ReadHandler(t *testing.T) {
	mocks, repo := _setupAdminControllerTest_(t)

	tests := []struct {
		name         string
		req          *dto.RestRequest
		expected     *dto.RestResponse
		isError      bool
		expectations func()
	}{
		{
			name:     "adminServices.Read failure",
			req:      &dto.RestRequest{},
			expected: nil,
			isError:  true,
			expectations: func() {
				mocks.mockAdminServices.EXPECT().Read(gomock.Any()).Return(errors.New("error"))
			},
		},
		{
			name:     "adminServices.Write failure",
			req:      &dto.RestRequest{},
			expected: nil,
			isError:  true,
			expectations: func() {
				mocks.mockAdminServices.EXPECT().Read(gomock.Any()).Return(nil)
				mocks.mockAdminServices.EXPECT().Write(gomock.Any()).Return(nil, errors.New("error"))
			},
		},
		{
			name:     "ReadHandler succeed",
			req:      &dto.RestRequest{},
			expected: &dto.RestResponse{},
			isError:  false,
			expectations: func() {
				mocks.mockAdminServices.EXPECT().Read(gomock.Any()).Return(nil)
				mocks.mockAdminServices.EXPECT().Write(gomock.Any()).Return(&dto.RestResponse{}, nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.expectations()

			actual, err := repo.ReadHandler(tt.req)

			if tt.isError {
				assert.Error(t, err)
				assert.Nil(t, actual)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.expected, actual)
			}
		})
	}
}
