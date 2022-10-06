package products

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

type Service interface {
	GetAll() ([]Product, error)
	Store(name, color string, price float64, stock int, code string, published bool, creationDate string) (Product, error)
	//Update(id int, name , color string, price float64, stock int, code string, published bool, creationDate string) (Product, error)
}

func (s *service) Store(name, color string, price float64, stock int, code string, published bool, creationDate string) (Product, error) {
	product, err := s.repository.Store(name, color, price, stock, code, published, creationDate)
	if err != nil {
		return Product{}, err
	}
	return product, nil
}

func (s *service) GetAll() ([]Product, error) {
	products, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return products, nil
}
