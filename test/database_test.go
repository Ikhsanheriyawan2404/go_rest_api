package test

import (
	// "context"
	"database/sql"
	"encoding/json"
	"fmt"
	"golang_rest_api/controller"
	"golang_rest_api/helper"
	// "golang_rest_api/model/domain"
	"golang_rest_api/repository"
	"golang_rest_api/service"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func setupTestDB() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/belajar_golang_restful_api_test")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func setupRouter(db *sql.DB) http.Handler {
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := httprouter.New()
	router.GET("/api/dsffdscategories/fdsaf", categoryController.FindAll)

	return router
}

func truncateCategory(db *sql.DB) {
	db.Exec("TRUNCATE category")
}

func TestListCategoriesSuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	// tx, _ := db.Begin()
	// categoryRepository := repository.NewCategoryRepository()
	// category1 := categoryRepository.Save(context.Background(), tx, domain.Category{
	// 	Name: "Gadget",
	// })
	// category2 := categoryRepository.Save(context.Background(), tx, domain.Category{
	// 	Name: "Computer",
	// })
	// tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories", nil)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])

	fmt.Println(responseBody)

	// var categories = responseBody["data"].([]interface{})

	// categoryResponse1 := categories[0].(map[string]interface{})
	// categoryResponse2 := categories[1].(map[string]interface{})

	// assert.Equal(t, category1.Id, int(categoryResponse1["id"].(float64)))
	// assert.Equal(t, category1.Name, categoryResponse1["name"])

	// assert.Equal(t, category2.Id, int(categoryResponse2["id"].(float64)))
	// assert.Equal(t, category2.Name, categoryResponse2["name"])
}
