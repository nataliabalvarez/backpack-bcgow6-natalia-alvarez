package products

import "fmt"

type Product struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	Color        string  `json:"color"`
	Price        float64 `json:"price"`
	Stock        int     `json:"stock"`
	Code         string  `json:"code"`
	Published    bool    `json:"published"`
	CreationDate string  `json:"creationDate"`
}

var products []Product
var lastId int

type Repository interface {
	Get(id int) (Product, error)
	GetAll() ([]Product, error)
	Store(id int, name, color string, price float64, stock int, code string, published bool, creationDate string) (Product, error)
	LastID() (int, error)
	Put(id int, name, color string, price float64, stock int, code string, published bool, creationDate string) (Product, error)
	Patch(id int, name, color string, price float64, stock int, code string, published bool, creationDate string) (Product, error)
	Delete(id int) error
}
type repository struct{} //struct implementa los metodos de la interfaz

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) Store(id int, name, color string, price float64, stock int, code string, published bool, creationDate string) (Product, error) {
	prod := Product{id, name, color, price, stock, code, published, creationDate}
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

func (r *repository) Put(id int, name, color string, price float64, stock int, code string, published bool, creationDate string) (Product, error) {
	aux := Product{id, name, color, price, stock, code, published, creationDate}
	for i := range products {
		if products[i].Id == id {
			products[i] = aux
			return products[i], nil
		}
	}
	return Product{}, fmt.Errorf("producto %d no encontrado", id)
}

func (r *repository) Patch(id int, name, color string, price float64, stock int, code string, published bool, creationDate string) (Product, error) {

	for i := range products {
		if products[i].Id == id {
			if name != "" {
				products[i].Name = name
			}
			if color != "" {
				products[i].Color = color
			}
			if price != 0 {
				products[i].Price = price
			}
			if stock != 0 {
				products[i].Stock = stock
			}
			if code != "" {
				products[i].Code = code
			}
			if published {
				products[i].Published = published				
			}
			if creationDate != "" {
				products[i].CreationDate = creationDate
			}

			return products[i], nil			
		}
	}

	return Product{}, fmt.Errorf("producto %d no encontrado", id)
}

func (r *repository) Delete(id int) error {
	for i := range products {
		if products[i].Id == id {
			products = append(products[:i], products[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("producto %d no encontrado", id)
}