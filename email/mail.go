package email

// Driver 邮件发送驱动接口定义
type Driver interface {
	// Send 发送邮件
	Send(to, subject, body string) error
	// Close 关闭链接
	Close()
}

// Send 发送邮件
func Send(to, subject, body string) error {
	Lock.RLock()
	defer Lock.RUnlock()

	if Client == nil {
		return nil
	}

	return Client.Send(to, subject, body)
}

// EmailConfig
// @Description: 邮件配置
type EmailConfig struct {
	Host      string `json:"host"`
	Port      int    `json:"port"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	ReplyTo   string `json:"reply_to"`
	KeepAlive int    `json:"keep_alive"`
}
