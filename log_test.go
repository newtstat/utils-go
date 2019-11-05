package utils

import (
	"log"
	"testing"
)

var testLog = Log

func TestLogT_Printfln(t *testing.T) {
	testLog.Printfln("test: %s", "TestLogT_Printfln()")
}

func TestLogT_Fatalfln(t *testing.T) {
	testLog.fatallnFunc = log.Println
	testLog.Fatalfln("test: %s", "TestLogT_Fatalfln()")
}
