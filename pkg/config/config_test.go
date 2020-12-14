package config

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestViperEnvVariable(t *testing.T) {
	r := viperEnvVariable("SEPTA_URL")
	fmt.Println(r)
	fmt.Println()
}

func TestSetEnv(t *testing.T) {
	SetEnv()
}

func TestIfEnvNot(t *testing.T) {
	os.Setenv("k", "stuff")
	IfEnvNot("k")
}

func TestCheckENV(t *testing.T) {
	SetEnv()

	for _, j := range tokens {
		r := os.Getenv(j)
		if r == "" {
			t.Fatalf("Env: %v not set\n", j)
		}
	}
	if !strings.Contains(os.Getenv("SEPTA_URL"), "cwxstat") {
		t.Fatalf("\n\nWe're not getting values\n\n")
	}

	if !strings.Contains(os.Getenv("SEPTA_TOKEN"), "TzlRuS4yeN") {
		t.Fatalf("\n\nWe're token Contains\n\n")
	}

	if !strings.HasPrefix(os.Getenv("SEPTA_TOKEN"), "TzlRuS4yeN") {
		t.Fatalf("\n\nWe're token Prefix\n\n")
	}

	if !strings.HasSuffix(os.Getenv("SEPTA_TOKEN"), "ewEuddmQ==") {
		t.Fatalf("\n\nWe're not getting token Suffix\n\n")
	}

}
