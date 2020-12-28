package slack

import (
	slack2 "github.com/balerter/balerter/internal/config/channels/slack"
	"github.com/nlopes/slack"
	"go.uber.org/zap"
)

type API interface {
	SendMessage(channel string, options ...slack.MsgOption) (string, string, string, error)
}

type Slack struct {
	logger  *zap.Logger
	name    string
	channel string
	api     API
}

func New(cfg *slack2.Slack, logger *zap.Logger) (*Slack, error) {
	m := &Slack{
		logger:  logger,
		name:    cfg.Name,
		channel: cfg.Channel,
	}

	m.api = slack.New(cfg.Token)

	return m, nil
}

func (m *Slack) Name() string {
	return m.name
}