package logger

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type level string

const (
	INFO  level = "INFO"
	ERROR level = "ERROR"
)

type LogMsg struct {
	Level level
	Msg   string
	Time  time.Time
}

type Logger struct {
	logChan  chan LogMsg
	wg       sync.WaitGroup
	file     *os.File
	stopChan chan struct{}
}

func NewLogger(filePath string, bufferSize int) (*Logger, error) {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	log := &Logger{
		logChan:  make(chan LogMsg, bufferSize),
		file:     file,
		stopChan: make(chan struct{}),
	}

	log.wg.Add(1)
	go log.loggerProcess()

	return log, nil

}

func (l *Logger) loggerProcess() {
	defer l.wg.Done()
	for {
		select {
		case msg, ok := <-l.logChan:
			if !ok {
				return
			}
			l.writerLog(msg)
		case <-l.stopChan:
			close(l.logChan)

			for msg := range l.logChan {
				l.writerLog(msg)
			}
			return
		}
	}
}

func (l *Logger) writerLog(msg LogMsg) {
	log := fmt.Sprintf("%s [%s] %s\n", msg.Time.Format(time.RFC3339), msg.Level, msg.Msg)
	if _, err := l.file.WriteString(log); err != nil {
		fmt.Printf("Failed to write log: %v", err)
	}

}

func (l *Logger) Info(msg string) {
	l.logChan <- LogMsg{Level: INFO, Msg: msg, Time: time.Now()}
}

func (l *Logger) Error(msg string) {
	l.logChan <- LogMsg{Level: ERROR, Msg: msg, Time: time.Now()}
}

func (l *Logger) Close() {
	close(l.stopChan)
	l.wg.Wait()
	l.file.Close()
}
