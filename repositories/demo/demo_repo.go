package demoRepo

import (
	"fmt"
	db "go-api-template/bootstrap"
)

// GetDemoRepo implements interfaces.MemberRepo.
func GetDemoRepo(value string) (string, error) {
	test := db.Conn.Row()
	fmt.Println(test)
	return "", nil
}
