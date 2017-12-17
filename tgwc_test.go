package tgwc

import (
	"reflect"
	"testing"
)

func TestGenerate(t *testing.T) {
	type args struct {
		word       string
		dictionary []string
	}
	tests := []struct {
		name    string
		args    args
		want    *Puzzle
		wantErr bool
	}{
		{
			name: "Generates solution",
			args: args{
				word: "abc",
				dictionary: []string{
					"cre",
					"drc",
					"creadrc",
					"crebdrc",
					"crecdrc",
				},
			},
			want: &(Puzzle{
				Letter{
					'a',
					[]Row{
						Row{"cre", "drc"},
					},
				},
				Letter{
					'b',
					[]Row{
						Row{"cre", "drc"},
					},
				},
				Letter{
					'c',
					[]Row{
						Row{"cre", "drc"},
					},
				},
			}),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Generate(tt.args.word, tt.args.dictionary)
			if (err != nil) != tt.wantErr {
				t.Errorf("Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Generate() = %v, want %v", got, tt.want)
			}
		})
	}
}
