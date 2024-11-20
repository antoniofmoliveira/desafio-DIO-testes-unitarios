package main

import (
	"testing"
)

func TestSoma(t *testing.T) {
	type args struct {
		i []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Soma",
			args: args{i: []int{1, 2, 3}},
			want: 6,
		},

		{
			name: "Soma",
			args: args{i: []int{1, 2, 3, 4, 5}},
			want: 15,
		},
		{
			name: "Soma",
			args: args{i: []int{}},
			want: 0,
		},
		{
			name: "Soma",
			args: args{i: []int{1}},
			want: 1,
		},
		{
			name: "Soma",
			args: args{i: []int{1, -2}},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Soma(tt.args.i...); got != tt.want {
				t.Errorf("Soma() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiplica(t *testing.T) {
	type args struct {
		i []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Multiplica",
			args: args{i: []int{1, 2, 3}},
			want: 6,
		},
		{
			name: "Multiplica",
			args: args{i: []int{1, 2, 3, 4, 5}},
			want: 120,
		},
		{
			name: "Multiplica",
			args: args{i: []int{}},
			want: 1,
		},
		{
			name: "Multiplica",
			args: args{i: []int{1}},
			want: 1,
		},
		{
			name: "Multiplica",
			args: args{i: []int{1, 0}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Multiplica(tt.args.i...); got != tt.want {
				t.Errorf("Multiplica() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	type args struct {
		i []float64
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "Divide",
			args: args{i: []float64{10, 2}},
			want: 5,
		},
		{
			name:    "Divide",
			args:    args{i: []float64{10, 0}},
			wantErr: true,
		},
		{
			name: "Divide",
			args: args{i: []float64{10, 2, 2}},
			want: 2.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Divide(tt.args.i...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Divide() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Divide() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubtrai(t *testing.T) {
	type args struct {
		i []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Subtrai",
			args: args{i: []int{1, 2, 3}},
			want: -4,
		},
		{
			name: "Subtrai",
			args: args{i: []int{1, 2, 3, 4, 5}},
			want: -13,
		},
		{
			name: "Subtrai",
			args: args{i: []int{}},
			want: 0,
		},
		{
			name: "Subtrai",
			args: args{i: []int{1}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Subtrai(tt.args.i...); got != tt.want {
				t.Errorf("Subtrai() = %v, want %v", got, tt.want)
			}
		})
	}
}
