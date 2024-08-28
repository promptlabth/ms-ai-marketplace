package user

import (
	"context"

	"github.com/promptlabth/proto-lib/user"
)

func (u *UserUsecase) LoginService(ctx context.Context, req LoginRequestDomain) (*LoginResponseDomain, error) {

	// get token from firebase
	token, err := u.userAdaptor.ValidateToken(ctx, req.Authorization)
	if err != nil {
		return nil, err
	}

	// get user detail from firebase
	userDetail, err := u.userAdaptor.FirebaseRetrieveUserData(ctx, token.UID)
	if err != nil {
		return nil, err
	}

	usr, err := u.grpcUserServer.UpsertUser(ctx, &user.UpsertUserReq{
		FirebaseId: token.UID,
		Name:       userDetail.Name,
		Email: func() *string {
			if userDetail.Email != "" {
				return &userDetail.Email
			}
			return nil
		}(),
		ProfilePic: func() *string {
			if userDetail.ProfilePicUrl != "" {
				return &userDetail.ProfilePicUrl
			}
			return nil
		}(),
		Platform:    &req.Platform,
		AccessToken: &req.AccessToken,
	})
	if err != nil {
		return nil, err
	}
	return &LoginResponseDomain{
		User: LoginUserDetailDomain{
			FirebaseId:     usr.UserDetail.FirebaseId,
			Name:           usr.UserDetail.Name,
			Email:          usr.UserDetail.Email,
			ProfilePic:     usr.UserDetail.ProfilePic,
			Platform:       usr.UserDetail.Platform,
			AccessToken:    usr.UserDetail.AccessToken,
			StripeId:       usr.UserDetail.StripeId,
			BalanceMessage: usr.UserDetail.BalanceMessage,
		},
		Plan: LoginPlanDetailDomain{
			PlanType:    usr.PlanDetail.PlanType,
			MaxMessages: usr.PlanDetail.MaxMessages,
		},
	}, nil
}
