package serve

import (
	"fmt"
	"os"

	"bitbucket.org/dyfrag-internal/mass-media-core/pkg/configs"
	"bitbucket.org/dyfrag-internal/mass-media-core/pkg/database"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
)

func upload() error {
	// The session the S3 Uploader will use
	sess := session.Must(session.NewSession(&aws.Config{
			Region: aws.String("eu-west-1"),
			Credentials: credentials.NewStaticCredentials("AKIA2N72UTRD75EH65IK", "2ULV3FoOdsqGu+u8t07NKO5Vjf5NIfORlN8/dUYR", ""),
		}))

	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)

	filename := "./main.go"
	f, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file %q, %v", filename, err)
	}

	// Upload the file to S3.
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("mass-media-core"),
		Key:    aws.String("some/"),
		Body:   f,
	})
	if err != nil {
		return fmt.Errorf("failed to upload file, %v", err)
	}
	
	return fmt.Errorf("file upload")
	// fmt.Printf("file uploaded to, %s\n", aws.StringValue(result.Location))
}

func main() {
	app := fiber.New()
	initialization()

	app.Get("/", func(c *fiber.Ctx) error {
		err := upload()
		
		fmt.Println(err)
		return c.SendString("Hello, World!")
	})

	// routes.SetUpRoutes(app)

	port := os.Getenv("SERVER_PORT")
	connection := ":" + port
	app.Listen(connection)
}

func initialization() {
	configs.SetUpConfigs()
	database.SetUpDB()
}

func New() *cobra.Command {

	return &cobra.Command{
		Use:   "serve",
		Short: "runs http server",
		Run: func(cmd *cobra.Command, args []string) {
			main()
		},
	}
}

