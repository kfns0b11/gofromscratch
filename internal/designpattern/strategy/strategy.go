package strategy

type IStrategy interface {
	Do(int, int) int
}

type Add struct{}

func (*Add) Do(a, b int) int {
	return a + b
}

type Reduce struct{}

func (*Reduce) Do(a, b int) int {
	return a - b
}

type Operator struct {
	Strategy IStrategy
}

func (operator *Operator) SetStrategy(strategy IStrategy) {
	operator.Strategy = strategy
}

func (operator *Operator) Caculate(a, b int) int {
	return operator.Strategy.Do(a, b)
}
