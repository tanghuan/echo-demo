package utils

import (
	"regexp"
	"testing"
)

// TestSayHelloName 测试 name 有效时
func TestSayHelloName(t *testing.T) {

	name := "TangHuan"

	want := regexp.MustCompile(`\b` + name + `\b`)

	msg, err := SayHello(name)
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("TangHuan") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// TestSayHelloEmpty 测试 name 为空字符串时
func TestSayHelloEmpty(t *testing.T) {
	msg, err := SayHello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}

// TestSayHelloSpaceName 测试 name 为空白字符串时
func TestSayHelloSpaceName(t *testing.T) {
	msg, err := SayHello("  ")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("  ") = %q, %v, want "", error`, msg, err)
	}
}
