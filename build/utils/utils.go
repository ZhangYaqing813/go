package utils

import (
	"fmt"
	"os"
	"path"
)

type Proj struct {
	Name       string
	Repository string
	Target_dir string
	Date       string
}

func Fat(name string, repo string, target string, date string) *Proj {
	return &Proj{
		name,
		repo,
		target,
		date,
	}
}

func (p *Proj) Get(proj *Proj) {
	build_dir := path.Join("/data/porj/", proj.Name, "/", proj.Date)
	_, err := os.Stat(build_dir)
	if err != nil {
		fmt.Printf("==== %v\n", err)
		if os.IsNotExist(err) {
			fmt.Println("当前文件不存在，准备创建")
			os.MkdirAll(build_dir, os.ModePerm)
		}
	}

}
