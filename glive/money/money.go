package glive

import "github.com/aimey-a/go-tools/gtype"

var (
	year  float64 = 365
	ratio float64 = 100
)

//计算利息 money 总金额  annualYield 年化
func CalculationOfInterest[T gtype.BaseTypeNumber](money, annualYield T) T {
	return T(float64(money) * (float64(annualYield) / ratio) / year)
}

//推算利息 money 总金额  annual 利息
func CalculateAccrual[T gtype.BaseTypeNumber](money, annual T) T {
	return T(float64(annual) * year / float64(money) * ratio)
}
