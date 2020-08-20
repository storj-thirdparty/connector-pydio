package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/go-openapi/strfmt"
	cells_sdk "github.com/pydio/cells-sdk-go"
	"github.com/pydio/cells-sdk-go/client"
	"github.com/pydio/cells-sdk-go/client/meta_service"
	"github.com/pydio/cells-sdk-go/models"
	"github.com/pydio/cells-sdk-go/transport"
	awstransport "github.com/pydio/cells-sdk-go/transport/aws"
)

// configPydioRest defines the Pydio REST configuration variables
var configPydioRest = &cells_sdk.SdkConfig{
	Url:        `json:"url"`,
	User:       `json:"user"`,
	Password:   `json:"password"`,
	SkipVerify: true,
}

// configPydioS3 defines the Pydio S3 configuration variables
var configPydioS3 = &cells_sdk.S3Config{
	Bucket:                 `json:"bucket"`,
	Region:                 `json:"region"`,
	ApiSecret:              `json:"apiSecret"`,
	UsePydioSpecificHeader: true,
	IsDebug:                false,
}

// ConfigPydioRest type alias for cells_sdk
type ConfigPydioRest = *cells_sdk.SdkConfig

// ConfigPydioS3 type alias for cells_sdk
type ConfigPydioS3 = *cells_sdk.S3Config

// Stack type alias for LIFO data structure implementation
type Stack []string

// IsEmpty checks if the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str string) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Pop removes and returns top element of stack. Returns false if stack is empty.
func (s *Stack) Pop() (string, bool) {
	var fileName string
	var stackNotEmpty bool

	if s.IsEmpty() {
		fileName = ""
		stackNotEmpty = false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		fileName = element
		stackNotEmpty = true
	}

	return fileName, stackNotEmpty
}

// LoadPydioProperty reads and parses the JSON file
// that contains a Pydio Cells instance's property
// and returns all the properties as REST object
func LoadPydioProperty(fullFileName string) ConfigPydioRest { // fullFileName for fetching Pydio Cells instance credentials from  given JSON filename.

	// Open and read the file.
	fileHandle, err := os.Open(filepath.Clean(fullFileName))
	if err != nil {
		log.Fatal(err)
	}

	jsonParser := json.NewDecoder(fileHandle)
	if err = jsonParser.Decode(&configPydioRest); err != nil {
		log.Fatal(err)
	}

	if err = fileHandle.Close(); err != nil {
		log.Fatal(err)
	}

	// Open and read the file.
	fileHandle2, err2 := os.Open(filepath.Clean(fullFileName))
	if err2 != nil {
		log.Fatal(err2)
	}

	jsonParser2 := json.NewDecoder(fileHandle2)
	if err2 := jsonParser2.Decode(&configPydioS3); err2 != nil {
		log.Fatal(err2)
	}

	if err = fileHandle2.Close(); err != nil {
		log.Fatal(err)
	}

	// Display the read Pydio Cells instance REST configuration properties.
	fmt.Println("\nRead Pydio REST configuration from the ", fullFileName, " file")
	fmt.Println("Pydio Cells Instance Url\t: ", configPydioRest.Url)
	fmt.Println("User\t\t\t\t: ", configPydioRest.User)
	fmt.Println("Password\t\t\t: ", configPydioRest.Password)
	fmt.Println("SkipVerify\t\t\t: ", configPydioRest.SkipVerify)

	// Display the read Pydio Cells instance S3 configuration properties.
	fmt.Println("\nRead Pydio S3 configuration from the ", fullFileName, " file")
	fmt.Println("Bucket\t\t\t: ", configPydioS3.Bucket)
	fmt.Println("Endpoint\t\t: ", configPydioRest.Url)
	fmt.Println("Region\t\t\t: ", configPydioS3.Region)
	fmt.Println("ApiKey\t\t\t: ", configPydioS3.ApiKey)
	fmt.Println("ApiSecret\t\t: ", configPydioS3.ApiSecret)
	fmt.Println("UsePydioSpecificHeader\t: ", configPydioS3.UsePydioSpecificHeader)
	fmt.Println("IsDebug\t\t\t: ", configPydioS3.IsDebug)

	return configPydioRest
}

// GetApiClient connects to the Pydio Cells server defined by this config, by sending an authentication
// request to the OIDC service to get a valid JWT (or taking the JWT from cache).
// Also returns a context to be used in subsequent requests.
func GetApiClient(sdkConfig *cells_sdk.SdkConfig, anonymous ...bool) (context.Context, *client.PydioCellsRest, error) {

	anon := false
	if len(anonymous) > 0 && anonymous[0] {
		anon = true
	}

	c, t, e := transport.GetRestClientTransport(sdkConfig, anon)
	if e != nil {
		return nil, nil, e
	}
	cl := client.New(t, strfmt.Default)
	return c, cl, nil

}

