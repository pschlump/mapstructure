package mapstructure

import "testing"

type Basic struct {
	Vstring string
	Vint    int
	Vbool   bool
	Vextra  string
}

type Nested struct {
	Vfoo string
	Vbar Basic
}

func TestBasicTypes(t *testing.T) {
	t.Parallel()

	input := map[string]interface{}{
		"vstring": "foo",
		"vint":    42,
		"vbool":   true,
	}

	var result Basic
	err := MapToStruct(input, &result)
	if err != nil {
		t.Errorf("got an err: %s", err.Error())
		t.FailNow()
	}

	if result.Vstring != "foo" {
		t.Errorf("vstring value should be 'foo': %#v", result.Vstring)
	}

	if result.Vint != 42 {
		t.Errorf("vint value should be 42: %#v", result.Vint)
	}

	if result.Vbool != true {
		t.Errorf("vbool value should be true: %#v", result.Vbool)
	}

	if result.Vextra != "" {
		t.Errorf("vextra value should be empty: %#v", result.Vextra)
	}
}

func TestNestedType(t *testing.T) {
	t.Parallel()

	input := map[string]interface{}{
		"vfoo": "foo",
		"vbar": map[string]interface{}{
			"vstring": "foo",
			"vint":    42,
			"vbool":   true,
		},
	}

	var result Nested
	err := MapToStruct(input, &result)
	if err != nil {
		t.Errorf("got an err: %s", err.Error())
		t.FailNow()
	}

	if result.Vfoo != "foo" {
		t.Errorf("vfoo value should be 'foo': %#v", result.Vfoo)
	}

	if result.Vbar.Vstring != "foo" {
		t.Errorf("vstring value should be 'foo': %#v", result.Vbar.Vstring)
	}

	if result.Vbar.Vint != 42 {
		t.Errorf("vint value should be 42: %#v", result.Vbar.Vint)
	}

	if result.Vbar.Vbool != true {
		t.Errorf("vbool value should be true: %#v", result.Vbar.Vbool)
	}

	if result.Vbar.Vextra != "" {
		t.Errorf("vextra value should be empty: %#v", result.Vbar.Vextra)
	}
}

func TestInvalidType(t *testing.T) {
	t.Parallel()

	input := map[string]interface{}{
		"vstring": 42,
	}

	var result Basic
	err := MapToStruct(input, &result)
	if err == nil {
		t.Error("error should exist")
		t.FailNow()
	}

	if err.Error() != "'root.Vstring' expected type 'string', got 'int'" {
		t.Errorf("got unexpected error: %s", err)
	}
}

func TestNonPtrValue(t *testing.T) {
	t.Parallel()

	err := MapToStruct(map[string]interface{}{}, Basic{})
	if err == nil {
		t.Error("error should exist")
		t.FailNow()
	}

	if err.Error() != "val must be a pointer" {
		t.Errorf("got unexpected error: %s", err)
	}
}

func TestNontStructValue(t *testing.T) {
	t.Parallel()

	result := 42
	err := MapToStruct(map[string]interface{}{}, &result)
	if err == nil {
		t.Error("error should exist")
		t.FailNow()
	}

	if err.Error() != "val must be an addressable struct" {
		t.Errorf("got unexpected error: %s", err)
	}
}