package main

import "net/http"

func main() {
	http.ListenAndServe(":3000", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://github.com/a10adotapp/blog.a10a.app", http.StatusMovedPermanently)
	}))
}
