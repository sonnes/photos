package main

import (
	"log/slog"
	"strings"
	"sync"

	"github.com/karrick/godirwalk"
	"github.com/urfave/cli/v2"
)

func index(c *cli.Context) error {
	paths := c.StringSlice("paths")

	q := make(chan string)

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		ImportWorker(c.Context, q)
		wg.Done()
	}()

	for _, path := range paths {
		if err := walk(path, q); err != nil {
			return err
		}
	}

	close(q)
	wg.Wait()

	return nil
}

func walk(path string, q chan string) error {
	return godirwalk.Walk(path, &godirwalk.Options{
		ErrorCallback: func(filename string, err error) godirwalk.ErrorAction {
			slog.Error("Error", "path", filename, "error", err)

			return godirwalk.SkipNode
		},
		Callback: func(filename string, de *godirwalk.Dirent) error {
			if strings.Contains(filename, ".git") {
				return godirwalk.SkipThis
			}

			if !IsMedia(filename) {
				return nil
			}

			q <- filename

			return nil
		},
		Unsorted: true,
	})
}
