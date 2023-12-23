
# golang

## PACKAGES
- a package is a way to organize and structure code. It provides a mechanism for grouping related Go source files together
- each go source file begins with a package declaration, indicating the package to which the file belongs.
- package main is a reserved package name which tells start point, we just run `go build` which creates a executable file, 
- we should have atleast once package per application or go code
- a single package can have multiple files

## MODULES
- one module consists of multiple packages
- modules provide a way to manage dependencies and versioning for your Go projects


## INPUT VALUES
- fmt.Scan(&<var_name>) `var_name` should be declared with a type and we need to pass the address of it for scan
- We can't scan string with spaces using scan

## SPRINTF
- fmt.sprintf() can be used to stored the formatted output into some string

## Functions
- func calculatePower(val1, val2, val3, val4 int) // here all values are expected to be of int type

## Write data to file
- we can write string so use sprintf() to get the string

## strconv
- string conversion used to convert string from string to float or int, will return two values one is converted value and other error

## error handling
- we should write function such that it doesn't crash application as it doesn't have try catch, we need to import errors module
