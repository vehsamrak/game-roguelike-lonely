package logger

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
)

type TextFormatter struct {
}

func (formatter *TextFormatter) Format(entry *log.Entry) ([]byte, error) {
	message := fmt.Sprintf(
		"%s [%s] %s",
		entry.Time.Format("2006-01-02 15:04:05"),
		entry.Level.String(),
		entry.Message,
	)

	if len(entry.Data) > 0 {
		serializedData, _ := json.Marshal(entry.Data)
		message = fmt.Sprintf("%s | %s", message, serializedData)
	}

	return []byte(message + "\n"), nil
}
