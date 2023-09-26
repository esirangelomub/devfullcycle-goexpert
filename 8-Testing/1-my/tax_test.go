package tax

import "testing"

func TestCalculateTax(t *testing.T) {
	amount := 500.00
	expected := 5.0

	result := CalculateTax(amount)

	if result != expected {
		t.Errorf("Expected %.2f, got %.2f", expected, result)
	}
}

func TestCalculateTaxBatch(t *testing.T) {
	type calcTax struct {
		amount, expected float64
	}

	table := []calcTax{
		{amount: 500, expected: 5.0},
		{amount: 628, expected: 5.0},
		{amount: 1000, expected: 10.0},
		{amount: 1500, expected: 10.0},
		{amount: 2000, expected: 10.0},
		{amount: 0, expected: 0},
	}

	for _, c := range table {
		result := CalculateTax(c.amount)
		if result != c.expected {
			t.Errorf("Expected %.2f, got %.2f", c.expected, result)
		}
	}
}

func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(500)
	}
}

func BenchmarkCalculateTax2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax2(500)
	}
}

func FuzzCalculateTax(f *testing.F) {
	seed := []float64{-1, -2, -2.5, 500, 628, 1000, 1500, 2000, 0, 1501.12}
	for _, amount := range seed {
		f.Add(amount)
	}
	f.Fuzz(func(t *testing.T, amount float64) {
		result := CalculateTax(amount)
		if amount <= 0 && result != 0 {
			t.Errorf("Reveived %f but expected 0", result)
		}
		if amount > 20000 && result != 20 {
			t.Errorf("Reveived %f but expected 20", result)
		}
	})
}
