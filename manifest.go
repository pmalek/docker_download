package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

func repoNameToRegistryImageTuple(repo string) (string, string, error) {
	s := strings.Split(repo, "/")

	var registryURL, image string
	switch len(s) {
	case 2:
		registryURL = "registry.hub.docker.com"
		image = s[0] + "/" + s[1]
	case 3:
		registryURL = s[0]
		image = s[1] + "/" + s[2]
	default:
		return "", "", fmt.Errorf("couldn't get repo name correctly: %s", repo)
	}

	return registryURL, image, nil

}

func getManifest(t tokenGetter, repo, tag string) (*ManifestResponse, error) {
	registryURL, image, err := repoNameToRegistryImageTuple(repo)
	if err != nil {
		return nil, err
	}

	token, err := t.get(registryURL, image)
	if err != nil {
		log.Fatalf("Couldn't get auth token: %v", err)
	}

	if err := checkIfImageExists(token, registryURL, image, tag); err != nil {
		return nil, errors.Wrapf(err, "Image %s doesn't exist in %s", image, registryURL)
	}

	fmt.Printf("Image %s:%s exists at registry %s\n", image, tag, registryURL)

	URL := fmt.Sprintf(DockerAPIManifestf, registryURL, image, tag)

	req, _ := http.NewRequest(http.MethodGet, URL, nil)
	if len(token) > 0 {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Fail during HTTP GET %s: %v", URL, err)
	}
	defer res.Body.Close()

	//b, _ := ioutil.ReadAll(res.Body)
	//fmt.Printf("%s\n", b)

	contentType := res.Header.Get("Content-Type")

	if !strings.Contains(contentType, "application/vnd.docker.distribution.manifest.v1+json") &&
		!strings.Contains(contentType, "application/vnd.docker.distribution.manifest.v1+prettyjws") {
		return nil,
			fmt.Errorf("Content-Type (schema version) %s not supported", contentType)
	}

	// V1
	jsonManResp := &ManifestResponse{}
	if err = json.NewDecoder(res.Body).Decode(jsonManResp); err != nil {
		return nil, fmt.Errorf("Couldn't decode JSON manifest: %v", err)
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
