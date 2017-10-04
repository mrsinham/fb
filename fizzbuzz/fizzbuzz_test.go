package fizzbuzz

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	type args struct {
		string1 string
		string2 string
		int1    int
		int2    int
		limit   int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "fizzbuzz classic",
			args: args{
				int1:    3,
				int2:    5,
				limit:   20,
				string1: "fizz",
				string2: "buzz",
			},
			want: []string{
				"1",
				"2",
				"fizz",
				"4",
				"buzz",
				"fizz",
				"7",
				"8",
				"fizz",
				"buzz",
				"11",
				"fizz",
				"13",
				"14",
				"fizzbuzz",
				"16",
				"17",
				"fizz",
				"19",
				"buzz",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := Naive(context.Background(), tt.args.string1, tt.args.string2, tt.args.int1, tt.args.int2, tt.args.limit); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fizzBuzzNaive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fizzBuzzSwitch(t *testing.T) {
	type args struct {
		string1 string
		string2 string
		int1    int
		int2    int
		limit   int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "fizzbuzz classic",
			args: args{
				int1:    3,
				int2:    5,
				limit:   20,
				string1: "fizz",
				string2: "buzz",
			},
			want: []string{
				"1",
				"2",
				"fizz",
				"4",
				"buzz",
				"fizz",
				"7",
				"8",
				"fizz",
				"buzz",
				"11",
				"fizz",
				"13",
				"14",
				"fizzbuzz",
				"16",
				"17",
				"fizz",
				"19",
				"buzz",
			},
		},
		{
			name: "limit 0",
			args: args{
				int1:    3,
				int2:    5,
				limit:   0,
				string1: "fizz",
				string2: "buzz",
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Switch(tt.args.string1, tt.args.string2, tt.args.int1, tt.args.int2, tt.args.limit); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fizzBuzzSwitch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkFizzBuzzNaive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Naive(context.Background(), "fizz", "buzz", 3, 5, 20)
	}
}

func BenchmarkFizzBuzzSwitch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Switch("fizz", "buzz", 3, 5, 20)
	}
}

func ExampleNaive() {
	res, _ := Naive(context.Background(), "fizz", "buzz", 3, 5, 20)
	fmt.Println(res)
	// Output:[1 2 fizz 4 buzz fizz 7 8 fizz buzz 11 fizz 13 14 fizzbuzz 16 17 fizz 19 buzz]
}
