package main

import (
	"FaisalBudiono/poc-totp/internal/app/core/ascii"
	"FaisalBudiono/poc-totp/internal/app/core/tfa"
	"bufio"
	"bytes"
	"fmt"
	"image"
	"os"
)

func main() {
	accountName := "john@example.com"

	totp := tfa.NewTOTP()

	key, err := totp.Generate(accountName)
	if err != nil {
		panic(err)
	}

	printImage(key.Image())

	fmt.Println("Validating TOTP...")
	passcode := promptForPasscode()
	valid := totp.Validate(key.Secret(), passcode)
	for !valid {
		fmt.Println("INVALID passcode!")
		fmt.Println("Try again...")
		passcode = promptForPasscode()
		valid = totp.Validate(key.Secret(), passcode)
	}

	println("Valid passcode!")
	os.Exit(0)
}

func promptForPasscode() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Passcode: ")
	text, _ := reader.ReadString('\n')
	return text
}

func printImage(img image.Image) {
	var bufImg bytes.Buffer
	ascii.Draw(&bufImg, img)

	fmt.Println("Scan the QR code with your OTP App")
	fmt.Println("=============")
	fmt.Printf("%s\n", bufImg.String())
	fmt.Println("=============")
	fmt.Println("=============")
	fmt.Println("=============")
}
