package folders

import (
	"io/ioutil"
)

type Path struct {
	path string
}

func NewPath(p string) *Path {
	return &Path{path: p}
}

func (p Path) GetFolders() ([]string, error) {
	var folders []string

	files, err := ioutil.ReadDir(p.path)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			//x := file.Name()
			folders = append(folders, file.Name())
			//fmt.Printf("%T\n", x)
			//fmt.Println(file.Name())
		}
	}
	return folders, nil
}
