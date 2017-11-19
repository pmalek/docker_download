package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// DockerAPIGetAuthToken is the URL for getting auth token for particular image
const DockerAPIGetAuthToken = "https://auth.docker.io/token?service=registry.docker.io&scope=repository:%s:pull"

// DockerAPIManifest is the URL for getting the manifest for particular image
const DockerAPIManifest = "https://registry.hub.docker.com/v2/%s/manifests/%s"

//DockerHubURL  is the URL of docker hub's registry
const DockerHubURL = "registry.hub.docker.com"

// Token is a docker auth token
type Token struct {
	Token     string `json:"token"`
	ExpiresIn int    `json:"expires_in"`
	IssuedAt  string `json:"issued_at"`
}

func getAuthToken(repository string) (string, error) {
	response, err := http.Get(fmt.Sprintf(DockerAPIGetAuthToken, repository))
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	var jsonTokenResp Token
	if err = json.NewDecoder(response.Body).Decode(&jsonTokenResp); err != nil {
		return "", err
	}

	return jsonTokenResp.Token, nil
}

func checkIfImageExists(token, repo, tag string) error {
	URL := fmt.Sprintf(DockerAPIManifest, repo, tag)

	req, _ := http.NewRequest(http.MethodHead, URL, nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("Host", DockerHubURL)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("Fail during HTTP HEAD %s: %v", DockerHubURL, err)
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("Couldn't get the image %s/%s: HTTP status code %v", repo, tag, res.StatusCode)
	}
	return nil
}

func getManifest(token, repo, tag string) (ManifestResponse, error) {
	if err := checkIfImageExists(token, repo, tag); err != nil {
		return ManifestResponse{}, err
	}

	URL := fmt.Sprintf(DockerAPIManifest, repo, tag)

	req, _ := http.NewRequest(http.MethodGet, URL, nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("Host", DockerHubURL)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return ManifestResponse{}, fmt.Errorf("Fail during HTTP GET %s: %v", DockerHubURL, err)
	}
	defer res.Body.Close()

	//b, _ := ioutil.ReadAll(manifest)
	//fmt.Printf("%s\n", b)

	contentType := res.Header.Get("Content-Type")

	if !strings.Contains(contentType, "application/vnd.docker.distribution.manifest.v1+json") &&
		!strings.Contains(contentType, "application/vnd.docker.distribution.manifest.v1+prettyjws") {
		return ManifestResponse{},
			fmt.Errorf("Content-Type (schema version) %s not supported", contentType)
	}

	// V1
	var jsonManResp ManifestResponse
	if err = json.NewDecoder(res.Body).Decode(&jsonManResp); err != nil {
		return ManifestResponse{}, fmt.Errorf("Couldn't decode JSON manifest: %v", err)
	}
	return jsonManResp, nil
}

type FsLayers []struct {
	BlobSum string `json:"blobSum"`
}

// ManifestResponse is the json response from manifests API v2
type ManifestResponse struct {
	Name          string `json:"name"`
	Tag           string `json:"tag"`
	Architecture  string `json:"architecture"`
	SchemaVersion int    `json:"schemaVersion"`

	FsLayers `json:"fsLayers"`

	History []struct {
		V1CompatibilityRaw string `json:"v1Compatibility"`
		V1Compatibility    v1Compatibility
	} `json:"history"`
}

type v1Compatibility struct {
	Architecture  string `json:"architecture,omitempty"`
	OS            string `json:"os,omitempty"`
	ID            string `json:"id,omitempty"`
	Parent        string `json:"parent,omitempty"`
	Created       string `json:"created,omitempty"`
	DockerVersion string `json:"docker_version,omitempty"`
	Throwaway     bool   `json:"throwaway,omitempty"`

	//Tty       bool `json:"Tty"`
	//OpenStdin bool `json:"OpenStdin"`
	//StdinOnce bool `json:"StdinOnce"`

	//Env []string `json:"Env"`
	ContainerConfig struct {
		CmdRaw []string `json:"Cmd"`
		Cmd    string
	} `json:"container_config"`
}

type manifestResponse ManifestResponse

// UnmarshalJSON implement json.Unmarshaller interface
func (m *ManifestResponse) UnmarshalJSON(b []byte) error {
	var jsonManResp manifestResponse
	if err := json.Unmarshal(b, &jsonManResp); err != nil {
		return err
	}

	for i := range jsonManResp.History {
		var comp v1Compatibility
		if err := json.Unmarshal([]byte(jsonManResp.History[i].V1CompatibilityRaw), &comp); err != nil {
			continue
		}

		jsonManResp.History[i].V1Compatibility = comp

		jsonManResp.History[i].V1Compatibility.ContainerConfig.Cmd =
			strings.Join(jsonManResp.History[i].V1Compatibility.ContainerConfig.CmdRaw, " ")
	}

	*m = ManifestResponse(jsonManResp)
	return nil
}
