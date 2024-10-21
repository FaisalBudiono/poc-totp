package tfa

import (
	"image"
	"time"

	libtotp "github.com/pquerna/otp/totp"
)

type totpKey struct {
	issuer      string
	accountName string
	secret      string
	image       image.Image
}

func (c *totpKey) Issuer() string {
	return c.issuer
}

func (c *totpKey) AccountName() string {
	return c.accountName
}

func (c *totpKey) Secret() string {
	return c.secret
}

func (c *totpKey) Image() image.Image {
	return c.image
}

type totp struct{}

func NewTOTP() *totp {
	return &totp{}
}

func (t *totp) Generate(accountName string) (*totpKey, error) {
	k, err := libtotp.Generate(libtotp.GenerateOpts{
		Issuer:      "Local Example",
		AccountName: accountName,
	})
	if err != nil {
		return nil, err
	}

	img, err := k.Image(200, 200)
	if err != nil {
		return nil, err
	}

	return &totpKey{
		issuer:      k.Issuer(),
		accountName: k.AccountName(),
		secret:      k.Secret(),
		image:       img,
	}, nil
}

func (t *totp) Validate(secret string, passcode string) bool {
	valid, _ := libtotp.ValidateCustom(passcode, secret, time.Now(), libtotp.ValidateOpts{
		Digits: 6,
	})
	return valid
}
