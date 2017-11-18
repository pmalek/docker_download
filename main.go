package main

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

func main() {
	/*
	 *docker_download layers
	 */

	var cmdLayers = &cobra.Command{
		Use:   "layers",
		Short: "Get layers info about specified image",
		Long: `Downloads docker layers from specified image
at specified tag and prints them on screen.`,
		Run: func(cmd *cobra.Command, args []string) {
			imageF := cmd.Flag("image")
			tagF := cmd.Flag("tag")
			layers(imageF.Value.String(), tagF.Value.String())
		},
	}

	cmdLayers.Example = "docker_download layers --image mysql/mysql-server --tag 5.6.23"

	/*
	 *root
	 */

	var rootCmd = &cobra.Command{Use: "docker_download"}
	rootCmd.PersistentFlags().String("image", "", "image to get info/image on from docker registry")
	rootCmd.MarkPersistentFlagRequired("image")
	rootCmd.PersistentFlags().String("tag", "", "tag of the image to get info/imageimage on from docker registry")
	rootCmd.MarkPersistentFlagRequired("tag")

	rootCmd.AddCommand(cmdLayers)
	rootCmd.Execute()
}

func layers(image, tag string) {
	manifest, err := getManifest(image, tag)
	if err != nil {
		log.Fatal(err)
	}

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
