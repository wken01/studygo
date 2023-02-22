package main

import (
	"fmt"
	"os/exec"
)

//管道是一种半双工（单向）的通信方式，只能用于父进程与子进程，以及同祖先的子进程之间的通信。
//管道的优点就是简单，而缺点就是只能单向通信以及通信双方关系上的严格限制
func main() {
	cmd0 := exec.Command("echo", "-n", "My first command comes from golang.")
	//获取cmd命令的输出管道
	stdout0, err := cmd0.StdoutPipe()
	if err != nil {
		fmt.Printf("Error: couldn't obtain the stdout pipe for command No.0 %s\n", err)
		return
	}

	if err := cmd0.Start(); err != nil {
		fmt.Print("Error: The command No.0 can not be startup: %s\n", err)
		return
	}



	output0 := make([]byte, 30)
	n, err := stdout0.Read(output0)
	if err != nil {
		fmt.Printf("Error: couldn't read data from the pipe %s\n", err)
		return
	}

	fmt.Printf("%s\n",output0[:n])


}
