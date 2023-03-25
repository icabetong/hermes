package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Medicine struct {
	Id string 						`json:"id" gorm:"primaryKey"`
	Description string 		`json:"description"`
	Quantity int 					`json:"quantity"`
	Unit string 					`json:"unit"`
	Batch string 					`json:"batch"`
	Expiry time.Time 			`json:"expiry"`
	Price float32 				`json:"price"`
	CreatedAt time.Time		`json:"createdAt"`	
	UpdatedAt time.Time 	`json:"updatedAt"`
}

func main() {
	err := godotenv.Load("local.env")
	if err != nil {
		fmt.Printf("An Error Occurred while loading the environment: %s", err)
	}

	username := os.Getenv("DATABASE_USERNAME")
	password := os.Getenv("DATABASE_PASSWORD")
	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")
	name := os.Getenv("DATABASE_NAME")

	dataSource := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + name +"?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dataSource), &gorm.Config{})

	if err != nil {
		fmt.Println("An error occured connecting to the database")
		return
	}

	database.AutoMigrate(&Medicine{})

	fmt.Println("Starting the server at port 8080")

	http.HandleFunc("/medicines", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
			case "GET": handleGetRequest(database, w, r)
			case "POST": handlePostRequest(database, w, r)
		default: w.WriteHeader(http.StatusNotFound)
		}
	})

	http.HandleFunc("/medicines/", func (w http.ResponseWriter, r *http.Request)  {
		id := r.URL.Path[11:]
		switch r.Method {
		case "GET": handleDynamicGetRequest(id, database, w, r)
		case "PATCH": handleDynamicPatchRequest(id, database, w, r)
		case "DELETE": handleDynamicDeleteRequest(id, database, w, r)
		default: 
		w.WriteHeader(http.StatusNotFound)
		}
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
}

func handleGetRequest(database *gorm.DB, w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	
	var medicines []Medicine
	database.Find(&medicines)

	data, err := json.Marshal(medicines)
	if (err != nil) {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func handlePostRequest(database *gorm.DB, w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var medicine Medicine
	
	err := json.NewDecoder(r.Body).Decode(&medicine)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	database.Create(&medicine)
	
	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(http.StatusOK)
}

func handleDynamicGetRequest(id string, database *gorm.DB, w http.ResponseWriter,  r *http.Request) {
	enableCors(&w)
	var medicine Medicine

	if id == "" {
		http.Error(w, "ID cannot be blank", http.StatusBadRequest)
		return
	}

	database.First(&medicine, id)
	data, err := json.Marshal(medicine)
	if (err != nil) {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
func handleDynamicPatchRequest(id string, database *gorm.DB, w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var medicine Medicine

	err := json.NewDecoder(r.Body).Decode(&medicine)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	database.Model(Medicine{}).Where("id = ?", id).Updates(&medicine)
	w.WriteHeader(http.StatusOK)
}
func handleDynamicDeleteRequest(id string, database *gorm.DB, w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var medicine Medicine

	err := json.NewDecoder(r.Body).Decode(&medicine)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	database.Delete(&medicine)
	w.WriteHeader(http.StatusOK)
}