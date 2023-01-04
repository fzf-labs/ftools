package email

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"time"
)

type Website struct {
	WebsiteURL    string `json:"home_url"`
	WebsiteName   string `json:"website_name"`
	WebsiteDomain string `json:"website_domain"`
}

func NewWebsite(websiteURL string, websiteName string, websiteDomain string) *Website {
	return &Website{WebsiteURL: websiteURL, WebsiteName: websiteName, WebsiteDomain: websiteDomain}
}

// ActiveUserMailData 激活用户模板数据
type ActiveUserMailData struct {
	UserName      string `json:"user_name"`
	WebsiteURL    string `json:"home_url"`
	WebsiteName   string `json:"website_name"`
	WebsiteDomain string `json:"website_domain"`
	ActivateURL   string `json:"activate_url"`
	Year          int    `json:"year"`
}

// ActivationHTMLEmail 发送激活邮件 html
func (w *Website) ActivationHTMLEmail(username, activateURL string) (subject string, body string) {
	mailData := ActiveUserMailData{
		UserName:      username,
		WebsiteURL:    w.WebsiteURL,
		WebsiteName:   w.WebsiteName,
		WebsiteDomain: w.WebsiteDomain,
		ActivateURL:   activateURL,
		Year:          time.Now().Year(),
	}
	mailTplContent := getEmailHTMLContent("pkg/email/templates/active-mail.html", mailData)
	return "帐号激活链接", mailTplContent
}

// VerificationCodeData
// @Description: 邮件验证码
type VerificationCodeData struct {
	WebsiteURL    string `json:"home_url"`
	WebsiteName   string `json:"website_name"`
	WebsiteDomain string `json:"website_domain"`
	Code          string `json:"code"`
	Greeting      string `json:"greeting"`
	Intro         string `json:"intro"`
	Outro         string `json:"outro"`
	Year          int    `json:"year"`
}

// VerificationCode 邮件验证码
func (w *Website) VerificationCode(code string, greeting string, intro string, outro string) (subject string, body string) {
	mailData := &VerificationCodeData{
		WebsiteURL:    w.WebsiteURL,
		WebsiteName:   w.WebsiteName,
		WebsiteDomain: w.WebsiteDomain,
		Code:          code,
		Greeting:      greeting,
		Intro:         intro,
		Outro:         outro,
		Year:          time.Now().Year(),
	}
	mailTplContent := getEmailHTMLContent("pkg/email/templates/verification-code.html", mailData)
	return "验证码", mailTplContent
}

// ResetPasswordMailData 激活用户模板数据
type ResetPasswordMailData struct {
	UserName      string `json:"user_name"`
	WebsiteURL    string `json:"home_url"`
	WebsiteName   string `json:"website_name"`
	WebsiteDomain string `json:"website_domain"`
	ResetURL      string `json:"reset_url"`
	Year          int    `json:"year"`
}

// ResetPasswordHTMLEmail 发送重置密码邮件
func (w *Website) ResetPasswordHTMLEmail(username, resetURL string) (subject string, body string) {
	mailData := ResetPasswordMailData{
		WebsiteURL:    w.WebsiteURL,
		WebsiteName:   w.WebsiteName,
		WebsiteDomain: w.WebsiteDomain,
		UserName:      username,
		ResetURL:      resetURL,
		Year:          time.Now().Year(),
	}
	mailTplContent := getEmailHTMLContent("pkg/email/templates/reset-mail.html", mailData)
	return "密码重置", mailTplContent
}

type NotifyMailData struct {
	UserName      string `json:"user_name"`
	WebsiteURL    string `json:"home_url"`
	WebsiteName   string `json:"website_name"`
	WebsiteDomain string `json:"website_domain"`
	JumpURL       string `json:"jump_url"`
	Greeting      string `json:"greeting"`
	Intro         string `json:"intro"`
	Outro         string `json:"outro"`
	Year          int    `json:"year"`
}

func (w *Website) NotifyMailData(jumpURL string, greeting string, intro string, outro string) (subject string, body string) {
	mailData := &NotifyMailData{
		WebsiteURL:    w.WebsiteURL,
		WebsiteName:   w.WebsiteName,
		WebsiteDomain: w.WebsiteDomain,
		JumpURL:       jumpURL,
		Greeting:      greeting,
		Intro:         intro,
		Outro:         outro,
		Year:          time.Now().Year(),
	}
	mailTplContent := getEmailHTMLContent("pkg/email/templates/notify-mail.html", mailData)
	return "告警", mailTplContent
}

// getEmailHTMLContent 获取邮件模板
func getEmailHTMLContent(tplPath string, mailData interface{}) string {
	b, err := os.ReadFile(tplPath)
	if err != nil {
		fmt.Printf("[util.email] read file err: %v", err)
		return ""
	}
	mailTpl := string(b)
	tpl, err := template.New("email tpl").Parse(mailTpl)
	if err != nil {
		fmt.Printf("[util.email] template new err: %v", err)
		return ""
	}
	buffer := new(bytes.Buffer)
	err = tpl.Execute(buffer, mailData)
	if err != nil {
		fmt.Printf("[util.email] execute template err: %v", err)
	}
	return buffer.String()
}
