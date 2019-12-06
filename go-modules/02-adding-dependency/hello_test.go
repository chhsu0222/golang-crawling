package hello

import "testing"

func TestHello(t *testing.T) {
	want := "你好，世界。"
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
