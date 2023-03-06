package car

type Speed float64

const (
	MPH Speed = 1
	KPH       = 1.60934
)

type Color string

const (
	BlueColor  Color = "blue"
	GreenColor       = "green"
	RedColor         = "red"
)

type Wheels string

const (
	SportsWheels Wheels = "sports"
	SteelWheels         = "steel"
)

type Builder interface {
	Color(Color) Builder
	Wheels(Wheels) Builder
	TopSpeed(Speed) Builder
	Build() Interface
}

type Interface interface {
	Drive() error
	Stop() error
}

type builder struct {
	color Color
	wheels Wheels
	topSpeed Speed
}

type Car struct {
	Color Color
	Wheels Wheels
	TopSpeed Speed
}

func (c *Car) Drive() error {
	return nil
}

func (c *Car) Stop() error {
	return nil
}

func (c *builder) Color(color Color) Builder {
	c.color = color
	return c
}

func (c *builder) Wheels(wheels Wheels) Builder{
	c.wheels = wheels
	return c
}

func (c *builder) TopSpeed(speed Speed) Builder{
	c.topSpeed = speed
	return c
}

func (c *builder) Build() Interface {
	return &Car{c.color, c.wheels, c.topSpeed}
}

func NewBuilder() Builder {
	return &builder{}
}


