package usecases

import (
	"auth_service/internal/app/models"
	"auth_service/internal/app/usecases/mocks"
	"context"
	"errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_authUsecase_RegisterUser(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx      context.Context
		userInfo *RegisterUser
	}
	tests := []struct {
		name      string
		args      args
		wantErr   bool
		mockSetup func(m *mocks.AuthRepository)
	}{
		{
			name: "Valid registration",
			args: args{
				ctx: context.Background(),
				userInfo: &RegisterUser{
					Email:    "test@example.com",
					Password: "correctpassword",
				},
			},
			wantErr: false,
			mockSetup: func(m *mocks.AuthRepository) {
				m.On("RegisterUser", mock.Anything, mock.AnythingOfType("*models.User")).
					Return(nil).Once()
			},
		},
		{
			name: "Error registering user",
			args: args{
				ctx: context.Background(),
				userInfo: &RegisterUser{
					Email:    "test@example.com",
					Password: "correctpassword",
				},
			},
			wantErr: true,
			mockSetup: func(m *mocks.AuthRepository) {
				m.On("RegisterUser", mock.Anything, mock.AnythingOfType("*models.User")).
					Return(errors.New("registration error")).Once()
			},
		},
		{
			name: "Invalid user info",
			args: args{
				ctx:      context.Background(),
				userInfo: nil,
			},
			wantErr:   true,
			mockSetup: func(m *mocks.AuthRepository) {},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repoMock := mocks.NewAuthRepository(t)
			uc := &authUsecase{
				AuthDeps{
					AuthRepository: repoMock,
				},
			}

			tt.mockSetup(repoMock)

			_, err := uc.RegisterUser(tt.args.ctx, tt.args.userInfo)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			repoMock.AssertExpectations(t)
		})
	}
}

func Test_authUsecase_Login(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx       context.Context
		loginUser *LoginUser
	}
	tests := []struct {
		name      string
		args      args
		wantErr   bool
		mockSetup func(m *mocks.AuthRepository)
	}{
		{
			name: "Login_Success",
			args: args{
				ctx: context.Background(),
				loginUser: &LoginUser{
					Email:    "test@example.com",
					Password: "correctpassword",
				},
			},
			wantErr: false,
			mockSetup: func(m *mocks.AuthRepository) {
				user := &models.User{
					Email:          "test@example.com",
					HashedPassword: "$2a$14$Nhw00yj5zv8QVODzEGccUepuPx.FwZKhMUpCr7yR3Avltep9BPxvm", // hash
				}
				m.On("GetRegisterUserByEmail", mock.Anything, user.Email).
					Return(user, nil).Once()
			},
		},
		{
			name: "Invalid login - user not found",
			args: args{
				ctx: context.Background(),
				loginUser: &LoginUser{
					Email:    "notfound@example.com",
					Password: "somepassword",
				},
			},
			wantErr: true,
			mockSetup: func(m *mocks.AuthRepository) {
				m.On("GetRegisterUserByEmail", mock.Anything, models.Email("notfound@example.com")).
					Return(nil, errors.New("user not found")).Once()
			},
		},
		{
			name: "Invalid login - password mismatch",
			args: args{
				ctx: context.Background(),
				loginUser: &LoginUser{
					Email:    "test@example.com",
					Password: "wrongpassword",
				},
			},
			wantErr: true,
			mockSetup: func(m *mocks.AuthRepository) {
				user := &models.User{
					Email:          "test@example.com",
					HashedPassword: "$2a$10$7QWKhI2aGhRHz.QV5o/eEuVVRqszjsTpHDOQcODhtSuYkQZ7/TXC2", // hash
				}
				m.On("GetRegisterUserByEmail", mock.Anything, user.Email).
					Return(user, nil).Once()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repoMock := mocks.NewAuthRepository(t)
			uc := &authUsecase{
				AuthDeps{
					AuthRepository: repoMock,
				},
			}

			tt.mockSetup(repoMock)

			_, err := uc.Login(tt.args.ctx, tt.args.loginUser)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			repoMock.AssertExpectations(t)
		})
	}
}
