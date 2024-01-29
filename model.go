package main

import "time"

type CommandService struct{}

var AuthCredentials = map[string]string{
	"username": "password",
}

type CommandRequest struct {
	Username string
	Password string
	Command  string
}

type CommandResponse struct {
	Errno int    `json:"errno"`
	Error string `json:"error"`
	Data  CommandData
}

type CommandData struct {
	Status      string            `json:"status"`
	Stdout      string            `json:"stdout"`
	ExitCode    int               `json:"exit_code"`
	CreateTime  time.Time         `json:"create_time"`
	FinishTime  time.Time         `json:"finish_time"`
	Environment map[string]string `json:"environment"`
}
