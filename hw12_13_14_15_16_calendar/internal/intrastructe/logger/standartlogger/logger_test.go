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

	logger := New("info")
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
