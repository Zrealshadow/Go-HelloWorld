package bapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/go-programming-tour-book/tag-service/pkg/errcode"
	"golang.org/x/net/context/ctxhttp"
)

const (
	APP_KEY    = "lingze"
	APP_SECRET = "yuke"
)

type AccessToken struct {
	Token string `json:"token"`
}

type API struct {
	URL string
}

func NewAPI(url string) *API {
	return &API{URL: url}
}

func (a *API) getAccessToken(ctx context.Context) (string, error) {
	p := url.Values{}
	p.Add("app_key", APP_KEY)
	p.Add("app_secret", APP_SECRET)
	// body, err := a.httpGet(ctx, fmt.Sprintf("%s?app_key=%s&app_secret=%s", "auth", APP_KEY, APP_SECRET))
	body, err := a.httpPost(ctx, "auth", p)
	if err != nil {
		return "", errcode.TogRPCError(errcode.ErrorGetAccessTokenFail)
	}
	var accessToken AccessToken
	_ = json.Unmarshal(body, &accessToken)
	return accessToken.Token, nil
}

func (a *API) httpGet(ctx context.Context, path string) ([]byte, error) {
	// resp, err := http.Get(fmt.Sprintf("%s/%s", a.URL, path))
	resp, err := ctxhttp.Get(ctx, http.DefaultClient, fmt.Sprintf("%s/%s", a.URL, path))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body, nil
}

func (a *API) httpPost(ctx context.Context, path string, param url.Values) ([]byte, error) {
	resp, err := http.PostForm(fmt.Sprintf("%s/%s", a.URL, path), param)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body, nil
}

func (a *API) GetTagList(ctx context.Context, name string) ([]byte, error) {

	token, err := a.getAccessToken(ctx)
	if err != nil {
		return nil, err
	}

	body, err := a.httpGet(ctx, fmt.Sprintf("%s?token=%s&name=%s", "api/v1/tags", token, name))

	if err != nil {
		return nil, err
	}

	return body, nil
}
