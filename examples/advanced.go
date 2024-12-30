package main

import (
	"fmt"
	"go-runtimevalidation/validation"
)

func advanced() {
  ruletext := "eq:$ConfirmPassword"
	groups, err := validation.Parse(ruletext)
	if err != nil {
		fmt.Println(err)
		return
	}

	// test1 is an object with a alpha field
  test1 := AdvancedStruct{Password: "test", ConfirmPassword: "test"}

	errs := groups.Validate(test1.Password, test1)
	for _, err := range errs { // should be empty
		fmt.Println(err)
	}

	// test2 is an object with a non-alpha field
  test2 := AdvancedStruct{Password: "test", ConfirmPassword: "123"}

	errs = groups.Validate(test2.Password, test2)
	for _, err := range errs { // should contain an error
		fmt.Println(err)
	}
}

type AdvancedStruct struct {
	Password string
  ConfirmPassword string
}
