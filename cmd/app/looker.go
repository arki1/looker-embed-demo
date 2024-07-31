package main

import (
	"log"

	"github.com/looker-open-source/sdk-codegen/go/rtl"
	v4 "github.com/looker-open-source/sdk-codegen/go/sdk/v4"
)

var (
	// SessionLength defined how many seconds the user will remain authenticated.
	SessionLength = int64(300)
)

// ConnectWithLooker will setup the required Looker SDK from the environment.
func ConnectWithLooker() (*v4.LookerSDK, error) {
	cfg, err := rtl.NewSettingsFromEnv()
	if err != nil {
		return nil, err
	}
	log.Printf("Connecting with %v", cfg.BaseUrl)

	session := rtl.NewAuthSession(cfg)
	sdk := v4.NewLookerSDK(session)
	return sdk, nil
}

// SignedEmbedURL is a small wrapper that will call the corresponding Looker API
// and return a Signed Embedding URL. The user can either be redirected to this page,
// or we can use an iframe with this embedding.
func SignedEmbedURL(user string) (url string, err error) {
	sdk, err := ConnectWithLooker()
	if err != nil {
		return "", err
	}

	first_name := "External"
	last_name := "User"

	req := v4.EmbedSsoParams{
		TargetUrl:     *DashboardURL,
		SessionLength: &SessionLength,

		ExternalUserId: &user,
		FirstName:      &first_name,
		LastName:       &last_name,
		UserAttributes: &map[string]interface{}{
			"email": user,
		},

		Permissions: &[]string{"access_data", "see_looks", "see_user_dashboards"},
		Models:      &[]string{"google_cloud"},
	}

	resp, err := sdk.CreateSsoEmbedUrl(req, nil)
	if err != nil {
		return "", err
	}

	return *resp.Url, nil
}
