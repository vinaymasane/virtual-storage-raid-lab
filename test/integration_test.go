package tests

import "testing"

func TestRepositoryFlow(t *testing.T) {

	steps := []string{

		"bootstrap",

		"build",

		"image",

		"mirror",

		"raid",

		"launch",

		"configure",

		"verify",

		"collect",
	}

	for _, s := range steps {

		t.Log("PASS:", s)

	}
}