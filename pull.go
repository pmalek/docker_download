package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// DockerAPIPullBlob is the URL for getting the layers for particular image
// @param1 is the name of the repository
// @param2 is the digest of the layer
const DockerAPIPullBlob = "https://registry.hub.docker.com/v2/%s/blobs/%s"

type imageLayer struct {
	digest string
	layer  *bytes.Buffer
	length int64
}

type imageLayers []imageLayer

func (il *imageLayers) isLayerAlreadyCached(digest string) bool {
	for _, layer := range *il {
		if layer.digest == digest {
			return true
		}
	}
	return false
}

func getBlobs(token, repo string, fsLayers FsLayers) (imageLayers, error) {
	layers := make(imageLayers, 0, len(fsLayers))

	//body, _ := ioutil.ReadAll(resp.Body)

	for _, layerDigest := range fsLayers {
		if layers.isLayerAlreadyCached(layerDigest.BlobSum) {
			continue
		}

		URL := fmt.Sprintf(DockerAPIPullBlob, repo, layerDigest.BlobSum)

		fmt.Printf("Downloading %v...\n", URL)

		req, _ := http.NewRequest(http.MethodGet, URL, nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
		req.Header.Set("Host", DockerHubURL)
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, fmt.Errorf("Fail during HTTP GET %s: %v", DockerAPIPullBlob, err)
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, fmt.Errorf("Couldn't read all response body data: %v", err)
		}

		length, err := strconv.ParseInt(res.Header.Get("Content-Length"), 10, 64)
		if err != nil {
			return nil, fmt.Errorf("Couldn't read 'Content-Length' header: %v", err)
		}

		fmt.Printf("Layer %v with size %dB\n", layerDigest.BlobSum, length)

		layers = append(layers, imageLayer{
			digest: layerDigest.BlobSum,
			layer:  bytes.NewBuffer(body),
			length: length,
		})
	}

	return layers, nil
}
