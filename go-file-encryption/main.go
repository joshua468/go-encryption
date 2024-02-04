package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/joshua468/go-file-encryption/filecrypt"
	"golang.org/x/term"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(0)
	}
	function := os.Args[1]

	switch function {
	case "help":
		printHelp()
	case "encrypt":
		encryptHandle()
	case "decrypt":
		decryptHandle()

	default:
		fmt.Println("run encrypt to encrypt a file,and decrypt to decrypt a file.")
		os.Exit(1)
	}
}
func printHelp() {
	fmt.Println("file encryption")
	fmt.Println("simple file encrpter for your day-to-day needs")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("")
	fmt.Println("\t  go run .  encrypt /path/to/your/file")
	fmt.Println("")
	fmt.Println("Commands")
	fmt.Println("")
	fmt.Println("\t encrpyt\tEncrypt a file  using a password")
	fmt.Println("\t decrpyt\t Tries to decrypt a file using a password")
	fmt.Println("\t help \t\t Displays help text")
	fmt.Println()
}
func encryptHandle() {
	if len(os.Args) < 3 {
		fmt.Println("missing path to the file. for  more info, run")
		os.Exit(0)
	}
	file := os.Args[2]
	if !validateFile(file) {
		panic("file not found")
	}
	password := getpassword()
	fmt.Println("\nEncrypting...")
	filecrypt.Encrypt(file, password)
	fmt.Println("\n file successfully protected")
}

func decryptHandle() {
	if len(os.Args) < 3 {
		fmt.Println("missing the path to the file. For more info, run go run . help")
		os.Exit(0)
	}
	file := os.Args[2]
	if !validateFile(file) {
		panic("file not found")
	}
	fmt.Print("Enter password:")
	password, _ := term.ReadPassword(0)
	fmt.Println("\nDecrypting...")
	filecrypt.decrpyt(file, password)
	fmt.Println("file successfully decrypted")

}

func getpassword() []byte {
	fmt.Print("Enter password")
	password, _ := term.Readpassword(0)
	fmt.Print("\n confirm Password")
	password2, _ := term.ReadPassword(0)
	if !validatePassword(password, password2) {
		fmt.Printf("\nPasswords do not match. please try again\n")
		return getpassword()
	}
	return password
}
func validatePassword(password1 []byte, password2 []byte) bool {
	if !bytes.Equal(password1, password2) {
		return false
	}
	return true

}

func validateFile(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}
