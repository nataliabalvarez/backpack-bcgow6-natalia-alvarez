package products

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServiceIntegrationGetAll(t *testing.T) {
	// Arrange
	dataBase := []Product{
		{
			Id:           1,
			Name:         "",
			Color:        "",
			Price:        0,
			Stock:        0,
			Code:         "",
			Published:    false,
			CreationDate: "",
		},
		{
			Id:           2,
			Name:         "",
			Color:        "",
			Price:        0,
			Stock:        0,
			Code:         "",
			Published:    false,
			CreationDate: "",
		},
	}

	mockStore := StubStore{
		Data: dataBase,
	}
	repository := NewRepository(&mockStore)
	service := NewService(repository)

	// Act
	result, err := service.GetAll()
	assert.Nil(t, err)
	assert.Equal(t, dataBase, result)

}

func TestServiceIntegrationGetAllFail(t *testing.T) {
	// Arrange
	expectedError := errors.New("error en lectura")

	mockStore := StubStore{
		Data:       nil,
		errOnWrite: nil,
		errOnRead:  errors.New("error en lectura"),
	}
	repository := NewRepository(&mockStore)
	service := NewService(repository)

	// Act
	result, err := service.GetAll()

	// Assert
	assert.EqualError(t, err, expectedError.Error())
	assert.Nil(t, result)
}

func TestServiceIntegrationWrite(t *testing.T) {

	// Arrange
	expected := []Product{
		{
			Id:           1,
			Name:         "",
			Color:        "",
			Price:        0,
			Stock:        0,
			Code:         "",
			Published:    false,
			CreationDate: "",
		},
		{
			Id:           2,
			Name:         "",
			Color:        "",
			Price:        0,
			Stock:        0,
			Code:         "",
			Published:    false,
			CreationDate: "",
		},
	}

	initialDb := []Product{
		{
			Id:           1,
			Name:         "",
			Color:        "",
			Price:        0,
			Stock:        0,
			Code:         "",
			Published:    false,
			CreationDate: "",
		},
	}

	mockStore := StubStore{
		Data: initialDb,
	}

	repository := NewRepository(&mockStore)
	service := NewService(repository)

	// Act
	productToCreate := Product{
		Id:           2,
		Name:         "",
		Color:        "",
		Price:        0,
		Stock:        0,
		Code:         "",
		Published:    false,
		CreationDate: "",
	}
	product, err := service.Store(productToCreate.Name, productToCreate.Color, productToCreate.Price, productToCreate.Stock, productToCreate.Code, productToCreate.Published, productToCreate.CreationDate)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, productToCreate, product)
	assert.Equal(t, expected, mockStore.Data)
}

func TestServiceIntegrationStoreFail(t *testing.T) {
	// Arrange
	expectedError := errors.New("error en escritura")

	mockStore := StubStore{
		Data:       nil,
		errOnWrite: errors.New("error en escritura"),
		errOnRead:  nil,
	}
	repository := NewRepository(&mockStore)
	service := NewService(repository)

	// Act
	productToCreate := Product{
		Id:           2,
		Name:         "",
		Color:        "",
		Price:        0,
		Stock:        0,
		Code:         "",
		Published:    false,
		CreationDate: "",
	}
	product, err := service.Store(productToCreate.Name, productToCreate.Color, productToCreate.Price, productToCreate.Stock, productToCreate.Code, productToCreate.Published, productToCreate.CreationDate)

	// Assert
	assert.EqualError(t, err, expectedError.Error())
	assert.Empty(t, product)
}
