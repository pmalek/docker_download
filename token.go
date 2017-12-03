package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/pkg/errors"
)

type tokenGetter interface {
	resetCache()
	get(registryURL, image string) (string, error)
}

type onlineTokenGetter struct {
	tokenCached bool
	token       string
}

func (o *onlineTokenGetter) resetCache() {
	o.token = ""
	o.tokenCached = false
}

func (o *onlineTokenGetter) get(registryURL, image string) (string, error) {
	if o.tokenCached {
		return o.token, nil
	}

	URL := fmt.Sprintf(DockerAPIManifestf, registryURL, image, "latest")

	res, err := http.Get(URL)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode >= http.StatusOK && res.StatusCode < http.StatusMultipleChoices {
		// token not needed for this registry
		return "", nil
	}

	if res.StatusCode != http.StatusUnauthorized {
		return "",
			errors.Wrapf(err,
				"Coudn't get the token, statusCode: %d", res.StatusCode)
	}

	if len(res.Header.Get("Www-Authenticate")) < 1 {
		return "",
			errors.Wrapf(err,
				"Coudn't obtain Www-Authenticate challenge header, statusCode: %d",
				res.StatusCode)
	}

	// Www-Authenticate Bearer realm="https://auth.docker.io/token",service="registry.docker.io",scope="repository:mysql/mysql-server:pull"
	wwwAuth := res.Header.Get("Www-Authenticate")
	s := strings.Split(strings.TrimPrefix(wwwAuth, "Bearer "), ",")
	if len(s) != 3 {
		return "", errors.Wrapf(err, "Wrong Www-Authenticate header: %s", wwwAuth)
	}

	realm := strings.TrimSuffix(strings.TrimPrefix(s[0], "realm=\""), "\"")
	service := strings.TrimSuffix(strings.TrimPrefix(s[1], "service=\""), "\"")
	scope := strings.TrimSuffix(strings.TrimPrefix(s[2], "scope=\""), "\"")

	req, err := http.NewRequest(http.MethodGet, realm, nil)
	if err != nil {
		return "", errors.Wrap(err, "Couldn't create new HTTP request")
	}
	v := url.Values{}
	v.Set("service", service)
	v.Set("scope", scope)
	req.URL.RawQuery = v.Encode()

	tokenResp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer tokenResp.Body.Close()

	var jsonTokenResp Token
	if err = json.NewDecoder(tokenResp.Body).Decode(&jsonTokenResp); err != nil {
		return "", err
	}

	o.tokenCached = true
	o.token = jsonTokenResp.Token

	return jsonTokenResp.Token, nil
}
