package user

type LoginRequestDomain struct {
	Authorization string `header:"authorization"`
	AccessToken   string `json:"accessToken"`
	Platform      string `json:"platform" binding:"required"`
}

type LoginResponseDomain struct {
	User LoginUserDetailDomain `json:"user"`
	Plan LoginPlanDetailDomain `json:"plan"`
}

type LoginUserDetailDomain struct {
	FirebaseId     string  `json:"firebaseId"`
	Name           string  `json:"name"`
	Email          *string `json:"email"`
	ProfilePic     *string `json:"profilePic"`
	Platform       *string `json:"platform"`
	AccessToken    *string `json:"accessToken"`
	StripeId       *string `json:"stripeId"`
	BalanceMessage int64   `json:"balanceMessage"`
}

type LoginPlanDetailDomain struct {
	PlanType    string `json:"planType"`
	MaxMessages int64  `json:"maxMessages"`
}

type UserDetailDomain struct {
	FirebaseId    string
	Name          string
	ProfilePicUrl string
	Email         string
}
