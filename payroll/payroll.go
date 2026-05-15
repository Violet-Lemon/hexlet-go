package payroll

import (
	"github.com/Violet-Lemon/hexlet-go/v2/employees"
)

func CalcGross(e employees.Employee, bonusPct float64) float64 {
	return e.BaseSalary() * (1 + bonusPct / 100)
}

func CalcNet(gross float64, taxPct float64) float64 {
	return gross * (1 - taxPct / 100)
}
