package main

import (
	"fmt"
	"go-runtimevalidation/validation"
)

func main() {
	ruletext := "required && alpha"
	groups, err := validation.Parse(ruletext)
	if err != nil {
		fmt.Println(err)
		return
	}

	// test1 is an object with a alpha field
	test1 := TestStruct{StrField: "test"}

	errs := groups.Validate(test1.StrField, nil)
	for _, err := range errs { // should be empty
		fmt.Println(err)
	}

	// test2 is an object with a non-alpha field
	test2 := TestStruct{StrField: "123"}

	errs = groups.Validate(test2.StrField, nil)
	for _, err := range errs { // should contain an error
		fmt.Println(err)
	}
}

type TestStruct struct {
	StrField string
}
