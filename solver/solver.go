package solver

type Solver struct {
	Run func([]string) (string, error)
}
