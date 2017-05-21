package zaplog

import (
	"log"
	"testing"
)

func TestFormatStdLog(t *testing.T) {
	FormatStdLog()
	log.Print("std log")
	log.Print("")
}

func TestNewLogger(t *testing.T) {
	logger := NewLogger(true)
	defer logger.Sync()
	logger.Debug("zap log debug msg")
}

func TestNewCustomLoggers(t *testing.T) {
	logger, noCallerLogger := NewCustomLoggers(true)
	logger.Info("normalLogger")
	noCallerLogger.Info("noCallerLogger")
}

func TestNewCompatibleLogger(t *testing.T) {
	compatibleLogger := NewCompatibleLogger(true)
	compatibleLogger.Print("compatibleLogger Info")
	compatibleLogger.Printf("compatibleLogger Infof :%v", 1)
	compatibleLogger.WithField("field", "value").Info("compatibleLogger with info")
	withFieldLogger := compatibleLogger.WithFields(map[string]interface{}{"field1": "value1", "field2": "value2"})
	withFieldLogger.Info("withFieldLogger Info")
	withFieldLogger.With("field3", "value3").Info("with filed3")
	withFieldLogger.Debugf("withFieldLogger debugf:%v", 1)
}

func BenchmarkStdLogger(b *testing.B) {
	for i := 0; i < b.N; i++ {
		log.Print("std log printf")
	}
}

func BenchmarkNewLogger(b *testing.B) {
	logger := NewLogger(true)
	defer logger.Sync()
	for i := 0; i < b.N; i++ {
		logger.Debug("zap log debug msg")
	}
}
