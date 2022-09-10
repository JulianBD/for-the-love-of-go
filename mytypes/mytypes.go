package mytypes

import "strings"

// Twice multiplies its receiver by 2 and returns
// the result.

type MyInt int
type MyString string
type MyBuilder struct {
	Contents strings.Builder
}

type StringUpperCaser struct {
	Contents strings.Builder
}

func (i MyInt) Twice() MyInt {
	return i * 2
}

func (s MyString) Len() int {
	return len(s)
}

func (mb MyBuilder) Hello() string {
	return "Hello, Gophers!"
}

func (uc StringUpperCaser) ToUpper() string {
	return strings.ToUpper(uc.Contents.String())
}

func (i *MyInt) Double() {
	*i *= 2
}
