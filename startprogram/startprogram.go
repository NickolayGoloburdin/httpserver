package startprogram

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

//1-start 2-end
func StartProgram(path string, state chan int, status *string) {
	if len(state) != 0 {
		return
	}
	state <- 1
	cmd := exec.Command(path)
	//cmdOutput := &bytes.Buffer{}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Start()
	*status = "loading"
	//var once bool
	//var res bool
	time.Sleep(time.Millisecond * 300)
	*status = "Enable"
	// for len(state) != 0 {
	// 	output := cmdOutput.Bytes()

	// 	if parseOutput(output, &res) {
	// 		*status = "enabled"

	// 		break
	// 	}
	// 	fmt.Println("Finding")

	// 	printError(err)
	// 	//printOutput(output)
	// 	time.Sleep(time.Millisecond * 300)

	// }
	cmd.Wait()
}
func EndProgram(path string, state chan int, status *string) {

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
func printCommand(cmd *exec.Cmd) {
	fmt.Printf("==> Executing: %s\n", strings.Join(cmd.Args, " "))
}

func printOutput(outs []byte) {
	if len(outs) > 0 {
		fmt.Printf("==> Output: %s\n", string(outs))
	}
}

func parseOutput(outs []byte, res *bool) bool {
	matched, _ := regexp.MatchString(`position finder ready`, string(outs))
	return matched

}
