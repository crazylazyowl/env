package env

import (
	"os"
	"testing"
)

const (
	strKey  = "KEY_STRING"
	intKey  = "KEY_INT"
	boolKey = "KEY_BOOL"
)

var environments = map[string]string{
	strKey:  "magic",
	intKey:  "31337",
	boolKey: "true",
}

func TestStr(t *testing.T) {
	setEnvs()
	defer unsetEnvs()
	type args struct {
		key     string
		choices []string
	}
	tests := []struct {
		args      args
		wantValue string
		wantErr   bool
	}{
		{args{strKey, []string{}}, "magic", false},
		{args{"some_key", []string{}}, "", true},
		{args{strKey, []string{"magic"}}, "magic", false},
		{args{strKey, []string{"some_value"}}, "", true},
	}
	for _, test := range tests {
		value, err := Str(test.args.key, test.args.choices...)
		if err != nil {
			if !test.wantErr {
				t.Fatalf("Str(%s) error = '%v', wantErr = %v", test.args.key, err, test.wantErr)
			}
			continue
		}
		if value != test.wantValue {
			t.Fatalf("Str(%s) = %v, wantValue = %v", test.args.key, value, test.wantValue)
		}
	}
}

func TestInt(t *testing.T) {
	setEnvs()
	defer unsetEnvs()
	type args struct {
		key     string
		choices []int
	}
	tests := []struct {
		args      args
		wantValue int
		wantErr   bool
	}{
		{args{intKey, []int{}}, 31337, false},
		{args{"some_key", []int{}}, 0, true},
		{args{intKey, []int{31337}}, 31337, false},
		{args{intKey, []int{0}}, 0, true},
	}
	for _, test := range tests {
		value, err := Int(test.args.key, test.args.choices...)
		if err != nil {
			if !test.wantErr {
				t.Fatalf("Int(%s) error = '%v', wantErr = %v", test.args.key, err, test.wantErr)
			}
			continue
		}
		if value != test.wantValue {
			t.Fatalf("Int(%s) = %v, wantValue = %v", test.args.key, value, test.wantValue)
		}
	}
}

func TestBool(t *testing.T) {
	setEnvs()
	defer unsetEnvs()
	type args struct {
		key string
	}
	tests := []struct {
		args      args
		wantValue bool
		wantErr   bool
	}{
		{args{boolKey}, true, false},
		{args{"some_key"}, false, true},
	}
	for _, test := range tests {
		value, err := Bool(test.args.key)
		if err != nil {
			if !test.wantErr {
				t.Fatalf("Bool(%s) error = '%v', wantErr = %v", test.args.key, err, test.wantErr)
			}
			continue
		}
		if value != test.wantValue {
			t.Fatalf("Bool(%s) = %v, wantValue = %v", test.args.key, value, test.wantValue)
		}
	}
}

func setEnvs() {
	for key, value := range environments {
		os.Setenv(key, value)
	}
}

func unsetEnvs() {
	for key := range environments {
		os.Unsetenv(key)
	}
}
