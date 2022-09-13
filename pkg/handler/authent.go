package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/T-jegou/myTravelNotebook/pkg/config"
	oidc "github.com/coreos/go-oidc" // Google OpenID client
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"golang.org/x/oauth2" // OAuth2 client
)

var (
	// state carries an internal token during the oauth2 workflow
	// we just need a non empty initial value
	state = "foobar" // Don't make this a global in production.

	// the credentials for this API (adapt values when registering API)
	clientID     = config.Config.ClientID     // <= enter registered API client ID here
	clientSecret = config.Config.ClientSecret // <= enter registered API client secret here

	// the Google login URL
	authURL = "https://accounts.google.com/o/oauth2/v2/auth"

	// the Google OAuth2 resource provider which delivers access tokens
	tokenURL    = "https://www.googleapis.com/oauth2/v4/token"
	userInfoURL = "https://www.googleapis.com/oauth2/v3/userinfo"

	// our endpoint to be called back by the redirected client
	// callbackURL = "http://127.0.0.1:55600/api/v1/auth/callback"
	callbackURL = "http://127.0.0.1:8080/homepage"

	// the description of the OAuth2 flow
	endpoint = oauth2.Endpoint{
		AuthURL:  authURL,
		TokenURL: tokenURL,
	}

	config_0auth = oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     endpoint,
		RedirectURL:  callbackURL,
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}
)

func Login(r *http.Request) middleware.Responder {
	// implements the login with a redirection
	return middleware.ResponderFunc(
		func(w http.ResponseWriter, pr runtime.Producer) {
			http.Redirect(w, r, config_0auth.AuthCodeURL(state), http.StatusFound)
		})
}

func Callback(r *http.Request) (string, error) {
	// we expect the redirected client to call us back
	// with 2 query params: state and code.
	// We use directly the Request params here, since we did not
	// bother to document these parameters in the spec.

	if r.URL.Query().Get("state") != state {
		logrus.Println("state did not match")
		return "", fmt.Errorf("state did not match")
	}

	myClient := &http.Client{}

	parentContext := context.Background()
	ctx := oidc.ClientContext(parentContext, myClient)

	authCode := r.URL.Query().Get("code")
	logrus.Printf("Authorization code: %v\n", authCode)

	// Exchange converts an authorization code into a token.
	// Under the hood, the oauth2 client POST a request to do so
	// at tokenURL, then redirects...
	oauth2Token, err := config_0auth.Exchange(ctx, authCode)
	if err != nil {
		logrus.Println("failed to exchange token", err.Error())
		return "", fmt.Errorf("failed to exchange token")
	}

	// the authorization server's returned token
	logrus.Println("Raw token data:", oauth2Token)
	fmt.Printf("oauth2Token.TokenType: %v\n", oauth2Token.TokenType)
	return oauth2Token.AccessToken, nil
}

func authenticated(token string) (bool, error) {
	// validates the token by sending a request at userInfoURL
	bearToken := "Bearer " + token
	req, err := http.NewRequest("GET", userInfoURL, nil)
	if err != nil {
		return false, fmt.Errorf("http request: %v", err)
	}

	req.Header.Add("Authorization", bearToken)

	cli := &http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		return false, fmt.Errorf("http request: %v", err)
	}
	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, fmt.Errorf("fail to get response: %v", err)
	}
	if resp.StatusCode != 200 {
		return false, nil
	}
	return true, nil
}
