package initializers

import(
	"github.com/joho/godotenv"
	"log"
)

func LoadEnvInitializers() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error while loading .env")
	}
}