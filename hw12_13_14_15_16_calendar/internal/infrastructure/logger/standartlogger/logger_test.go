package standartlogger

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestLogger_Info_Output(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	logger, err := New("info", "http_test.log")
	if err != nil {
		t.Errorf("Ошибка создания логгера: %v", err)
	}

	logger.Info("Привет, лог!")
	w.Close()
	var buf bytes.Buffer
	io.Copy(&buf, r)

	os.Stdout = old

	output := buf.String()

	if !strings.Contains(output, "[INFO] Привет, лог!") {
		t.Errorf("Ожидали '[INFO] Привет, лог!', получили: %s", output)
	}
}
