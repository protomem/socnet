package port

type Pair[Left any, Right any] struct {
	left  Left
	right Right
}

func NewPair[Left any, Right any](left Left, right Right) Pair[Left, Right] {
	return Pair[Left, Right]{
		left:  left,
		right: right,
	}
}

func (e Pair[Left, Right]) Left() Left {
	return e.left
}

func (e Pair[Left, Right]) Right() Right {
	return e.right
}
