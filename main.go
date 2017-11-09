package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	image := flag.String("image", "", "image to get info on e.g. library/nginx or mysql/mysql-server")
	tag := flag.String("tag", "", "tag of the image to get manifest on e.g. 1.2 or latest")
	flag.Parse()

	if len(os.Args) < 3 || len(*image) == 0 || len(*tag) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	manifest, err := getManifest(*image, *tag)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%+v\n", jsonManResp)

	fmt.Printf("SchemaVersion: %d\n", manifest.SchemaVersion)

	fmt.Printf("%s\n", "FsLayers")
	for i, l := range manifest.FsLayers {
		fmt.Printf("%d %v\n", i, l)
	}

	fmt.Printf("%s\n", "History")
	for i, l := range manifest.History {
		//fmt.Printf("%#v\n", l.V1CompatibilityRaw)
		fmt.Printf("%d %s %s\n", i, l.V1Compatibility.ID, l.V1Compatibility.ContainerConfig.Cmd)
	}
}
