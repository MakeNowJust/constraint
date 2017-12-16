package constraint_test

import (
	"fmt"

	"github.com/MakeNowJust/constraint"
)

func Example() {
	tags, _ := constraint.NewTagSet("foo", "bar")

	for _, s := range []string{
		// A simple tag name is valid constraint of course.
		"foo",
		"bar",

		// 'NOT' operator is bang, so it mean 'NOT fizz'.
		"!fizz",

		// 'OR' operator is space, so it means 'fizz OR foo OR buzz'.
		"fizz foo buzz",

		// 'AND' operator is comma, so it means 'foo AND bar'.
		"foo,bar",

		// It is complete example.
		"fizz,!buzz !foo,bar",
	} {
		c, _ := constraint.Parse(s)
		fmt.Printf("%-19s => %t\n", c, c.Eval(tags))
	}
	// Output:
	// foo                 => true
	// bar                 => true
	// !fizz               => true
	// fizz foo buzz       => true
	// foo,bar             => true
	// fizz,!buzz bar,!foo => false
}

func ExampleIsValidTag() {
	for _, s := range []string{
		// Letters are valid tag name.
		"foo",

		// Digits are valid also.
		"123",

		// They can be mixed.
		"foo123",

		// Tag name can contain '.' and '_'.
		"foo_bar.baz",

		// Unicode letters (and digits) are valid.
		"トムとジェリー",

		// Empty string is invalid.
		"",

		// Neither letter nor digits is invalid.
		"tom&jerry",
	} {
		fmt.Printf("%#v: %t\n", s, constraint.IsValidTag(s))
	}
	// Output:
	// "foo": true
	// "123": true
	// "foo123": true
	// "foo_bar.baz": true
	// "トムとジェリー": true
	// "": false
	// "tom&jerry": false
}
