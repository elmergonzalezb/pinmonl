package tag

import (
	"testing"

	"github.com/pinmonl/pinmonl/model"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		name  string
		in    *Input
		wants bool
	}{
		{
			name:  "name",
			in:    &Input{Tag: model.Tag{Name: "name.1093-04__(abc) : [abc]"}},
			wants: true,
		},
		{
			name:  "invalid name 01",
			in:    &Input{Tag: model.Tag{Name: "name="}},
			wants: false,
		},
		{
			name:  "invalid name 02",
			in:    &Input{Tag: model.Tag{Name: "name!"}},
			wants: false,
		},
		{
			name:  "invalid name 03",
			in:    &Input{Tag: model.Tag{Name: "name?"}},
			wants: false,
		},
	}

	for _, test := range tests {
		err := test.in.Validate()
		got := err == nil
		if got != test.wants {
			t.Errorf("case %q fails", test.name)
		}
	}
}
