package stack

import "testing"

func TestIsEmpty(t *testing.T) {
	test := Stack{}
	if !test.IsEmpty() {
		t.Errorf("%#v expected true but received false", test)
	}
}

func TestSize(t *testing.T) {
	sizeTests := []struct {
		name  string
		stack Stack
		want  int
	}{
		{name: "One long", stack: Stack{1}, want: 1},
		{name: "Three long", stack: Stack{1, "hello", true}, want: 3},
	}

	for _, tt := range sizeTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.stack.Size()
			if got != tt.want {
				t.Errorf("%#v expected %v but got %v", tt.stack, tt.want, got)
			}
		})
	}
}

func TestPush(t *testing.T) {
	pushTests := []struct {
		name       string
		stack      Stack
		currLength int
	}{
		{name: "One long", stack: Stack{1}, currLength: 1},
		{name: "Three long", stack: Stack{1, "hello", true}, currLength: 3},
	}

	for _, tt := range pushTests {
		t.Run(tt.name, func(t *testing.T) {
			tt.stack.Push("new")
			if tt.stack.Size() != tt.currLength+1 {
				t.Errorf("Size of %#v should have increased by 1 but stayed %v", tt.stack, tt.currLength)
			}

		})
	}

}

func TestPop(t *testing.T) {
	test := Stack{1, 2, 3}
	currLength := test.Size()
	want := 3
	got := test.Pop()
	if got != want {
		t.Errorf("%#v expected %v but got %v", test, want, got)
	}
	if test.Size() == currLength {
		t.Errorf("Size of %#v should have decreased by 1 but stayed %v", test, currLength)
	}
}

func TestPeek(t *testing.T) {
	test := Stack{1, 2, 3}
	currLength := test.Size()
	want := 3
	got := test.Peek()
	if got != want {
		t.Errorf("%#v expected %v but got %v", test, want, got)
	}
	if test.Size() != currLength {
		t.Errorf("Size of %#v should have stayed %v but decreased by 1", test, currLength)
	}
}
