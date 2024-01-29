package main

import (
	"fmt"
	"os/exec"
	"time"
)

func (cs *CommandService) ExecuteCommand(request CommandRequest, response *CommandResponse) error {
	if !authenticateUser(request.Username, request.Password) {
		response.Errno = 1
		response.Error = "authentication failed"
		return nil
	}

	startTime := time.Now()

	cmd := exec.Command(request.Command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		response.Errno = 2
		response.Error = fmt.Sprintf("command execution error: %s", err)
		return nil
	}

	err = cmd.Wait()

	response.Errno = 0
	response.Error = "succeed"
	response.Data = CommandData{
		Status:      "finished",
		Stdout:      string(output),
		ExitCode:    getExitCode(err),
		CreateTime:  startTime,
		FinishTime:  time.Now(),
		Environment: getSystemEnvironment(),
	}

	return nil
}
