package main

import (
	"context"
	"log/slog"

	"github.com/sourcegraph/conc/pool"
)

func ImportWorker(ctx context.Context, q chan string) {
	p := pool.New().WithContext(ctx)

	for path := range q {
		path := path

		p.Go(func(ctx context.Context) error {
			// read file

			// extract metadata

			// extract exif

			// extract or create thumbnail

			// store in db

			slog.Info("IMPORTED", "path", path)

			return nil
		})
	}
}
