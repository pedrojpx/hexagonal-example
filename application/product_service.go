package application

type ProductService struct {
	Persistence ProductPersistenceInterface
}

func (s *ProductService) NewProductService(persistence ProductPersistenceInterface) *ProductService {
	return &ProductService{
		Persistence: persistence,
	}
}

func (s *ProductService) Get(id string) (ProductInterface, error) {
	p, err := s.Persistence.Get(id)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (s *ProductService) Create(name string, price float64) (ProductInterface, error) {
	p := NewProduct()
	p.Name = name
	p.Price = price
	_, err := p.IsValid()
	if err != nil {
		return &Product{}, err
	}
	result, err := s.Persistence.Save(p)
	if err != nil {
		return &Product{}, err
	}
	return result, nil
}

func (s *ProductService) Enable(product ProductInterface) (ProductInterface, error) {
	err := product.Enable()
	if err != nil {
		return &Product{}, err
	}

	result, err := s.Persistence.Save(product)
	if err != nil {
		return &Product{}, err
	}
	return result, nil
}

func (s *ProductService) Disable(product ProductInterface) (ProductInterface, error) {
	err := product.Disable()
	if err != nil {
		return &Product{}, err
	}

	result, err := s.Persistence.Save(product)
	if err != nil {
		return &Product{}, err
	}
	return result, nil
}
