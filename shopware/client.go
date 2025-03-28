package shopware

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

const (
    TokenEndpoint = "/oauth/token"
)

type RepositoryInterface interface {
	SetClient(client *Client)
}

type SwAPIRequestBody interface {}
type SwAPIResponse interface {}

type SwAPIRequest struct {
	Method string
	URL string
	Body SwAPIRequestBody
}

type Client struct {
	httpClient *http.Client
	baseURL string
}

func NewSwClientFromIntegration(ctx context.Context, clientID string, clientSecret string, baseURL string) *Client {
	tokenURL := baseURL + TokenEndpoint
	cfg := clientcredentials.Config{
		ClientID:    clientID,
		ClientSecret: clientSecret,
		TokenURL:     tokenURL,
	}

	oAusClient := cfg.Client(ctx)
	
	return &Client{
		httpClient: oAusClient,
		baseURL: baseURL,
	}
}

func NewSwClientFromPassword(ctx context.Context, user string, passwd string, baseURL string) *Client {
    tokenURL := baseURL + TokenEndpoint
    cfg := &oauth2.Config{
        ClientID: "administration",
        Scopes:  []string{"write"},
        Endpoint: oauth2.Endpoint{
            TokenURL: tokenURL,
        },
    }
    token, err := cfg.PasswordCredentialsToken(ctx, user, passwd)
    if (err != nil) {
        panic(err)
    }
    oAusClient := oauth2.NewClient(ctx, oauth2.StaticTokenSource(token));
    
	return &Client{
		httpClient: oAusClient,
		baseURL: baseURL,
	}
}

func NewRepository[T any](client *Client) *T {
	repo := new(T)
	if repoInterface, ok := any(repo).(RepositoryInterface); ok {
        repoInterface.SetClient(client)
    }
	return repo
}

func (c *Client)InvokeRequest(request *SwAPIRequest, response SwAPIResponse) (error) {
	url := c.baseURL + request.URL
	req, err := http.NewRequest(request.Method, url, nil)
	if err != nil {
		return err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to get config, status code: %d", resp.StatusCode)
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(bodyBytes, response); err != nil {
		return err
	}
	return nil
}


// type PasswordCredentials struct {
// 	Username string
// 	Password string
// 	Scopes   []string
// }

// func NewPasswordCredentials(username, password string, scopes []string) PasswordCredentials {
// 	return PasswordCredentials{
// 		Username: username,
// 		Password: password,
// 		Scopes:   scopes,
// 	}
// }

// func (c PasswordCredentials) GetTokenSource(ctx context.Context, tokenURL string) (oauth2.TokenSource, error) {
// 	oauthConf := &oauth2.Config{
// 		ClientID: "administration",
// 		Scopes:   c.Scopes,
// 		Endpoint: oauth2.Endpoint{
// 			TokenURL: tokenURL,
// 		},
// 	}

// 	token, err := oauthConf.PasswordCredentialsToken(ctx, c.Username, c.Password)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return oauth2.StaticTokenSource(token), nil
// }