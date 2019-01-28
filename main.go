package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	fmt.Println("This is the value specified for the input 'example_step_input':", os.Getenv("example_step_input"))

	versioncode:= os.Getenv("input_version_code")
	versionname:= os.Getenv("input_version_name")

	t:= time.Now()
	timeString:= t.Format("0601021504")

	fmt.Println("input version code: 'versioncode' and versionname 'versionname'",versioncode,versionname)


	//
	// --- Step Outputs: Export Environment Variables for other Steps:
	// You can export Environment Variables for other Steps with
	//  envman, which is automatically installed by `bitrise setup`.
	// A very simple example:
	cmdLog, err := exec.Command("bitrise", "envman", "add", "--key", "OUTPUT_VERSION_NAME", "--value", timeString).CombinedOutput()
	if err != nil {
		fmt.Printf("Failed to expose output with envman, error: %#v | output: %s", err, cmdLog)
		os.Exit(1)
	}
	// You can find more usage examples on envman's GitHub page
	//  at: https://github.com/bitrise-io/envman

	//
	// --- Exit codes:
	// The exit code of your Step is very important. If you return
	//  with a 0 exit code `bitrise` will register your Step as "successful".
	// Any non zero exit code will be registered as "failed" by `bitrise`.
	os.Exit(0)
}
