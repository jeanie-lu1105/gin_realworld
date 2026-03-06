package logger

import (
	"context"

	"github.com/sirupsen/logrus"
)


func New(ctx context.Context) *logrus.Entry{
	return logrus.New().WithContext(ctx)
}