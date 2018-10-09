package birthday

import (
	"testing"
)

func TestNewBirthday(t *testing.T) {
	for i := 0; i < 100000; i++ {

		a := NewBirthday()
		if !(1 <= a && a <= 365) {
			t.Fatalf("誕生日が1~365じゃないよ。a=%v", a)
		}
	}
}

func TestCheckSameBirthday(t *testing.T) {

	tests := []struct {
		name  string
		class []int
		want  bool
	}{
		{"同じ誕生日がいるパターン", []int{1, 2, 3, 4, 4, 5}, true},
		{"同じ誕生日が居ないパターン", []int{23, 55, 1, 3, 77}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckSameBirthday(tt.class); got != tt.want {
				t.Errorf("CheckSameBirthday() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckSameBirthday2(t *testing.T) {

	tests := []struct {
		name  string
		class []int
		want  bool
	}{
		{"同じ誕生日がいるパターン", []int{1, 2, 3, 4, 4, 5}, true},
		{"同じ誕生日が居ないパターン", []int{23, 55, 1, 3, 77}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckSameBirthday2(tt.class); got != tt.want {
				t.Errorf("CheckSameBirthday2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckSameBirthday3(t *testing.T) {

	tests := []struct {
		name  string
		class []int
		want  bool
	}{
		{"同じ誕生日がいるパターン", []int{1, 2, 3, 4, 4, 5}, true},
		{"同じ誕生日が居ないパターン", []int{23, 55, 1, 3, 77}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckSameBirthday3(tt.class); got != tt.want {
				t.Errorf("CheckSameBirthday2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkCheckSameBirthday(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c := NewClass(100000)
		CheckSameBirthday(c)
	}
}

func BenchmarkCheckSameBirthday2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c := NewClass(100000)
		CheckSameBirthday2(c)
	}
}

func BenchmarkCheckSameBirthday3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c := NewClass(100000)
		CheckSameBirthday3(c)
	}
}
