package games

import "fmt"

type singleValResultType int

func (r singleValResultType) String() string {
	return fmt.Sprintf("%d", r)
}
