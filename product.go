package main

// Product is a structure that describes a single product
type Product struct {
	Backend // Backend is being declared as part of the Product structure

	Name            string
	Price           float32
	BillingInterval string
}

// Backend defines some custom network behavior we like all HTTP clients to implement
type Backend interface {
	Call(string) (string, error) // makes an API call
}

// Call could be an API call to get some data and is expected to return a simple string and an error state.
// This func can be used as long as the item using it satisfies the *Product object.
// This means any Product object can use `.Call()`
func (p *Product) Call(res string) (string, error) {
	return p.Name, nil
}

// NewProductParams is the defined dependency to CreateNewProduct
type NewProductParams struct {
	Name     string
	Interval string
	Price    float32
}

// CreateNewProduct creates a new product for us from the defined input
func CreateNewProduct(params *NewProductParams) *Product {
	product := &Product{
		Name:            params.Name,
		Price:           params.Price,
		BillingInterval: params.Interval,
	}

	return product
}

// SomeProductFunc is usable for anything that satisfies the Backend interface
func SomeProductFunc(product Backend) (string, error) {
	// use our call method
	res, err := product.Call("something")
	if err != nil {
		return "", err
	}

	return res, nil
}
