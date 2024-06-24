package migrations

import (
	"fmt"
	"log"

	"github.com/budisetionugroho123/money_manegement/internal/config"
	"github.com/budisetionugroho123/money_manegement/internal/models"
)

func RunMigration() {
	err := config.InitDB().AutoMigrate(&models.User{}, &models.Category{}, &models.Budget{}, &models.Transaction{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Success run migration")
}
