package ulog

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/evalphobia/logrus_fluent"
)

type Fluent struct {
	Host   string
	Port   int
	Tag    string
	Levels []string
}

func newHookFluent(fluent Fluent) (*logrus_fluent.FluentHook, error) {
	hook := logrus_fluent.NewHook(fluent.Host, fluent.Port)
	levels, err := getLogLevels(fluent.Levels)
	if err != nil {
		return nil, err
	}
	if len(fluent.Tag) == 0 {
		return nil, fmt.Errorf("fluent: unknown tag value")
	}
	hook.SetTag(fluent.Tag)
	hook.SetLevels(levels)
	return hook, nil
}

func getLogLevels(levels []string) ([]logrus.Level, error) {
	rets := []logrus.Level{}
	for _, level := range levels {
		ret, err := logrus.ParseLevel(level)
		if err != nil {
			return nil, err
		}
		rets = append(rets, ret)
	}
	return rets, nil
}
