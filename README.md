# Runtime Validation

--> [!NOTE]
> Not yet ready for use. Still working in additional essential rules.

go-runtimevalidation is a Go library that provides a simple way to validate values and structs in a Go program..The main attraction is that it provides the ability to parse and compile validation rules at runtime. Unlike (go-validator)[https://github.com/go-playground/validator] which requires struct decoration at compile time.

go-runtimevalidation allows you to define validation rules as strings and apply them to any struct at runtime. The library is designed to be simple and easy to use. Many of the ideation about which validations rules should be included comes from (go-validator)[https://github.com/go-playground/validator]. However, all the parsing/compilation logic is build from scratch and it is more flexible and can be used in a wider range of scenarios.

The library is still in the early stages of development, so there may be bugs and missing features. If you find any issues, please report them on the (GitHub issue tracker)[https://github.com/helshabini/go-runtimevalidation/issues].

## Installation

```bash
go get github.com/helshabini/go-runtimevalidation
```

## Usage

The library provides a simple API to define validation rules and apply them to structs. The following example demonstrates how to use the library to validate a struct.

```go
package main

import (
    "fmt"
    validation "github.com/helshabini/go-runtimevalidation"
)

type User struct {
    Name string
    Age  int
    Country string
}

func main() {
    user := User{
        Name: "John",
        Age:  30,
        Country: "USA",
    }

    namerule := "required&&alpha" // Validation string, will be compiled at runtime
    namevgroup , err := validation.Parse(namerule)
	if err != nil {
		fmt.Println(err)
		return
	}
    // The library validates all the rules (even if one of them fails) and returns all the errors, this is to allow the user to display all the errors at once
    errs = namevgroup.Validate(user.Name, nil)
	for _, err := range errs {
		fmt.Println(err)
	}

    // Validating age
    agerule := "required&&min:18&&max:100" // Validation string, will be compiled at runtime
    agevgroup , err := validation.Parse(agerule)
    if err != nil {
        fmt.Println(err)
        return
    }
    // The library validates all the rules (even if one of them fails) and returns all the errors, this is to allow the user to display all the errors at once
    errs = agevgroup.Validate(user.Age, nil)
    for _, err := range errs {
        fmt.Println(err)
    } 

    countryrule := "required&&oneof:USA,UK,Canada" // Validation string, will be compiled at runtime
    countryvgroup , err := validation.Parse(countryrule)
    if err != nil {
        fmt.Println(err)
        return
    }
    // The library validates all the rules (even if one of them fails) and returns all the errors, this is to allow the user to display all the errors at once
    errs = countryvgroup.Validate(user.Country, nil) 
    for _, err := range errs {
        fmt.Println(err)
    } 
}
```

## Advanced Usage

The library also has the ability to evaluate a value as it compares to another value within the same struct. The following example demonstrates how to use the library to validate a value as it compares to another value in its parent struct.

```go

package main

import (
    "fmt"
    validation "github.com/helshabini/go-runtimevalidation"
)

type User struct {
    Name string
    Age  int
    Country string
    Password string
    ConfirmPassword string
}

func main() {
    user := User{
        Name: "John",
        Age:  30,
        Country: "USA",
        Password: "password",
        ConfirmPassword: "password",
    }

    passwordrule := "required" // Validation string, will be compiled at runtime
    passwordvgroup , err := validation.Parse(passwordrule)
    if err != nil {
        fmt.Println(err)
        return
    }
    // The library validates all the rules (even if one of them fails) and returns all the errors, this is to allow the user to display all the errors at once
    errs = passwordvgroup.Validate(user.Password, nil)
    for _, err := range errs {
        fmt.Println(err)
    } 

    // Validating confirm password
    confirmpasswordrule := "eq:$ConfirmPassword" // Validation string, will be compiled at runtime
    confirmpasswordvgroup , err := validation.Parse(confirmpasswordrule)
    if err != nil {
        fmt.Println(err)
        return
    }
    // The library validates all the rules (even if one of them fails) and returns all the errors, this is to allow the user to display all the errors at once
    errs = confirmpasswordvgroup.Validate(user.ConfirmPassword, user) //Note we pass the parent compare twi values 
    for _, err := range errs {
        fmt.Println(err)
    } 
}
```


