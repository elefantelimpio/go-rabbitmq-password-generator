package main

import (
	"fmt"
	"go-rabbitmq-password-generator/pwdgenerator"
)

var srv pwdgenerator.PwdIService

func init() {
	srv = pwdgenerator.NewPwdService()
}

func main() {
	var pwd string

	_, err := fmt.Scan(&pwd)
	if err != nil {
		panic(err)
	}
	hashPwd, err := srv.GetSHA256Hash(pwd)
	if err != nil {
		panic(err)
	}

	fmt.Println("hashPwd", hashPwd)
}
