## Hash Quick start

```go
package main

import (
    "fmt"
    "crypto/md5"
    "crypto/sha1"
    "github.com/buzzxu/boys/common/structs"
)

type S struct {
    Str string
    Num int
}

func main() {
    s := S{"hello", 123}

    hash, err := structs.Hash(s, 1)
    if err != nil {
        panic(err)
    }
    fmt.Println(hash)
    // Prints: v1_41011bfa1a996db6d0b1075981f5aa8f

    fmt.Println(structs.Version(hash))
    // Prints: 1

    fmt.Printf("%x\n", structs.Md5(s, 1))
    // Prints: 41011bfa1a996db6d0b1075981f5aa8f

    fmt.Printf("%x\n", structs.Sha1(s, 1))
    // Prints: 5ff72df7212ce8c55838fb3ec6ad0c019881a772

    fmt.Printf("%x\n", md5.Sum(structs.Dump(s, 1)))
    // Prints: 41011bfa1a996db6d0b1075981f5aa8f

    fmt.Printf("%x\n", sha1.Sum(structs.Dump(s, 1)))
    // Prints: 5ff72df7212ce8c55838fb3ec6ad0c019881a772
}
```

## Struct tags

structhash supports struct tags in the following forms:

* `hash:"-"`, or
* `hash:"name:{string} version:{number} lastversion:{number} method:{string}"`

All fields are optional and may be ommitted. Their semantics is:

* `-` - ignore field
* `name:{string}` - rename field (may be useful when you want to rename field but keep hashes unchanged for backward compatibility)
* `version:{number}` - ignore field if version passed to structhash is smaller than given number
* `lastversion:{number}` - ignore field if version passed to structhash is greater than given number
* `method:{string}` - use the return value of a field's method instead of the field itself

Example:

```go
type MyStruct struct {
    Ignored    string `hash:"-"`
    Renamed    string `hash:"name:OldName version:1"`
    Legacy     string `hash:"version:1 lastversion:2"`
    Serialized error  `hash:"method:Error"`
}
```

## Nil values

When hash is calculated, nil pointers, nil slices, and nil maps are treated equally to zero values of corresponding type. E.g., nil pointer to string is equivalent to empty string, and nil slice is equivalent to empty slice.