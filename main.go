package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/lyquocnam/go-note-learning/handler"
	"github.com/lyquocnam/go-note-learning/model"
	"github.com/lyquocnam/go-note-learning/repo"
	"github.com/lyquocnam/go-note-learning/storage"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.LogMode(true)
	db.AutoMigrate(model.Note{})

	gin.SetMode(os.Getenv("GIN_MODE"))
	engine := gin.Default()

	noteStorage := storage.NewNotePostgresStorage(db)
	noteRepo := repo.NewNoteRepo(noteStorage)
	handler.NewNoteHandler(engine, noteRepo)

	log.Fatal(engine.Run(":8080"))
}
