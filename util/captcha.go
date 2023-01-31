package util

import (
	"github.com/mojocn/base64Captcha"
)

// TODO: Redis store
var store = base64Captcha.DefaultMemStore

func GenerateCaptcha() (string, string, error) {
	driver := base64Captcha.NewDriverDigit(60, 200, 6, 0.7, 80)

	captcha := base64Captcha.NewCaptcha(driver, store)

	id, b64s, err := captcha.Generate()

	return id, b64s, err
}

func VerifyCaptcha(id, answer string) bool {
	return store.Verify(id, answer, true)
}
