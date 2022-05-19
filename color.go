package main

type Color Vect3

func color_init(r float64, g float64, b float64) Color {
	return Color(vector_init(r, g, b))
}

func color_add(color1 Color, color2 Color) Color {
	return Color(vector_add(Vect3(color1), Vect3(color2)))
}

func color_scalar_mul(color Color, t float64) Color {
	return Color(vector_scalar_mul(Vect3(color), t))
}
