package log

import "testing"

func TestGetLogger(t *testing.T) {
	log := GetLogger()
	log.Info("this is a info message")
	log.Error("this is a error message")
	log.Debug("this is a debug message")
}
