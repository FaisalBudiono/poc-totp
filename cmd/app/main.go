package main

import (
	"FaisalBudiono/poc-totp/internal/app/core/ascii"
	"bytes"
	"fmt"

	"github.com/pquerna/otp/totp"
)

type User struct {
	Name  string
	Email string
}

func main() {
	u := User{
		Email: "john@example.com",
	}

	k, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "Local Example",
		AccountName: u.Email,
	})
	if err != nil {
		panic(err)
	}

	img, err := k.Image(300, 300)
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	ascii.Draw(&buf, img)

	fmt.Printf("Kambing img %s\n", buf.String())
}
