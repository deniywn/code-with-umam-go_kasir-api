package main

import (
	"encoding/json"
	"fmt"
	"kasir-api/database"
	"kasir-api/handlers"
	"kasir-api/repositories"
	"kasir-api/services"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Port    string `mapstructure:"PORT"`
	DB_CONN string `mapstructure:"DB_CONN"`
}

// Fungsi untuk MAIN & Routing
func main() {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}
	config := Config{
		Port:    viper.GetString("PORT"),
		DB_CONN: viper.GetString("DB_CONN"),
	}

	// Setup database
	db, err := database.InitDB(config.DB_CONN)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	//depedency injection
	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	// Transaction
	transactionRepo := repositories.NewTransactionRepository(db)
	transactionService := services.NewTransactionService(transactionRepo)
	transactionHandler := handlers.NewTransactionHandler(transactionService)

	// Report
	reportRepo := repositories.NewReportRepository(db)
	reportService := services.NewReportService(reportRepo)
	reportHandler := handlers.NewReportHandler(reportService)
	// Setup routes

	//**product**
	http.HandleFunc("/api/produk", productHandler.HandleProducts)
	http.HandleFunc("/api/produk/", productHandler.HandleProductByID)
	//**transaction**
	http.HandleFunc("/api/checkout", transactionHandler.HandleCheckout) // POST
	//**report
	http.HandleFunc("/api/report/hari-ini", reportHandler.GetTodaySalesSummary)
	http.HandleFunc("/api/report", reportHandler.GetSalesSummary)

	// localhost:8080/health
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Api running",
			"status":  "ok",
		})
	})
	fmt.Println("server running di localhost:" + config.Port)

	//	err := http.ListenAndServe(":8080", nil)
	err = http.ListenAndServe(":"+config.Port, nil)
	if err != nil {
		fmt.Println("gagal encoding server")
	}
}
