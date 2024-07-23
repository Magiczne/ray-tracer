package color

func Black() *Color {
	return &Color{0, 0, 0}
}

func White() *Color {
	return &Color{1.0, 1.0, 1.0}
}

func Red() *Color {
	return &Color{1.0, 0, 0}
}

func Green() *Color {
	return &Color{0, 1.0, 0}
}

func Blue() *Color {
	return &Color{0, 0, 1.0}
}

func Cyan() *Color {
	return &Color{0, 1.0, 1.0}
}
