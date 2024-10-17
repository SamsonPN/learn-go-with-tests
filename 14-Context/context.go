package main

import (
	"context"
	"fmt"
	"net/http"
)

type Store interface {
	Fetch(ctx context.Context) (string, error)
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// by passing in context,
		// the server code is simpler since it doesn't
		// handle cancellations anymore
		data, err := store.Fetch(r.Context())
		if err != nil {
			return
		}

		fmt.Fprint(w, data)
		// ctx := r.Context()

		// data := make(chan string, 1)

		// go func() {
		// 	data <- store.Fetch()
		// }()

		// // we use select to race between
		// // fetching the data
		// // or cancelling the fetch request

		// // if cancel is called first, then the store will cancel
		// // otherwise, it will grab data from channel and print it
		// select {
		// case d := <-data:
		// 	fmt.Fprint(w, d)
		// // when the context gets cancelled
		// // Done will return a closed channel
		// case <-ctx.Done():
		// 	store.Cancel()
		// }
	}
}
