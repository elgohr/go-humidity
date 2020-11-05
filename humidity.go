package humidity

import "math"

const gasConstantForSteam = 461.51

func ToAbsolute(temperature float64, relativeHumidity float64) float64 {
	saturationVaporPressure := magnusFormula(temperature)
	partialPressureForWater := relativeHumidity * saturationVaporPressure
	return partialPressureForWater / (gasConstantForSteam * temperature)
}

func magnusFormula(temperature float64) float64 {
	// see https://en.wikipedia.org/wiki/Vapour_pressure_of_water#Approximation_formulas
	return 0.61094 * math.Exp((17.625*temperature)/(243.04+temperature))
}
