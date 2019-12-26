package helper

import "testing"

func TestGenerateSHA512(t *testing.T) {
	funcName := "GenerateSHA512()"

	tests := []struct {
		Name string
		Text string
		Want string
	}{
		{
			Name: "Empty string",
			Text: "",
			Want: "cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e",
		},
		{
			Name: "string is sample",
			Text: "sample",
			Want: "39a5e04aaff7455d9850c605364f514c11324ce64016960d23d5dc57d3ffd8f49a739468ab8049bf18eef820cdb1ad6c9015f838556bc7fad4138b23fdf986c7",
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			if got := GenerateSHA512(tt.Text); got != tt.Want {
				t.Errorf("%v = %v, want %v", funcName, got, tt.Want)
			}
		})
	}
}
