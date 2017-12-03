package main

import (
	"fmt"
	"log"
	"net/http"
)

// DockerAPIManifestf is the URL for getting manifests for particular image
// @param1 is the URL of the registry
// @param2 is the name of the repository
// @param3 is the digest of the layer
const DockerAPIManifestf = "https://%s/v2/%s/manifests/%s"

// Token is a docker auth token
type Token struct {
	Token     string `json:"token"`
	ExpiresIn int    `json:"expires_in"`
	IssuedAt  string `json:"issued_at"`
}

func checkIfImageExists(token, registryURL, repo, tag string) error {
	URL := fmt.Sprintf(DockerAPIManifestf, registryURL, repo, tag)

	req, _ := http.NewRequest(http.MethodHead, URL, nil)
	if len(token) > 0 {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("Fail during HTTP HEAD %s: %v", URL, err)
	}
	if res.StatusCode != http.StatusOK {
		log.Printf("response : %+v", res)
		return fmt.Errorf("Couldn't get the image %s/%s: HTTP status code %v", repo, tag, res.StatusCode)
	}
	return nil
}
