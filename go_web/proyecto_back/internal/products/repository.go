package products

import (
	"fmt"

	"github.com/nataliabalvarez/backpack-bcgow6-natalia-alvarez/go_web/proyecto_back/pkg/store"
)


// var products []Product
var lastId int

type repository struct {
	db store.Store
} //struct implementa los metodos de la interfaz

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

type Repository interface {
	Get(id int) (Product, error)
	GetAll() ([]Product, error)
	Store(id int, name, color string, price float64, stock int, code string, published bool, creationDate string) (Product, error)
	LastID() (int, error)
	Put(id int, name, color string, price float64, stock int, code string, published bool, creationDate string) (Product, error)
	Patch(id int, name, color string, price float64, stock int, code string, published bool, creationDate string) (Product, error)
	Delete(id int) error
}

func (r *repository) Store(id int, name, color string, price float64, stock int, code string, published bool, creationDate string) (Product, error) {

	var ps []Product

	r.db.Read(&ps) //writes json into ps struct

	prod := Product{id, name, color, price, stock, code, published, creationDate}

	ps = append(ps, prod)
	if err := r.db.Write(ps); err != nil {
		return Product{}, err
	}

	lastId = prod.Id

	return prod, nil
}

func (r *repository) Get(id int) (Product, error) {
	var ps []Product
	err := r.db.Read(&ps)
	if err != nil {
		return Product{}, err
	}
	for i := range ps {
		if ps[i].Id == id {
			return ps[i], nil
		}
	}
	return Product{}, fmt.Errorf("producto %d no encontrado", id)
}

func (r *repository) GetAll() ([]Product, error) {
	var ps []Product
	err := r.db.Read(&ps)
	if err != nil {
		return nil, err
	}
	return ps, nil
}

func (r *repository) LastID() (int, error) {
	//return lastId, nil
	var ps []Product
	err := r.db.Read(&ps)
	if err != nil {
		return 0, err
	}
	if len(ps) == 0 {
		return 0, nil
	}

	return ps[len(ps)-1].Id, nil
}

func (r *repository) Put(id int, name, color string, price float64, stock int, code string, published bool, creationDate string) (Product, error) {
	aux := Product{id, name, color, price, stock, code, published, creationDate}
	var ps []Product
	err := r.db.Read(&ps)
	if err != nil {
		return Product{}, err
	}
	for i := range ps {
		if ps[i].Id == id {
			ps[i] = aux
			//modificar en json
			return ps[i], nil
		}
	}
	return Product{}, fmt.Errorf("producto %d no encontrado", id)
}

func (r *repository) Patch(id int, name, color string, price float64, stock int, code string, published bool, creationDate string) (Product, error) {
	var ps []Product
	err := r.db.Read(&ps)
	if err != nil {
		return Product{}, err
	}
	for i := range ps {
		if ps[i].Id == id {
			if name != "" {
				ps[i].Name = name
			}
			if color != "" {
				ps[i].Color = color
			}
			if price != 0 {
				ps[i].Price = price
			}
			if stock != 0 {
				ps[i].Stock = stock
			}
			if code != "" {
				ps[i].Code = code
			}
			if published {
				ps[i].Published = published
			}
			if creationDate != "" {
				ps[i].CreationDate = creationDate
			}

			return ps[i], nil
		}
	}

	return Product{}, fmt.Errorf("producto %d no encontrado", id)
}

func (r *repository) Delete(id int) error {
	var ps []Product
	err := r.db.Read(&ps)
	if err != nil {
		return err
	}
	for i := range ps {
		if ps[i].Id == id {
			ps = append(ps[:i], ps[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("producto %d no encontrado", id)
}
