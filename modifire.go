package main

import "math"

// PriceModifier представляет любую модификацию цены: принимает цену и возвращает новую.
type PriceModifier func(float64) float64

// floatEqual сообщает, равны ли a и b с относительной точностью 1e-9.
func floatEqual(a, b float64) bool {
	const eps = 1e-9
	diff := math.Abs(a - b)
	scale := math.Max(math.Abs(a), math.Abs(b))
	return diff <= eps*math.Max(1, scale)
}

// WithDiscount возвращает модификатор, уменьшающий цену на percent процентов.
func WithDiscount(percent float64) PriceModifier {
	return func(price float64) float64 {
		return price * (1 - percent/100)
	}
}

// WithTax возвращает модификатор, увеличивающий цену на percent процентов налога.
func WithTax(percent float64) PriceModifier {
	return func(price float64) float64 {
		return price * (1 + percent/100)
	}
}

// WithRounding возвращает модификатор, округляющий цену до precision знаков после запятой.
func WithRounding(precision int) PriceModifier {
	return func(price float64) float64 {
		multiplier := math.Pow(10, float64(precision))
		return math.Round(price*multiplier) / multiplier
	}
}

// WithMinimumPrice возвращает модификатор, не дающий цене опуститься ниже minPrice.
func WithMinimumPrice(minPrice float64) PriceModifier {
	return func(price float64) float64 {
		if !floatEqual(price, minPrice) && price < minPrice {
			return minPrice
		}
		return price
	}
}

// CombineModifiers — функция высшего порядка: объединяет modifiers в один модификатор,
// который применяет их по очереди в переданном порядке.
func CombineModifiers(modifiers ...PriceModifier) PriceModifier {
	return func(price float64) float64 {
		result := price
		for _, mod := range modifiers {
			result = mod(result)
		}
		return result
	}
}
