package app

import "github.com/KCTHack/kcthack-auth/internal/handler"

func Run() {
	r := handler.NewRouter()
	r.Run(":8000")
}
