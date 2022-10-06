package products

import "fmt"

type Product struct {
	Id           int     `json:"id"`
	Name         string  `json:"name" binding:"required"`
	Color        string  `json:"color" binding:"required"`
	Price        float64 `json:"price" binding:"required"`
	Stock        int     `json:"stock" binding:"required"`
	Code         string  `json:"code" binding:"required"`
	Published    bool    `json:"published" binding:"required"`
	CreationDate string  `json:"creationDate" binding:"required"`
}

var products []Product
var lastId int

type Repository interface {
	Get(id int) (Product, error)
	GetAll() ([]Product, error)
	Store(id int, name , color string, price float64, stock int, code string, published bool, creationDate string) (Product, error)
	LastID() (int, error)
	//Update(id int, name , color string, price float64, stock int, code string, published bool, creationDate string) (Product, error)
}
type repository struct{} //struct implementa los metodos de la interfaz

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) Store(id int, name , color string, price float64, stock int, code string, published bool, creationDate string) (Product, error) {
	prod := Product {id, name, color, price, stock, code, published, creationDate}
	products = append(products, prod)
	lastId = prod.Id
	return prod, nil
}

func (r *repository) Get(id int) (Product, error) {
	for i := range products {
		if products[i].Id == id {
			return products[i], nil
		}
	}
	return Product{}, fmt.Errorf("producto %d no encontrado", id)
}

func (r *repository) GetAll() ([]Product, error) {
	return products, nil
}

func (r *repository) LastID() (int, error) {
	return lastId, nil
}