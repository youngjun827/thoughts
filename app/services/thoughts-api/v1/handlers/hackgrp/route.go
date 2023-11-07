package hackgrp

import (
	"github.com/go-chi/chi/v5"
	"github.com/youngjun827/thoughts/foundation/web"
)

func Routes(router chi.Router) {
    router.Route("/hack", func(c chi.Router) {
        c.Get("/", web.HandlerAdapter(Hack))
    })
}
