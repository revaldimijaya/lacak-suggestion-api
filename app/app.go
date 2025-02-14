package app

import (
	"context"
	httphandler "github/revaldimijaya/lacak-api/app/http_handler"
	"github/revaldimijaya/lacak-api/app/repository"
	"github/revaldimijaya/lacak-api/app/usecase"
	"log"
	"os"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	handler httphandler.HttpHandler
	uc      usecase.Usecase
	repo    repository.Repository
)

func App() {
	envPath := path.Join(".env")
	if _, err := os.Stat(envPath); err == nil {
		_ = godotenv.Load(envPath)
	}

	httpAddr := os.Getenv("HTTP_ADDR")
	dataPath := os.Getenv("DATA_PATH")
	if dataPath == "" {
		log.Fatal("data path is empty")
	}

	repo = repository.InitRepository(repository.TrieNode{})
	err := repo.LoadDataCities(context.Background(), repository.LoadDataRequest{
		DataPath: dataPath,
	})
	if err != nil {
		log.Fatal(err)
	}

	uc = usecase.InitUsecase(repo)
	handler = httphandler.InitHTTPHandler(uc)

	r := gin.Default()

	r.GET("/suggestions", handler.GetCitySuggestions)
	r.Run(httpAddr)
}
