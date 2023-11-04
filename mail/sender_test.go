package mail

import (
	"path"
	"testing"

	"github.com/fauzanfebrian/simplebank/util"
	"github.com/stretchr/testify/require"
)

func TestSendEmailWithGmail(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	envPath := path.Join(util.GetProjectPath(), ".env")
	config, err := util.LoadConfig(envPath)
	require.NoError(t, err)

	sender := NewEmailSender(config)

	subject := "A test email"
	content := `
	<h1>Hello world</h1>
	<p>This is a test message from SimpleBank</p>
	`
	to := []string{"test@simplebank.com"}

	err = sender.SendEmail(subject, content, to, nil, nil, nil)
	require.NoError(t, err)
}
