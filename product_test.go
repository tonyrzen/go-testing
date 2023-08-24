package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockProduct is a fake struct that parries our Product struct, but with properties from mock.Mock
type MockProduct struct {
	mock.Mock
	Name string
}

// Call implements a mock version of Call from the Backend interface onto our MockProduct
func (mp *MockProduct) Call(res string) (string, error) {
	// let's assert this call has happened in our test
	args := mp.Called(res)

	return args.Get(0).(string), args.Error(1)
}

// TestCall is verifying the implementation our Call receiver matches expectations based on the
// Backend interface
func TestCall(t *testing.T) {
	fakeProduct := "my_product"

	// create a mock Product
	mockProduct := &MockProduct{
		Name: fakeProduct,
	}

	mockProduct.On("Call", mock.Anything).
		Return(fakeProduct, nil).
		Once()

	res, err := mockProduct.Call(mockProduct.Name)
	assert.Nil(t, err)
	assert.Equal(t, fakeProduct, res)
}

// TestSomeProductFunc is validating the functionality of SomeProductFunc
func TestSomeProductFunc(t *testing.T) {
	fakeProduct := "some product"

	// create a mock Product
	mockProduct := &MockProduct{
		Name: fakeProduct,
	}

	mockProduct.On("Call", mock.Anything).
		Return(fakeProduct, nil).
		Once()

	res, err := SomeProductFunc(mockProduct)
	assert.Nil(t, err)
	assert.Equal(t, fakeProduct, res)
}

// TestSomeProductFunc_Failure is validating the functionality of FreeConsult func
func TestSomeProductFunc_Failure(t *testing.T) {
	fakeProduct := "some product"
	expectedError := "something went wrong"

	// create a mock Product
	mockProduct := &MockProduct{
		Name: fakeProduct,
	}

	mockProduct.On("Call", mock.Anything).
		Return("", errors.New(expectedError)).
		Once()

	res, err := SomeProductFunc(mockProduct)

	assert.NotNil(t, err)
	assert.Equal(t, expectedError, err.Error())
	assert.Equal(t, "", res)
}

// TestCall_Failure verifies our return in the case of a Call error
func TestCall_Failure(t *testing.T) {
	fakeProduct := "my_product"
	expectedError := "something went wrong"
	// create a mock Product
	mockProduct := &MockProduct{
		Name: fakeProduct,
	}

	mockProduct.On("Call", mock.IsType("string")).
		Once().
		Return("", errors.New(expectedError))

	res, err := mockProduct.Call(mockProduct.Name)
	assert.NotNil(t, err)
	assert.Equalf(t, "something went wrong", err.Error(), "error must return correct string")
	assert.Equal(t, "", res)
}

// TestCreateNewProduct_Mock is testing the return of CreateNewProduct
func TestCreateNewProduct_Mock(t *testing.T) {
	newProductParams := &NewProductParams{
		Name:     "Tofu Tacos",
		Interval: "daily",
		Price:    15.99,
	}
	res := CreateNewProduct(newProductParams)

	assert.IsTypef(t, (*Product)(nil), res, "res should be of type Product")
	assert.Equal(t, res.Name, newProductParams.Name)
}

// TestCreateNewProduct is validating our CreateNewProduct func can return a successful type that passes into SomeProductFunc
func TestCreateNewProduct(t *testing.T) {
	newProductParams := &NewProductParams{
		Name:     "Tofu Tacos",
		Interval: "daily",
		Price:    15.99,
	}
	product := CreateNewProduct(newProductParams)

	res, err := SomeProductFunc(product)
	assert.Nil(t, err)
	assert.Equal(t, res, newProductParams.Name)
}

func TestBackendImplementation(t *testing.T) {
	mp := &MockProduct{}
	assert.Implementsf(t, (*Backend)(nil), mp, "MockProduct MUST implement the Backend interface")
}
