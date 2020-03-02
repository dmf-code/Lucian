package captcha

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"image/color"
	"sync"
)

//configJsonBody json request body.
type CaptchaConfig struct {
	Id            string
	CaptchaType   string
	VerifyValue   string
	DriverAudio   *base64Captcha.DriverAudio
	DriverString  *base64Captcha.DriverString
	DriverChinese *base64Captcha.DriverChinese
	DriverMath    *base64Captcha.DriverMath
	DriverDigit   *base64Captcha.DriverDigit
}

var store = base64Captcha.DefaultMemStore
var captchaConfigOnce sync.Once
var captchaConfig *CaptchaConfig
// 获取base64验证码基本配置
func GetCaptchaConfig() *CaptchaConfig {
	captchaConfigOnce.Do(func() {
		captchaConfig = &CaptchaConfig{
			Id:              "",
			CaptchaType:     "string",
			VerifyValue:     "",
			DriverAudio:     base64Captcha.DefaultDriverAudio,
			DriverString: 	 &base64Captcha.DriverString{
			Height:          30,
			Width:           60,
			NoiseCount:      0,
			ShowLineOptions: 2 | 4,
			Length:          4,
			Source:          "1234567890qwertyuioplkjhgfdsazxcvbnm",
			BgColor: &color.RGBA{
				R: 3,
				G: 102,
				B: 214,
				A: 125,
			},
			Fonts: []string{"wqy-microhei.ttc"},
		},
			DriverDigit:     base64Captcha.DefaultDriverDigit,
		}
	})
	return captchaConfig
}

// base64Captcha create http handler
func GenerateCaptchaHandler(c *gin.Context) {

	var driver base64Captcha.Driver
	var param CaptchaConfig
	err := c.BindJSON(&param)
	if err != nil {
		c.JSON(200, gin.H{"code": 400, "msg": "failed"})
	}
	//create base64 encoding captcha
	switch param.CaptchaType {
	case "audio":
		driver = GetCaptchaConfig().DriverAudio
	case "string":
		driver = GetCaptchaConfig().DriverString.ConvertFonts()
	case "math":
		driver = GetCaptchaConfig().DriverMath.ConvertFonts()
	case "chinese":
		driver = GetCaptchaConfig().DriverChinese.ConvertFonts()
	default:
		driver = GetCaptchaConfig().DriverDigit
	}
	captcha := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := captcha.Generate()
	body := gin.H{"code": 1, "data": b64s, "captchaId": id, "msg": "success"}
	if err != nil {
		body = gin.H{"code": 0, "msg": err.Error()}
	}

	c.JSON(200, body)
}

// base64Captcha verify http handler
func VerifyCaptchaHandler(id string, verifyValue string) bool {

	//verify the captcha
	if store.Verify(id, verifyValue, true) {
		return true
	}

	return false
}

