# Phone Package

The **phone** package provides functionality for formatting and validating phone numbers. It includes methods for formatting phone numbers according to international standards and validating phone numbers based on their formats.

## Installation

To use the **phone** package, import it into your Go project:

````
 import "github.com/aleslanger/phone"
````

## Usage

### Phone Struct

The **Phone** struct represents a phone number with a name and a code. It has the following fields:

* **Name** (string): The name of the country or region associated with the phone number.
* **Code** (string): The country code or prefix for the phone number.

### Codes
The **Codes** function returns a slice of **Phone** objects representing country names and their corresponding phone codes. 
It provides a list of commonly used country codes for reference.

Example usage:

```go
codes := phone.Codes()
for _, code := range codes {
fmt.Println("Name:", code.Name)
fmt.Println("Code:", code.Code)
fmt.Println()
}
```

### FormatPhone

The FormatPhone function takes a phone number as input and returns the formatted version of the phone number according to international standards. 
It removes any unnecessary characters and adds spaces between number groups.

Example usage:

```go
phoneNumber := "+14325551234"
formattedNumber := phone.FormatPhone(phoneNumber)
fmt.Println("Formatted number:", formattedNumber)
```

### ValidatePhoneNo

The **ValidatePhoneNo** function checks if a given phone number is valid based on its format. 
It returns **true** if the phone number is valid, and **false** otherwise.

Example usage:

```go
phoneNumber := "+14325551234"
isValid := phone.ValidatePhoneNo(phoneNumber)
if isValid {
	fmt.Println("Phone number is valid")
} else {
	fmt.Println("Phone number is invalid")
}
```

### ValidatePhonePrefix

The **ValidatePhonePrefix** function checks if a given phone number prefix is valid. 
It verifies if the prefix corresponds to a known country code.

Example usage:

```go
prefix := "+1"
isValidPrefix := phone.ValidatePhonePrefix(prefix)
if isValidPrefix {
	fmt.Println("Phone prefix is valid")
} else {
	fmt.Println("Phone prefix is invalid")
}
```

## License

This code is released under the [MIT License](LICENSE.md).





