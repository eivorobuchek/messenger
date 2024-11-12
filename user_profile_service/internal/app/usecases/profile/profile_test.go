package profile

import (
	"context"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"user_profile_service/internal/app/models"
	"user_profile_service/internal/app/usecases/profile/mocks"
)

func Test_profileUsecase_EditUserProfile(t *testing.T) {
	t.Parallel()

	// ARRANGE
	var (
		ctx = context.Background() // dummy
	)

	type args struct {
		ctx     context.Context
		profile *EditProfile
	}

	tests := []struct {
		name    string
		args    args
		wantErr assert.ErrorAssertionFunc
		mock    func(t *testing.T) DepsProfile
	}{
		{
			name: "Test 1. Positive. Profile edited successfully",
			args: args{
				ctx: ctx,
				profile: &EditProfile{
					Nickname:        "user123",
					ChangedNickname: "newUser123",
					Bio:             "new bio",
					Avatar:          "new-avatar-url",
				},
			},
			mock: func(t *testing.T) DepsProfile {
				profileRepoMock := mocks.NewRepositoryProfile(t)

				// mock FindUsersByNickname
				profileRepoMock.EXPECT().
					FindUsersByNickname(ctx, models.Nickname("user123")).
					Return([]models.Profile{
						{
							Nickname: "user123",
							Bio:      "old bio",
							Avatar:   "old-avatar-url",
						},
					}, nil).Once()

				// mock EditUserProfile
				profileRepoMock.EXPECT().
					EditUserProfile(ctx, mock.Anything).
					Return(nil).Once()

				return DepsProfile{
					ProfileRepository: profileRepoMock,
				}
			},
			wantErr: assert.NoError,
		},
		{
			name: "Test 3. Negative. User not found",
			args: args{
				ctx: ctx,
				profile: &EditProfile{
					Nickname: "wrongUser",
				},
			},
			mock: func(t *testing.T) DepsProfile {
				profileRepoMock := mocks.NewRepositoryProfile(t)

				// mock FindUsersByNickname returning empty list
				profileRepoMock.EXPECT().
					FindUsersByNickname(ctx, models.Nickname("wrongUser")).
					Return([]models.Profile{}, errors.New("not found")).Once()

				return DepsProfile{
					ProfileRepository: profileRepoMock,
				}
			},
			wantErr: assert.Error,
		},
		{
			name: "Test 4. Negative. Profile validation error",
			args: args{
				ctx:     ctx,
				profile: &EditProfile{Nickname: ""}, // invalid data
			},
			mock: func(t *testing.T) DepsProfile {
				profileRepoMock := mocks.NewRepositoryProfile(t)

				return DepsProfile{
					ProfileRepository: profileRepoMock,
				}
			},
			wantErr: assert.Error,
		},
		{
			name: "Test 5. Negative. Repository error",
			args: args{
				ctx: ctx,
				profile: &EditProfile{
					Nickname:        "user123",
					ChangedNickname: "newUser123",
					Bio:             "new bio",
					Avatar:          "new-avatar-url",
				},
			},
			mock: func(t *testing.T) DepsProfile {
				profileRepoMock := mocks.NewRepositoryProfile(t)

				// mock FindUsersByNickname to return valid profile
				profileRepoMock.EXPECT().
					FindUsersByNickname(ctx, models.Nickname("user123")).
					Return([]models.Profile{
						{
							Nickname: "user123",
							Bio:      "old bio",
							Avatar:   "old-avatar-url",
						},
					}, nil).Once()

				// mock EditUserProfile to return error
				profileRepoMock.EXPECT().
					EditUserProfile(ctx, mock.Anything).
					Return(fmt.Errorf("repository error")).Once()

				return DepsProfile{
					ProfileRepository: profileRepoMock,
				}
			},
			wantErr: assert.Error,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// ARRANGE
			useCase := &profileUsecase{
				DepsProfile{
					ProfileRepository: tt.mock(t).ProfileRepository,
				},
			}

			// ACT
			err := useCase.EditUserProfile(tt.args.ctx, tt.args.profile)

			// ASSERT
			tt.wantErr(t, err)
		})
	}
}

func Test_profileUsecase_FindUsersByNickname(t *testing.T) {
	t.Parallel()

	// ARRANGE
	var (
		ctx = context.Background() // dummy
	)

	type args struct {
		ctx      context.Context
		nickname models.Nickname
	}

	tests := []struct {
		name    string
		args    args
		want    []models.Profile
		wantErr assert.ErrorAssertionFunc
		mock    func(t *testing.T) DepsProfile
	}{
		{
			name: "Test 1. Positive. Find users successfully",
			args: args{
				ctx:      ctx,
				nickname: models.Nickname("user123"),
			},
			want: []models.Profile{
				{
					Nickname: "user123",
					Bio:      "bio",
					Avatar:   "avatar-url",
				},
			},
			mock: func(t *testing.T) DepsProfile {
				profileRepoMock := mocks.NewRepositoryProfile(t)

				// mock FindUsersByNickname
				profileRepoMock.EXPECT().
					FindUsersByNickname(ctx, models.Nickname("user123")).
					Return([]models.Profile{
						{
							Nickname: "user123",
							Bio:      "bio",
							Avatar:   "avatar-url",
						},
					}, nil).Once()

				return DepsProfile{
					ProfileRepository: profileRepoMock,
				}
			},
			wantErr: assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// ARRANGE
			useCase := &profileUsecase{
				DepsProfile{
					ProfileRepository: tt.mock(t).ProfileRepository,
				},
			}

			// ACT
			got, err := useCase.FindUsersByNickname(tt.args.ctx, tt.args.nickname)

			// ASSERT
			assert.Equal(t, tt.want, got)
			tt.wantErr(t, err)
		})
	}
}
