package content

import (
	"testing"
)

type test struct {
	data        interface{}
	expectError bool
}

var tests = []test{
	{File{}, true},
	{&File{}, true},
	{&File{Contents: "Something"}, true},
	{&File{Name: "Something"}, false},
	{&File{Name: "Something", Contents: "Something"}, false},
	{&SourceLocation{}, true},
	{&SourceLocation{Line: 1, Column: 1}, true},
	{&SourceLocation{File: File{Name: "Something"}, Line: 1, Column: 1}, false},
	{&FullyQualifiedName{}, true},
	{&FullyQualifiedName{Absolute: "Something"}, true},
	{&FullyQualifiedName{Relative: "Something"}, false},
	{&FullyQualifiedName{"Something.Something", "Something.Something"}, true},
	{&FullyQualifiedName{"Something", "Something.Something"}, false},
	{&FullyQualifiedName{".Something", "Something.Something"}, true},
	{&FullyQualifiedName{"\\.Something", "Something.Something"}, false},
	{&FullyQualifiedName{"\\.$omething", "Something.Something"}, true},
	{&Type{}, true},
	{&Type{Flags: FLAG_TYPE_ARRAY}, true},
	{&Type{Flags: FLAG_TYPE_ARRAY, Specialization: []Type{Type{Name: FullyQualifiedName{Absolute: "Test"}}}}, false},
}

func TestValidate(t *testing.T) {
	for i, test := range tests {
		err := Validate(test.data)
		if err == nil && test.expectError {
			t.Errorf("Expected an error in test %d", i)
		} else if err != nil && !test.expectError {
			t.Errorf("Didn't expected an error: %s", err)
		}
	}
}
