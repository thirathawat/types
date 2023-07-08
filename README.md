# Package `types`

The `types` package provides custom types for various purposes. Currently, it includes the `Hash` type for representing hashed strings.

## Installation

To install the `types` package, use the following command:

```shell
go get -u github.com/thirathawat/types
```

## Usage

Import the package into your Go program:

```go
import "github.com/thirathawat/types"
```

### Hash Type

The `Hash` type is a custom type based on the `string` type. It represents a hashed string value.

#### Creating a Hash

To create a `Hash` value, simply assign a string value to it:

```go
var hashedPassword types.Hash = "mysecretpassword"
```

#### JSON Serialization

The `Hash` type implements the `json.Marshaler` and `json.Unmarshaler` interfaces, allowing you to easily serialize and deserialize JSON representations of the `Hash` value.

Example usage:

```go
hashedPassword := types.Hash("mysecretpassword")

// Serialize to JSON
data, err := json.Marshal(hashedPassword)
if err != nil {
    log.Fatal(err)
}

// Deserialize from JSON
var restoredHash types.Hash
err = json.Unmarshal(data, &restoredHash)
if err != nil {
    log.Fatal(err)
}
```

#### String Representation

The `Hash` type also implements the `fmt.Stringer` interface, providing a string representation of the hashed value.

Example usage:

```go
hashedPassword := types.Hash("mysecretpassword")

str := hashedPassword.String()
fmt.Println(str) // Output: "$2a$10$ldYn/9LioR9v/EsHf/6qQ.mtD7iDQwB58gUxCW.GgNt3G82Mu4Ebe"
```

#### Comparison

The `Hash` type provides a `Compare` method to compare the hashed value with a plain string.

Example usage:

```go
hashedPassword := types.Hash("mysecretpassword")

isMatch := hashedPassword.Compare("mysecretpassword")
fmt.Println(isMatch) // Output: true

isMatch = hashedPassword.Compare("wrongpassword")
fmt.Println(isMatch) // Output: false
```

## License

This package is distributed under the MIT License. See the `LICENSE` file for more information.
