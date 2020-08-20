package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// storeCmd represents the store command
var storeCmd = &cobra.Command{
	Use:   "store",
	Short: "Command to upload data to storj V3 network.",
	Long:  `Command to connect and transfer all data from a desired Pydio instance to given Storj Bucket.`,
	Run:   pydioStore,
}

func init() {

	// Setup the store command with its flags.
	rootCmd.AddCommand(storeCmd)
	var defaultPydioFile string
	var defaultStorjFile string
	storeCmd.Flags().BoolP("accesskey", "a", false, "Connect to storj using access key(default connection method is by using API Key).")
	storeCmd.Flags().BoolP("share", "s", false, "For generating share access of the uploaded backup file.")
	storeCmd.Flags().StringVarP(&defaultPydioFile, "pydio", "p", "./config/pydio_property.json", "full filepath contaning Pydio configuration.")
	storeCmd.Flags().StringVarP(&defaultStorjFile, "storj", "u", "./config/storj_config.json", "full filepath contaning storj V3 configuration.")
}

func pydioStore(cmd *cobra.Command, args []string) {

	// Process arguments from the CLI.
	pydioConfigfilePath, _ := cmd.Flags().GetString("pydio")
	fullFileNameStorj, _ := cmd.Flags().GetString("storj")
	useAccessKey, _ := cmd.Flags().GetBool("accesskey")
	useAccessShare, _ := cmd.Flags().GetBool("share")

	// Read Pydio instance's REST configurations from an external file and create a Pydio configuration object.
	configPydioRest := LoadPydioProperty(pydioConfigfilePath)

	// Read storj network configurations from and external file and create a storj configuration object.
	storjConfig := LoadStorjConfiguration(fullFileNameStorj)

	// Connect to storj network using the specified credentials.
	access, project := ConnectToStorj(fullFileNameStorj, storjConfig, useAccessKey)

	// Establish connection with Pydio Cells and get the workspaces and the backupPrefixPath.
	workspaces, backupPrefixPath := ConnectToPydio(configPydioRest)

	// Fetch all the file names in the Pydio Cells instance
	uploadFileNames := GetFileNames(configPydioRest, workspaces)

	fmt.Printf("\nInitiating back-up.\n")
	// Fetch all backup files from Pydio instance and simultaneously store them into desired Storj bucket.
	for i := 0; i < len(uploadFileNames); i++ {

		// Get io.Reader implementor for a file path in Pydio Cells.
		pydioObjectReader, _, err := GetFile(uploadFileNames[i])
		if err != nil {
			log.Fatal(err)
		}

		UploadData(project, storjConfig, backupPrefixPath+uploadFileNames[i], pydioObjectReader)
	}
	fmt.Printf("\nBack-up complete.\n\n")

	// Create restricted shareable serialized access if share is provided as argument.
	if useAccessShare {
		ShareAccess(access, storjConfig)
	}
}
