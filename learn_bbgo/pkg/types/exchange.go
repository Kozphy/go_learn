package types

type ExchangeName string

func (n ExchangeName) String() string {
	return string(n)
}
