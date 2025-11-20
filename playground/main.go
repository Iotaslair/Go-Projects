package main

import "strings"

func main() {

}

func IsLinkedInEmployee(email string) bool {
	return strings.HasSuffix(email, "@linkedin.com")
}
