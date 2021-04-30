package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

func main() {
	// read in the contents of an image to memory
	dat, err := ioutil.ReadFile("./some.png")
	if err != nil {
		panic(err)
	} // dat isn't currently used
	// seed our rng using time method
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	// generate a random number to check on the image
	x := rng.Intn(3000000) // later make this the resolution of the image

	fmt.Println(x)
}
