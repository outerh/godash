package main


type Iterator[T any] struct{}
type Mapper[T, R any] struct {
	it Iterator[T]
	fn func(T) R
}

func (xs List[T]) filter(fn func(T) bool) List[T] {
	r := make(List[T], 0, len(xs)/2)
	for _, e := range xs {
		if fn(e) {
			r = append(r, e)
		}
	}
	return r
}



func main(){
	x := []int{1, 2, 3}
	// y := List[int](x)
	// x[0] = 4
	// y = y.filter(func (e int) bool { return e % 2 == 0})

	it := newIterator(x)
	r := map(it, func(int) float64 { return 2 * 2.})

	dash(x, [map(func(int) float64 { return 2 * 2.})])

	// fmt.Println(x)
	// fmt.Println(y)
}