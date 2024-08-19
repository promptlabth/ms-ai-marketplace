package user

import (
	"context"

	"github.com/promptlabth/proto-lib/user"
)

func (u *UserUsecase) LoginService(ctx context.Context, req LoginRequestDomain) (*LoginResponseDomain, error) {

	token, err := u.userAdaptor.ValidateToken(ctx, req.Authorization)
	if err != nil {
		return nil, err
	}

	usr, err := u.userAdaptor.UpsertUser(ctx, &user.UpsertUserReq{
		FirebaseId: token.UID,
		Name:       token.Claims["name"].(string),
		Email: func() *string {
			if val, ok := token.Claims["email"].(string); ok {
				return &val
			}
			return nil
		}(),
		ProfilePic: func() *string {
			if val, ok := token.Claims["picture"].(string); ok {
				return &val
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
