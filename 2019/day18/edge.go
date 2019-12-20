package day18

import "fmt"

type edge struct {
	a tile
	b tile
}

// hashKey ...
func (e edge) hashKey() string {
	// if e.a == nil && e.b == nil {
	// 	return fmt.Sprintf("edge<nil <-> nil>")
	// }

	// if e.a == nil && e.b != nil {
	// 	return fmt.Sprintf("edge<nil <-> %v>", e.b.hashKey())
	// }

	// if e.a != nil && e.b == nil {
	// 	return fmt.Sprintf("edge<%v <-> nil>", e.a.hashKey())
	// }

	return fmt.Sprintf("edge<%v <-> %v>", e.a.hashKey(), e.b.hashKey())
}
