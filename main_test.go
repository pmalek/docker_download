package main

import (
	"encoding/json"
	"strings"
	"testing"
)

func Test_Successfull_Unmarshalling_ManifestResponse1(t *testing.T) {
	json1 := `{
		"schemaVersion": 1,
   "name": "library/redis",
   "tag": "latest",
   "architecture": "amd64",
   "history": [
      {
         "v1Compatibility": "{\"id\":\"ef8a93741134ad37c834c32836aefbd455ad4aa4d1b6a6402e4186dfc1feeb88\",\"parent\":\"9c8b347e3807201285053a5413109b4235cca7d0b35e7e6b36554995cfd59820\",\"created\":\"2017-10-10T02:53:19.011435683Z\",\"container_config\":{\"Cmd\":[\"/bin/sh -c #(nop)  ENTRYPOINT [\\\"docker-entrypoint.sh\\\"]\"]},\"throwaway\":true}"
      }
   ]
}`

	var jsonManResp ManifestResponse
	str := strings.NewReader(json1)
	if err := json.NewDecoder(str).Decode(&jsonManResp); err != nil {
		t.Error(err)
	}

	if jsonManResp.Name != "library/redis" {
		t.Errorf("Expected name: %s but got name: %s", "library/redis", jsonManResp.Name)
	}
	if jsonManResp.Tag != "latest" {
		t.Errorf("Expected name: %s but got name: %s", "latest", jsonManResp.Tag)
	}
	if jsonManResp.Architecture != "amd64" {
		t.Errorf("Expected name: %s but got name: %s", "amd64", jsonManResp.Architecture)
	}
	if len(jsonManResp.History) != 1 {
		t.Errorf("Expected history to be of length %d but got length  %d",
			1, len(jsonManResp.History))
	}
}

func Test_Successfull_Unmarshalling_ManifestResponse2(t *testing.T) {
	json1 := `{
   "schemaVersion": 1,
   "name": "library/redis",
   "tag": "latest",
   "architecture": "amd64",
   "fsLayers": [
      {
         "blobSum": "sha256:a3ed95caeb02ffe68cdd9fd84406680ae93d633cb16422d00e8a7c22955b46d4"
      },
      {
         "blobSum": "sha256:a3ed95caeb02ffe68cdd9fd84406680ae93d633cb16422d00e8a7c22955b46d4"
      }
   ],
   "history": [
      {
         "v1Compatibility": "{\"id\":\"ef8a93741134ad37c834c32836aefbd455ad4aa4d1b6a6402e4186dfc1feeb88\",\"parent\":\"9c8b347e3807201285053a5413109b4235cca7d0b35e7e6b36554995cfd59820\",\"created\":\"2017-10-10T02:53:19.011435683Z\",\"container_config\":{\"Cmd\":[\"/bin/sh -c #(nop)  ENTRYPOINT [\\\"docker-entrypoint.sh\\\"]\"]},\"throwaway\":true}"
      },
      {
         "v1Compatibility": "{\"id\":\"9c8b347e3807201285053a5413109b4235cca7d0b35e7e6b36554995cfd59820\",\"parent\":\"37ede13d1e9570252c0e0803e63d15b2bc6d6deac842b9ae8044e64bdd3b8d82\",\"created\":\"2017-10-10T02:53:18.84358639Z\",\"container_config\":{\"Cmd\":[\"/bin/sh -c #(nop) COPY file:9c29fbe8374a97f9c2d953c9c8b7224554607eeb7a610a930844f2bec678265c in /usr/local/bin/ \"]}}"
      }
   ]
}`

	var jsonManResp ManifestResponse
	str := strings.NewReader(json1)
	if err := json.NewDecoder(str).Decode(&jsonManResp); err != nil {
		t.Error(err)
	}

	if jsonManResp.Name != "library/redis" {
		t.Errorf("Expected name: %s but got name: %s", "library/redis", jsonManResp.Name)
	}
	if jsonManResp.Tag != "latest" {
		t.Errorf("Expected name: %s but got name: %s", "latest", jsonManResp.Tag)
	}
	if jsonManResp.Architecture != "amd64" {
		t.Errorf("Expected name: %s but got name: %s", "amd64", jsonManResp.Architecture)
	}
	if len(jsonManResp.FsLayers) != 2 {
		t.Errorf("Expected FsLayers to be of length %d but got length  %d",
			2, len(jsonManResp.FsLayers))
	}
	if len(jsonManResp.History) != 2 {
		t.Errorf("Expected history to be of length %d but got length  %d",
			2, len(jsonManResp.History))
	}
}
