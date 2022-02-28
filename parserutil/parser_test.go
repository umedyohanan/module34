package parserutil

import "testing"

func TestParseString(t *testing.T) {
	type args struct {
		inputString string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		want1   int
		want2   string
		wantErr bool
	}{
		{
			name: "Sum",
			args: args{
				inputString: "5+4=?",
			},
			want:    5,
			want1:   4,
			want2:   "+",
			wantErr: false,
		},
		{
			name: "String",
			args: args{
				inputString: "It is good weather",
			},
			want:    0,
			want1:   0,
			want2:   "",
			wantErr: true,
		},
		{
			name: "First operand is string literal",
			args: args{
				inputString: "a+4=?",
			},
			want:    0,
			want1:   0,
			want2:   "",
			wantErr: true,
		},
		{
			name: "Second operand is string literal",
			args: args{
				inputString: "5+a=?",
			},
			want:    0,
			want1:   0,
			want2:   "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, err := ParseString(tt.args.inputString)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseString() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ParseString() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("ParseString() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
