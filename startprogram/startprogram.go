package startprogram

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"time"
)

//1-start 2-end
func StartProgram(path string, state chan int, status *string) {
	if len(state) != 0 {
		return
	}
	state <- 1
	cmd := exec.Command(path)
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	err := cmd.Start()

	var once bool
	for len(state) != 0 {
		// output := cmdOutput.Bytes()

		// if parseOutput(output, &res) && !once {
		// 	once = true
		// }
		if !once {
			once = true
			*status = "loading"
		} else {
			*status = "enabled"
		}
		printError(err)
		printOutput(string(cmdOutput.Bytes()))
		time.Sleep(time.Millisecond * 2000)

	}
}
func EndProgram(path string, state chan int, status *string) {

	if len(state) == 0 {
		return
	}

	*status = "loading"
	cmd := exec.Command(path)
	err := cmd.Run()
	*status = "disabled"
	<-state
	printError(err)

}

func printError(err error) {
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("==> Error: %s\n", err.Error()))
	}
}
func printOutput(output string) {
	if output != "" {
		os.Stdout.WriteString(fmt.Sprintf(output))
	}
}

func parseOutput(outs []byte, res *bool) bool {
	matched, _ := regexp.MatchString(`True`, string(outs))
	*res = matched
	return matched

}
