package main

import (
	"io"
	"os/exec"
)

/*
	TODO change ExecAndMarshall():
		1. Move 'out' to be used as a return value, not a parameter. Adapt main() accordingly.
		2. Use generics for the return type, inferring the type to unmarshal from what the caller defines.

	Thanks to Norman for this example.
*/

import (
	"encoding/json"
	"fmt"
)

func ExecAndMarshall(command []string, out interface{}) error {

	cmd := exec.Command(command[0], command[1:]...)
	stdout, err := cmd.StdoutPipe()

	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	bytes, err := io.ReadAll(stdout)

	if err != nil {
		return err
	}

	if err := cmd.Wait(); err != nil {
		return err
	}

	// var out T
	if err = json.Unmarshal(bytes, out); err != nil {
		return err
	}

	// return T, err
	return nil
}

func main() {
	type P struct {
		Name    string
		Country string
	}

	var err error
	var p P
	err = ExecAndMarshall([]string{"echo", `{"name": "Jane","country": "AU"}`}, &p)
	if err != nil {
		panic(err)
	}

	fmt.Println(p)

	type C struct {
		Name     string
		Timezone int
	}

	var c C
	err = ExecAndMarshall([]string{"echo", `{ "name": "Australia", "timezone": 1030 }`}, &c)
	if err != nil {
		panic(err)
	}

	fmt.Println(c)
}
