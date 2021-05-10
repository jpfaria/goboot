package server

import (
	"context"
	"sync"
)

type Server interface {
	Serve(context.Context)
}

func Serve(ctx context.Context, srvs ...Server) {

	switch len(srvs) {
	case 0:
		panic("no servers configured")
	case 1:
		srvs[0].Serve(ctx)
	default:
		wg := new(sync.WaitGroup)
		wg.Add(len(srvs))

		for _, srv := range srvs {
			srv := srv
			go func() {
				srv.Serve(ctx)
				wg.Done()
			}()
		}

		wg.Wait()
	}

}
