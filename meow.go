package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	if len(os.Args) >= 2 {
		fmt.Println("🐱 <" + os.Args[1])
	} else {
		fmt.Println(getGohoushi())
	}
}

func getGohoushi() (s string) {
	neko := []string{
		"🐱",
		"😸",
		"😻",
		"😼",
		"😺",
		"😽",
		"😹",
	}
	suru := []string{
		"する",
		"する",
		"する",
		"する",
		"しない",
		"しない",
		"したくない",
		"するとかありえない",
		"されたい",
		"しろ",
	}
	nyan := []string{
		"にゃん♪",
		"にゃん☆",
		"ニャン☆",
		"ニャン♪",
		"ニャン☆・*:.｡.:",
		"にゃん☆・*:.｡.:",
	}
	return neko[rand.Intn(len(neko))] +
		"「ご奉仕" +
		suru[rand.Intn(len(suru))] +
		nyan[rand.Intn(len(nyan))] +
		"」 " +
		neko[rand.Intn(len(neko))] +
		"「" +
		suru[rand.Intn(len(suru))] +
		nyan[rand.Intn(len(nyan))] +
		"」"
}
