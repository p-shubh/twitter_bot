package bot

import (
	"github.com/rajatjindal/twitter-bot/twitter"
	"github.com/sirupsen/logrus"
)

const (
	webhookHost = "https://your-webhook-domain"
	webhookPath = "/webhook/twitter"
	//has to be same as provided when getting tokens from twitter developers console
	environmentName = "prod"
)

func BotFunction() *twitter.Bot {
	// use the tokens saved when following above steps
	bot, err := twitter.NewBot(
		&twitter.BotConfig{
			Tokens: twitter.Tokens{
				ConsumerKey:   "<consumer-key>",
				ConsumerToken: "<consumer-token>",
				Token:         "<token>",
				TokenSecret:   "<token-secret>",
			},
			Environment:          environmentName,
			WebhookHost:          webhookHost,
			WebhookPath:          webhookPath,
			OverrideRegistration: true,
		},
	)
	if err != nil {
		logrus.Fatal(err)
	}

	return bot

}
