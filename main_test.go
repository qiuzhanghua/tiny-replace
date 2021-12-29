package main

import (
	"os"
	"testing"
)

func TestReplaceString(t *testing.T) {
	content := "Hello, ${JAVA_HOME}"
	_ = os.Setenv("JAVA_HOME", "/usr/bin")
	actual := ReplaceString(content)
	expected := "Hello, /usr/bin"
	if actual != expected {
		t.Errorf("Expected : %v, actual is '%v'", expected, actual)
	}
}

func TestReplaceString2(t *testing.T) {
	content := "Hello, %JAVA_HOME%"
	_ = os.Setenv("JAVA_HOME", "/usr/bin")
	actual := ReplaceString(content)
	expected := "Hello, /usr/bin"
	if actual != expected {
		t.Errorf("Expected : %v, actual is '%v'", expected, actual)
	}

}
