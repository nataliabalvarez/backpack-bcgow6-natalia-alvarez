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
	Get(id int) (Product, error)
	GetAll() ([]Product, error)
	Store(name, color string, price float64, stock int, code string, published bool, creationDate string) (Product, error)
	Put(id int, name, color string, price float64, stock int, code string, published bool, creationDate string) (Product, error)
	Patch(id int, name, color string, price float64, stock int, code string, published bool, creationDate string) (Product, error)
	Delete(id int) error
}

func (s *service) Store(name, color string, price float64, stock int, code string, published bool, creationDate string) (Product, error) {

	lastId, err := s.repository.LastID()
	if err != nil {
		return Product{}, err
	}

	lastId++

	product, err := s.repository.Store(lastId, name, color, price, stock, code, published, creationDate)
	if err != nil {
		return Product{}, err
	}
	return product, nil
}

func (s *service) Get(id int) (Product, error) {
	return s.repository.Get(id)
}

func (s *service) GetAll() ([]Product, error) {
	products, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *service) Put(id int, name, color string, price float64, stock int, code string, published bool, creationDate string) (Product, error) {

	product, err := s.repository.Put(id, name, color, price, stock, code, published, creationDate)

	if err != nil {
		return Product{}, err
	}

	return product, nil
}

func (s *service) Patch(id int, name, color string, price float64, stock int, code string, published bool, creationDate string) (Product, error) {
	product, err := s.repository.Patch(id, name, color, price, stock, code, published, creationDate)
	if err != nil {
		return Product{}, err
	}

	return product, nil
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}