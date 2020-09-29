package objects

type Consumer interface {
	Accept(arg interface{})
}
