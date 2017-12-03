package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func main() {
	appName := strings.TrimPrefix(os.Args[0], "./")

	/*
	 *docker_download layers
	 */

	var cmdLayers = &cobra.Command{
		Use:   "layers [image]",
		Short: "Get layers info about specified image",
		Long: `Downloads docker layers info about specified image
at specified tag and prints them on screen.`,
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			image := args[0]
			tagF := cmd.Flag("tag")
			layers(image, tagF.Value.String())
		},
	}
	cmdLayers.Example = appName + " layers mysql/mysql-server --tag 5.6.23"

	/*
	 *docker_download pull
	 */

	var cmdPull = &cobra.Command{
		Use:   "pull [image]",
		Short: "Downloads layers from specified image",
		Long: `Downloads docker layers from specified image
at specified tag and saves them in directory named 
by repo name and tag name in the current directory.`,
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			image := args[0]
			tagF := cmd.Flag("tag")
			pull(image, tagF.Value.String())
		},
	}
	cmdPull.Example = appName + " pull mysql/mysql-server --tag 5.6.23"

	/*
	 *root
	 */

	var rootCmd = &cobra.Command{Use: "docker_download"}
	rootCmd.PersistentFlags().String("tag", "", "tag of the image to get info on from docker registry")
	rootCmd.MarkPersistentFlagRequired("tag")

	rootCmd.AddCommand(cmdLayers, cmdPull)
	rootCmd.Execute()
}

func layers(repo, tag string) {
	token, err := getAuthToken(repo)
	if err != nil {
		log.Fatalf("Couldn't get auth token: %v", err)
	}

	manifest, err := getManifest(token, repo, tag)
	if err != nil {
		log.Fatalf("Couldn't get the manifest: %v", err)
	}

	fmt.Printf("SchemaVersion: %d\n", manifest.SchemaVersion)

	fmt.Printf("%s\n", "FsLayers")
	for i, l := range manifest.FsLayers {
		fmt.Printf("%2d %v\n", i, l)
	}

	fmt.Printf("%s\n", "History")
	for i, l := range manifest.History {
		//fmt.Printf("%#v\n", l.V1CompatibilityRaw)
		fmt.Printf("%2d %s %s\n", i, l.V1Compatibility.ID, l.V1Compatibility.ContainerConfig.Cmd)
	}
}

func pull(repo, tag string) {
	token, err := getAuthToken(repo)
	if err != nil {
		log.Fatalf("Couldn't get auth token: %v", err)
	}

	manifest, err := getManifest(token, repo, tag)
	if err != nil {
		log.Fatalf("Couldn't get the manifest: %v", err)
	}

	imageLayers, err := getBlobs(token, repo, manifest.FsLayers)
	if err != nil {
		log.Fatalf("Couldn't get the layers' blobs: %v", err)
	}

	dirName := fmt.Sprintf("%s_%s", strings.Replace(repo, "/", "_", -1), tag)
	err = os.Mkdir(dirName, 0777)
	if err != nil {
		log.Fatalf("Couldn't create directory for downloaded layers: %v", err)
	}

	for _, layer := range imageLayers {
		// layerDigest is in form 'sha256:$THE_DIGEST'
		if strings.Contains(layer.digest, ":") {
			layer.digest = strings.Split(layer.digest, ":")[1]
		}

		f, err := os.Create(dirName + "/" + layer.digest)
		if err != nil {
			log.Fatalf("Couldn't create file for downloaded layer %s: %v", layer.digest, err)
		}
		defer f.Close()
		io.Copy(f, layer.layer)
	}
}
