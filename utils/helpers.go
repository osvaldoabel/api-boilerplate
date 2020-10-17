package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
)

/**
	Description: Verify if a string is a valid email
		and return true/false depending if it matches
		the email pattern or not.
	Return: true|false
**/
func IsValidEmail(email string) bool {
	pattern := regexp.MustCompile(`^(\w|\.)+@(\w)+(.(\w)+){1,2}$`)
	return pattern.MatchString(email)
}

/**
	Description: Receives a variable, parse to json
		and prints it in a pretty way
	Return: void
**/
func Dd(variable interface{}) {
	res, _ := json.MarshalIndent(variable, "", "  ")

	fmt.Println("= = = = = = = = = = = =")
	fmt.Println(string(res))
	fmt.Println("= = = = = = = = = = = =")
	os.Exit(1)
}
