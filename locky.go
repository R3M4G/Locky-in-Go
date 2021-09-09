package main

import (
	"fmt"
	"os"
	"math/rand"
	"time"
)

func main() {
	name := os.Args[1]
	first := []string{"had", "throws", "buys", "bites", "blows", "hates", "likes", "bangs", "kisses", "has", "boils", "wants", "roasts", "fucks", "eats", "licks", "sucks", "rides", "kills", "screws", "destroys", "cuts"}
	second := []string{"your", "our", "racist", "hairy", "many", "green", "giant", "all", "three", "five", "hundred", "ugly", "my", "black", "nice", "small", "horny", "little", "ten"}
	third := []string{"hammers", "apples", "monkeys", "babies", "kids", "hands", "horses", "trees","stones", "socks", "cats", "asses", "computers", "mangoes", "dildos", "poles", "dogs", "pussies", "aliens"}
	seperators := []string{"-", ".", "*", "+", "~", "_", "/", "\\", "@"}
	for i := 0; i< 5; i++ {
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(first))
	pick := first[randomIndex]
	randomIndex2 := rand.Intn(len(second))
	pick2 := second[randomIndex2]
	randomIndex3 := rand.Intn(len(third))
	pick3 := third[randomIndex3]
	randomIndex4 := rand.Intn(len(seperators))
	pick4 := seperators[randomIndex4]
	generated := name+pick4+pick+pick4+pick2+pick4+pick3
	fmt.Println(generated)
}
}