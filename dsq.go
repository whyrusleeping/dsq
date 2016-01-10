package main

import (
	"fmt"
	"os"

	ds "github.com/ipfs/go-ipfs/Godeps/_workspace/src/github.com/ipfs/go-datastore"
	"github.com/ipfs/go-ipfs/repo/fsrepo"
)

func main() {
	bkp, err := fsrepo.BestKnownPath()
	if err != nil {
		panic(err)
	}

	r, err := fsrepo.Open(bkp)
	if err != nil {
		panic(err)
	}

	k := ds.NewKey(os.Args[1])

	val, err := r.Datastore().Get(k)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(string(val.([]byte)))
}
