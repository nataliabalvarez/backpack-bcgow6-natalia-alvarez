package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/nataliabalvarez/backpack-bcgow6-natalia-alvarez/go_web/proyecto_back/internal/products"
	"github.com/stretchr/testify/assert"
)



func createServer(mockStore products.StubStore) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	repo := products.NewRepository(&mockStore)
	service := products.NewService(repo)
	handler := NewProduct(service)

	r := gin.Default()

	pr := r.Group("/products")
	pr.POST("/", handler.Store())
	pr.GET("/", handler.GetAll())

	return r
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {

	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	// req.Header.Add("token", "123456")

	return req, httptest.NewRecorder()
}

func TestFunctionalGetAll(t *testing.T) {

	// Arrange
	mockDB := []products.Product {
		{
		   Id: 1,
		   Name: "",
		   Color: "",
		   Price: 0,
		   Stock: 0,
		   Code: "",
		   Published: false,
		   CreationDate: "",
		},
		{
		   Id: 2,
		   Name: "",
		   Color: "",
		   Price: 0,
		   Stock: 0,
		   Code: "",
		   Published: false,
		   CreationDate: "",
		},
	}
	
	mockStore := products.StubStore  {
		Data: mockDB,
	}

	var resp []Product
	r := createServer(mockStore)
	req, rr := createRequestTest(http.MethodGet, "/products/", "")
	
	// Act
	r.ServeHTTP(rr, req)

	// Assert
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, len(mockStore.Data), len(resp)) // fails
}

func TestFunctionalSave(t *testing.T) {
	// Arrange
	mockStorage := products.StubStore{
		Data: []products.Product{
		 {
			Id: 1,
			Name: "",
			Color: "",
			Price: 0,
			Stock: 0,
			Code: "",
			Published: false,
			CreationDate: "",
		 },
		 {
			Id: 2,
			Name: "",
			Color: "",
			Price: 0,
			Stock: 0,
			Code: "",
			Published: false,
			CreationDate: "",
		 },
		},
	}
	r := createServer(mockStorage)
	req, rr := createRequestTest(http.MethodPost, "/products/", `{name: "pipa",
	color: "",
	price: 0,
	stock: 0,
	code: "",
	published: false,
	creationDate: ""}`)
	
	// Act
	r.ServeHTTP(rr, req)
	
	// Assert
	var resp Product
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, rr.Code) // fails
}