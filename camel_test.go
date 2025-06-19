// Copyright (c) 2017, A. Stoewer <adrian.stoewer@rz.ifi.lmu.de>
// All rights reserved.

package strcase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpperCamelCase(t *testing.T) {
	data := map[string]string{
		"":                      "",
		"f":                     "F",
		"foo":                   "Foo",
		"fooBar":                "FooBar",
		"FooBarBla":             "FooBarBla",
		"foo_barBla":            "FooBarBla",
		" foo_bar\n":            "FooBar",
		" foo-bar\t":            "FooBar",
		" foo bar\r":            "FooBar",
		"HTTP_status_code":      "HttpStatusCode",
		"skip   many spaces":    "SkipManySpaces",
		"skip---many-dashes":    "SkipManyDashes",
		"skip___many_underline": "SkipManyUnderline",
		"XRequestId":            "XRequestId",
		"HTTPStatusCode":        "HttpStatusCode",
		"cfg2_test":             "Cfg2Test",
		"Cfg2Test":              "Cfg2Test",
	}

	for in, out := range data {
		converted := UpperCamelCase(in)
		assert.Equal(t, out, converted)
	}
}

func TestLowerCamelCase(t *testing.T) {
	data := map[string]string{
		"":                      "",
		"F":                     "f",
		"foo":                   "foo",
		"FooBar":                "fooBar",
		"fooBarBla":             "fooBarBla",
		"foo_barBla":            "fooBarBla",
		" foo_bar\n":            "fooBar",
		" foo-bar\t":            "fooBar",
		" foo bar\r":            "fooBar",
		"HTTP_status_code":      "httpStatusCode",
		"skip   many spaces":    "skipManySpaces",
		"skip---many-dashes":    "skipManyDashes",
		"skip___many_underline": "skipManyUnderline",
		"XRequestId":            "xRequestId",
		"HTTPStatusCode":        "httpStatusCode",
		"cfg2_test":             "cfg2Test",
		"cfg2Test":              "cfg2Test",
	}

	for in, out := range data {
		converted := LowerCamelCase(in)
		assert.Equal(t, out, converted)
	}
}

// TestMultipleCamelCaseApplications tests that applying the camelCase functions
// multiple times to strings with numbers preserves the correct capitalization.
func TestMultipleCamelCaseApplications(t *testing.T) {
	// Test LowerCamelCase
	input := "cfg2_test"
	expected := "cfg2Test"

	// First application
	result := LowerCamelCase(input)
	assert.Equal(t, expected, result)

	// Second application
	result = LowerCamelCase(result)
	assert.Equal(t, expected, result)

	// Test UpperCamelCase
	input = "cfg2_test"
	expected = "Cfg2Test"

	// First application
	result = UpperCamelCase(input)
	assert.Equal(t, expected, result)

	// Second application
	result = UpperCamelCase(result)
	assert.Equal(t, expected, result)
}
