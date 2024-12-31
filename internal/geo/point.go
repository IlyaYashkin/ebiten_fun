package geo

type Point struct {
	X, Y float64
}

func (p *Point) Move(dir Vector) {
	p.X += dir.X
	p.Y += dir.Y
}

func (p *Point) GetVector(p2 Point) Vector {
	return Vector{p2.X - p.X, p2.Y - p.Y}
}

func (p *Point) ToVector() Vector {
	return Vector{p.X, p.Y}
}
