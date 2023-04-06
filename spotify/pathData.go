package spotify

import (
	"strings"
)

type PathData struct {
	Name string
	ID   string
}

func (pd *PathData) getPathData(path string) {
	p := strings.Split(path, "/")
	for i, value := range p {
		if i == 1 {
			pd.Name = value
		} else if i == 2 {
			pd.ID = value
		}
	}

}
