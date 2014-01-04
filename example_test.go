// Copyright 2013 Frank Schroeder. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goproperties

import (
	"fmt"
	"log"
)

func ExampleLoad_ISO() {
	buf := []byte("key = ISO-8859-1 value with unicode literal \\u2318 and umlaut ")
	buf = append(buf, 0xE4) // 0xE4 == ä
	p, _ := Load(buf, ISO_8859_1)
	v, ok := p.Get("key")
	fmt.Println(ok)
	fmt.Println(v)
	// Output:
	// true
	// ISO-8859-1 value with unicode literal ⌘ and umlaut ä
}

func ExampleLoad_UTF8() {
	p, _ := Load([]byte("key = UTF-8 value with unicode character ⌘ and umlaut ä"), UTF8)
	v, ok := p.Get("key")
	fmt.Println(ok)
	fmt.Println(v)
	// Output:
	// true
	// UTF-8 value with unicode character ⌘ and umlaut ä
}

func Example_Properties_GetDefault() {
	p, _ := Load([]byte("key=value"), ISO_8859_1)
	v := p.GetDefault("another key", "default value")
	fmt.Println(v)
	// Output:
	// default value
}

func Example() {
	// Decode some key/value pairs with expressions
	p, err := Load([]byte("key=value\nkey2=${key}"), ISO_8859_1)
	if err != nil {
		log.Fatal(err)
	}

	// Get a valid key
	if v, ok := p.Get("key"); ok {
		fmt.Println(v)
	}

	// Get an invalid key
	if _, ok := p.Get("does not exist"); !ok {
		fmt.Println("invalid key")
	}

	// Get a key with a default value
	v := p.GetDefault("does not exist", "some value")
	fmt.Println(v)

	// Dump the expanded key/value pairs of the Properties
	fmt.Println("Expanded key/value pairs")
	fmt.Println(p)

	// Dump the raw key/value pairs.
	fmt.Println("Raw key/value pairs")
	fmt.Println(p.Dump())
	// Output:
	// value
	// invalid key
	// some value
	// Expanded key/value pairs
	// key = value
	// key2 = value
	//
	// Raw key/value pairs
	// key = value
	// key2 = ${key}
}