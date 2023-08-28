package main

import "fmt"

type Status int

const (
	InValidLogin Status = iota + 1
	NotFound
)

type StatusErr struct {
	Status  Status
	Message string
}

func (se StatusErr) Error() string {
	return se.Message
}

func LoginAndGetData(uid, pwd, file string) ([]byte, error) {
	err := login(uid, pwd)
	if err != nil {
		return nil, StatusErr{
			Status:  InValidLogin,
			Message: fmt.Sprintf("invalid credentials for %s", uid),
		}
	}
	data, err := getData(file)
	if err != nil {
		return nil, StatusErr{
			Status:  NotFound,
			Message: fmt.Sprintf("%s not found", file),
		}
	}
	return data, nil
}

func login(uid, pwd string) error {
	return nil
}

func getData(file string) ([]byte, error) {
	var dummy []byte
	return dummy, nil
}
