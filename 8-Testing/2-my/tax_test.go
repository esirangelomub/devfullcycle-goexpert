package tax

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateTax(t *testing.T) {
	tax, err := CalculateTax(1000)
	assert.Nil(t, err)
	assert.Equal(t, 10.0, tax, "Tax for 1000 should be 10.0")

	tax, err = CalculateTax(0)
	assert.NotNil(t, err, "amount must be greater than zero")
	assert.Equal(t, 0.0, tax)
	assert.Contains(t, err.Error(), "greater than zero")
}

func TestCalculateTaxAndSave(t *testing.T) {
	repository := &RepositoryMock{}
	repository.On("SaveTax", 10.0).Return(nil)
	repository.On("SaveTax", 0.0).Return(errors.New("error saving tax"))

	err := CalculateTaxAndSave(1000.0, repository)
	assert.Nil(t, err)

	err = CalculateTaxAndSave(0.0, repository)
	assert.Error(t, err, "error saving tax")

	repository.AssertExpectations(t)
}
