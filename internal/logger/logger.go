package logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type LokiLogger struct {
	batch       int64
	serviceName string
	LokiPayload []LokiStream
	path        string
}

type LokiStream struct {
	Stream map[string]string `json:"stream"`
	Values [][]string        `json:"values"`
}
type LokiStreams struct {
	Streams []LokiStream `json:"streams"`
}

func NewLokiLogger(serviceName string, batch int64, path string) *LokiLogger {
	return &LokiLogger{
		batch:       batch,
		serviceName: serviceName,
		LokiPayload: []LokiStream{},
		path:        path,
	}
}

func (l *LokiLogger) Log(message, level string) {
	l.LokiPayload = append(l.LokiPayload, LokiStream{
		Stream: map[string]string{
			"job":   l.serviceName,
			"level": level,
		},
		Values: [][]string{
			{strconv.FormatInt(time.Now().UnixNano(), 10), message},
		},
	})
	if len(l.LokiPayload) >= int(l.batch) {
		l.Send()
		l.LokiPayload = []LokiStream{}
	}
}
func (l *LokiLogger) Send() {
	jsonData, err := json.Marshal(LokiStreams{Streams: l.LokiPayload})
	if err != nil {
		log.Fatalf("Ошибка сериализации JSON: %v", err)
	}
	fmt.Println(jsonData)
	// Отправляем POST-запрос в Loki
	resp, err := http.Post(l.path+"/loki/api/v1/push", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Ошибка отправки запроса: %v", err)
	}
	defer resp.Body.Close()

}
