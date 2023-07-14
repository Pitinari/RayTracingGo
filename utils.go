package main

func clamp(x float64, min float64, max float64) float64 {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}

func smoothstep(x float64, min float64, max float64) float64 {
	x = clamp((x-min)/(max-min), min, max)
	return x * x * (3.0 - 2.0*x)
}
