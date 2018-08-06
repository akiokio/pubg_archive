package main

import (
	"log"
	"os/user"
	"path"
	"sync"
	"time"

	"github.com/akiokio/pubg-archive/copydir"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	currentUser, err := user.Current()
	if err != nil {
		log.Println(err)
	}
	if currentUser.HomeDir == "" {
		log.Println("cannot find user-specific home dir")
	}
	source := path.Join(currentUser.HomeDir, "\\AppData\\Local\\TslGame\\Saved\\Demos")
	dest := path.Join(currentUser.HomeDir, "\\Desktop\\PUBG - Matches - all")
	go func() {
		err = copydir.CopyDir(source, dest)
		wg.Done()
	}()

	wg.Wait()
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("Done working. Closing in 3 seconds.")
		time.Sleep(3 * time.Second)
	}
}
