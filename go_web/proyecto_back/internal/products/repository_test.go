package products

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
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

	// Act
	result, err := repository.GetAll()

	// Asert
	assert.Nil(t, err)
	assert.Equal(t, dataBase, result)
}

func TestStore(t *testing.T) {

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
	result, err := repository.Store(productToCreate.Id, productToCreate.Name, productToCreate.Color, productToCreate.Price, productToCreate.Stock, productToCreate.Code, productToCreate.Published, productToCreate.CreationDate)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, mockStore.Data, expected)
	assert.Equal(t, productToCreate, result)
}
