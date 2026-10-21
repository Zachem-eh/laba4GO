package main

import "math"

type PriceModifierFunc func(float64) float64

func WithDiscount(percent float64) PriceModifierFunc {
	return func(price float64) float64 {
		return price * (1 - percent/100)
	}
}
func WithTax(percent float64) PriceModifierFunc {
	return func(price float64) float64 {
		return price * (1 + percent/100)
	}
}
func WithRounding(precision int) PriceModifierFunc {
	return func(price float64) float64 {
		multiplier := math.Pow(10, float64(precision))
		return math.Round(price*multiplier) / multiplier
	}
}
func WithMinimumPrice(minPrice float64) PriceModifierFunc {
	return func(price float64) float64 {
		if price < minPrice {
			return minPrice
		}
		return price
	}
}

// CombineModifiers функция высшего порядка
func CombineModifiers(modifiers ...PriceModifierFunc) PriceModifierFunc {
	return func(price float64) float64 {
		result := price
		for _, mod := range modifiers {
			result = mod(result)
		}
		return result
	}
}
