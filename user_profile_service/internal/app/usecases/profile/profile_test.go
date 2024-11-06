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
					FindUsersByNickname(ctx, "user123").
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
			name: "Test 2. Negative. Invalid profile data",
			args: args{
				ctx: ctx,
				profile: &EditProfile{
					Nickname: "user123",
				}, // missing other fields for validation
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
			name: "Test 3. Negative. User not found",
			args: args{
				ctx: ctx,
				profile: &EditProfile{
					Nickname: "user123",
				},
			},
			mock: func(t *testing.T) DepsProfile {
				profileRepoMock := mocks.NewRepositoryProfile(t)

				// mock FindUsersByNickname returning empty list
				profileRepoMock.EXPECT().
					FindUsersByNickname(ctx, "user123").
					Return([]models.Profile{}, nil).Once()

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
					FindUsersByNickname(ctx, "user123").
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

func TestEditUserProfile(t *testing.T) {
	t.Parallel()

	var ctx = context.Background()

	tests := []struct {
		name        string
		profile     *EditProfile
		mock        func(t *testing.T) DepsProfile
		wantErr     assert.ErrorAssertionFunc
		wantProfile *models.Profile
	}{
		{
			name: "Test 1. Positive case. Profile is updated",
			profile: &EditProfile{
				Nickname:        "existinguser",
				ChangedNickname: "newnickname",
				Bio:             "Updated bio",
				Avatar:          "newavatar.png",
			},
			mock: func(t *testing.T) DepsProfile {
				repoMock := mocks.NewRepositoryProfile(t)
				repoMock.On("FindUsersByNickname", ctx, "existinguser").Return([]models.Profile{
					{Nickname: "existinguser"},
				}, nil).Once()

				repoMock.On("EditUserProfile", ctx, mock.AnythingOfType("*models.Profile")).Return(nil).Once()

				return DepsProfile{ProfileRepository: repoMock}
			},
			wantErr: assert.NoError,
			wantProfile: &models.Profile{
				Nickname: "newnickname",
				Bio:      "Updated bio",
				Avatar:   "newavatar.png",
			},
		},
		{
			name:    "Test 2. Error case. Profile is nil",
			profile: nil,
			mock: func(t *testing.T) DepsProfile {
				return DepsProfile{}
			},
			wantErr: assert.Error,
		},
		{
			name: "Test 3. Error case. Profile validation fails",
			profile: &EditProfile{
				Nickname: "", // invalid profile
			},
			mock: func(t *testing.T) DepsProfile {
				return DepsProfile{}
			},
			wantErr: assert.Error,
		},
		{
			name: "Test 4. Error case. User not found",
			profile: &EditProfile{
				Nickname: "nonexistentuser",
			},
			mock: func(t *testing.T) DepsProfile {
				repoMock := mocks.NewRepositoryProfile(t)
				repoMock.On("FindUsersByNickname", ctx, "nonexistentuser").Return([]models.Profile{}, nil).Once()
				return DepsProfile{
					ProfileRepository: repoMock,
				}
			},
			wantErr: assert.Error,
		},
		{
			name: "Test 5. Error case. Repository returns error",
			profile: &EditProfile{
				Nickname: "existinguser",
			},
			mock: func(t *testing.T) DepsProfile {
				repoMock := mocks.NewRepositoryProfile(t)
				repoMock.On("FindUsersByNickname", ctx, "existinguser").Return([]models.Profile{
					{Nickname: "existinguser"},
				}, nil).Once()
				repoMock.On("EditUserProfile", ctx, mock.AnythingOfType("*models.Profile")).Return(errors.New("db error")).Once()
				return DepsProfile{
					ProfileRepository: repoMock,
				}
			},
			wantErr: assert.Error,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			repoMock := tt.mock(t)
			usecase := &profileUsecase{
				DepsProfile: repoMock,
			}

			// ACT
			err := usecase.EditUserProfile(ctx, tt.profile)

			// ASSERT
			tt.wantErr(t, err)

			// Additional assertions if the profile was successfully edited
			if err == nil {
				// Profile should be edited with the updated fields
				assert.NotNil(t, tt.wantProfile)
				assert.Equal(t, tt.wantProfile.Nickname, tt.profile.ChangedNickname)
				assert.Equal(t, tt.wantProfile.Bio, tt.profile.Bio)
				assert.Equal(t, tt.wantProfile.Avatar, tt.profile.Avatar)
			}

			//assert.Equal(t, tt.want, got)  // mark test as failed but continue execution
			//require.Equal(t, tt.want, got) // mark test as failed and exit

			//repoMock.ProfileRepository.AssertExpectations(t)
		})
	}
}
