package unknownjson_test

import (
	"fmt"

	unknownjson2 "github.com/chaos-io/core/x/encoding/unknownjson"
)

func Example() {
	var s struct {
		A int `json:"a"`

		Unknown unknownjson2.Store `json:"-" unknown:",store"`
	}

	js := []byte(`{"a":1,"b":2}`)
	if err := unknownjson2.Unmarshal(js, &s); err != nil {
		panic(err)
	}

	var err error
	js, err = unknownjson2.Marshal(s)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(js))
	// Output:
	// {"a":1,"b":2}
}
