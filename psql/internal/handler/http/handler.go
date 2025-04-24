package http

import "net/http"

func InitHandlers(
	ah *AccountHandler,
	rh *RoomHandler,
) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/transfer", ah.HandleTransfer)
	mux.HandleFunc("/room", rh.Reserve)
	return mux
}
