package main

import (
	"testing"
)

func Test_plusPlus(t *testing.T) {
	type args struct {
		a int
		b int
		c int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"name1",
			args{
				a: 1, b: 2, c: 3,
			},
			7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := plusPlus(tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("plusPlus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_plus(t *testing.T) {
	r := plus(1, 2)
	t.Logf("result = %v", r)
}

func Test_add(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"one",
			args{
				nums: []int{2, 3},
			},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := add(tt.args.nums...); got != tt.want {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 模糊测试
func FuzzAdd(f *testing.F) {
	testcases := []int{2, 5}
	for _, tc := range testcases {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}
	//第二个参数只能是
	/*
		string, []byte
		int, int8, int16, int32/rune, int64
		uint, uint8/byte, uint16, uint32, uint64
		float32, float64
		bool
	*/
	f.Fuzz(func(t *testing.T, i int) {
		r := add(i)
		t.Errorf("r= %d", r)
	})
}
