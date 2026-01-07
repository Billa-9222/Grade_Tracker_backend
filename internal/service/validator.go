package service

import "regexp"

var (
	nameRegex = regexp.MustCompile(`^[A-Za-z]{2,}$`)
	surnameRegex = regexp.MustCompile(`^[A-Za-z]{3,}$`)
	classRegex = regexp.MustCompile(`^(?:[1-9]|1[0-2])[A-Z]$`)
)

func isValidName(is string) bool{
	return nameRegex.MatchString(is)
}

func isValidSurname(is string) bool{
	return surnameRegex.MatchString(is)
}

func isValidClass(is string) bool {
	return classRegex.MatchString(is)
}