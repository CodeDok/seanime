package mkvparser

import (
	"strings"
	"testing"
)

func TestConvertToASS_SRT(t *testing.T) {
	srt := `1
00:00:01,000 --> 00:00:02,000
Hello world
`
	out, err := ConvertToASS(srt, SubtitleTypeSRT)
	if err != nil {
		t.Fatalf("valid SRT should convert: %v", err)
	}
	if !strings.Contains(out, "Hello world") {
		t.Fatal("converted ASS is missing the dialogue text")
	}
}

func TestConvertToASS_NoCues(t *testing.T) {
	// Empty input and non-subtitle input (e.g. an HTML page fetched by mistake)
	// parse to zero cues without a read error; conversion must fail clearly.
	for name, content := range map[string]string{
		"empty": "",
		"html":  "<!DOCTYPE html><html><body>Login</body></html>",
	} {
		if _, err := ConvertToASS(content, SubtitleTypeSRT); err == nil || !strings.Contains(err.Error(), "no cues") {
			t.Errorf("%s: expected 'no cues' error from ConvertToASS, got: %v", name, err)
		}
	}
	if _, err := ConvertToVTT("", SubtitleTypeSRT); err == nil || !strings.Contains(err.Error(), "no cues") {
		t.Errorf("expected 'no cues' error from ConvertToVTT, got: %v", err)
	}
}
