package utils

import (
	"github.com/mojocn/base64Captcha"
)

func CodeCaptchaCreate(height, width int) (string, string) {
	var configC = base64Captcha.ConfigCharacter{
		Height: height,
		Width:  width,
		//CaptchaModeNumber:数字,CaptchaModeAlphabet:字母,CaptchaModeArithmetic:算术,CaptchaModeNumberAlphabet:数字字母混合.
		Mode:               base64Captcha.CaptchaModeArithmetic,
		ComplexOfNoiseText: base64Captcha.CaptchaComplexLower,
		ComplexOfNoiseDot:  base64Captcha.CaptchaComplexLower,
		IsUseSimpleFont:    true,
		IsShowHollowLine:   true,
		IsShowNoiseDot:     true,
		IsShowNoiseText:    true,
		IsShowSlimeLine:    true,
		IsShowSineLine:     true,
		CaptchaLen:         4,
	}
	//GenerateCaptcha first parameter is empty string,so the package will generate a random uuid for you.
	idKeyC, capC := base64Captcha.GenerateCaptcha("", configC)
	base64stringC := base64Captcha.CaptchaWriteToBase64Encoding(capC)
	return idKeyC, base64stringC
}

func VerfiyCaptcha(idkey, verifyValue string) bool {
	verifyResult := base64Captcha.VerifyCaptcha(idkey, verifyValue)
	if verifyResult {
		return true
	}
	return false
}
