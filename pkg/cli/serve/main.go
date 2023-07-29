package serve

import (
	"os"

	"bitbucket.org/dyfrag-internal/mass-media-core/pkg/configs"
	"bitbucket.org/dyfrag-internal/mass-media-core/pkg/database"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
)

func main() {
	app := fiber.New()
	initialization()

	app.Get("/", func(c *fiber.Ctx) error {
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

