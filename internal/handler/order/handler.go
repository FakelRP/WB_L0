package order

import (
	"context"
	"log"
)

type Handler struct {
}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) Handle(ctx context.Context, msg []byte) error {

	log.Printf("recieved msg: %+v", string(msg))

	//

	return nil
}
