package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
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

func execCommand(commandName string, params []string) bool {
	//函数返回一个*Cmd，用于使用给出的参数执行name指定的程序
	cmd := exec.Command(commandName, params...)

	//显示运行的命令
	fmt.Println(cmd.Args)
	//StdoutPipe方法返回一个在命令Start后与命令标准输出关联的管道。Wait方法获知命令结束后会关闭这个管道，一般不需要显式的关闭该管道。
	stdout, err := cmd.StdoutPipe()

	if err != nil {
		fmt.Println(err)
		return false
	}

	cmd.Start()
	//创建一个流来读取管道内内容，这里逻辑是通过一行一行的读取的
	reader := bufio.NewReader(stdout)

	//实时循环读取输出流中的一行内容
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		fmt.Println(line)
	}

	//阻塞直到该命令执行完成，该命令必须是被Start方法开始执行的
	cmd.Wait()
	return true
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