// ConnectToPydio will connect to a Pydio Cells instance,
// based on the read property from an external file.
// It returns list of workspaces in the Pydio Cells instance
// and the backup prefix path name
func ConnectToPydio(configPydioRest ConfigPydioRest) ([]string, string) {

	fmt.Println("\nConnecting to Pydio Cells.")

	//connects to the pydio api via the sdkConfig
	ctx, apiClient, err := GetApiClient(configPydioRest)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Fetching all the Pydio Cells workspaces.")

	workspaces := make([]string, 0)

	/*GetBulkMetaParams contains all the parameters to send to the API endpoint
	for the get bulk meta operation typically these are written to a http.Request
	*/
	params := &meta_service.GetBulkMetaParams{
		Body: &models.RestGetBulkMetaRequest{NodePaths: []string{
			//the workspaces from whom the files are listed
			"/*",
		}},
		Context: ctx,
	}

	//assigns the files data retrieved above in the results variable
	result, err := apiClient.MetaService.GetBulkMeta(params)
	if err != nil {
		fmt.Printf("could not list meta: %s\n", err.Error())
		log.Fatal(err)
	}

	//prints the path therefore the name of the files listed
	if len(result.Payload.Nodes) > 0 {
		fmt.Printf("* %d workspaces:\n", len(result.Payload.Nodes))
		for _, u := range result.Payload.Nodes {
			fmt.Println(u.Path)
			workspaces = append(workspaces, u.Path+"/*")
		}
	}

	t := time.Now()
	timeNow := t.Format("2006-01-02_15_04_05")
	backupPrefixPath := "pydio-backup" + "_" + timeNow + "/"

	return workspaces, backupPrefixPath
}

// GetFileNames will return all the file names in the Pydio Cells instance
func GetFileNames(configPydioRest ConfigPydioRest, workspaces []string) []string {

	fmt.Println("\nConnecting to Pydio Cells.")

	//connects to the pydio api via the sdkConfig
	ctx, apiClient, err := GetApiClient(configPydioRest)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Fetching all the file names from the Pydio Cells instance.")

	fileNames := make([]string, 0)

	totalFiles := 0

	var stack Stack // create a stack variable of type Stack

	stack.Push("")

	for {

		possiblePath, _ := stack.Pop()

		/*GetBulkMetaParams contains all the parameters to send to the API endpoint
		  for the get bulk meta operation typically these are written to a http.Request
		*/
		params := &meta_service.GetBulkMetaParams{
			Body:    &models.RestGetBulkMetaRequest{NodePaths: []string{possiblePath + "/*"}},
			Context: ctx,
		}

		//assigns the files data retrieved above in the results variable
		result, err := apiClient.MetaService.GetBulkMeta(params)
		if err != nil {
			fmt.Printf("could not list meta: %s\n", err.Error())
			log.Fatal(err)
		}

		//prints the path therefore the name of the files listed
		for _, u := range result.Payload.Nodes {

			if u.Type == models.TreeNodeTypeLEAF {
				fileNames = append(fileNames, u.Path)
				totalFiles++
			} else {
				flag := 1
				for i := range workspaces {
					if u.Path == workspaces[i][:len(workspaces[i])-2]+"/recycle_bin" {
						flag = 0
						break
					}
				}
				if flag != 0 {
					stack.Push(u.Path)
				}

			}
		}

		if stack.IsEmpty() {
			break
		}
	}
	fmt.Println("Total files: ", totalFiles)

	sort.Strings(fileNames)

	for i := range fileNames {
		fmt.Println(fileNames[i])
	}

	return fileNames
}

// GetS3Client returns the Pydio Cells S3 Client
// which is required for getting io.Reader implementor from GetFile function
func GetS3Client() (*s3.S3, string, error) {
	bucketName := configPydioS3.Bucket
	configPydioS3.Endpoint = configPydioRest.Url
	s3Client, e := awstransport.GetS3CLient(configPydioRest, configPydioS3)
	return s3Client, bucketName, e
}

// GetFile returns the io.Reader implementor for a file path in Pydio Cells
func GetFile(pathToFile string) (io.Reader, int, error) {

	s3Client, bucketName, e := GetS3Client()
	if e != nil {
		return nil, 0, e
	}
	hO, err := s3Client.HeadObject((&s3.HeadObjectInput{}).
		SetBucket(bucketName).
		SetKey(pathToFile),
	)
	if err != nil {
		return nil, 0, err
	}
	size := int(*hO.ContentLength)

	obj, err := s3Client.GetObject((&s3.GetObjectInput{}).
		SetBucket(bucketName).
		SetKey(pathToFile),
	)
	if err != nil {
		return nil, 0, err
	}
	return obj.Body, size, nil
}
