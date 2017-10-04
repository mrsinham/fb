package http

import (
	"net/http"
	"fmt"
)

func fizzBuzzHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: validation of the parameters

	q := r.URL.Query()

	if s1 := q.Get("string1"); s1 == "" || len(s1) == 0 {
		writeError(w)
		return
	}

	// TODO: fizzbuzz



}

func fizzBuzz(r *http.Request)  {
	
} 

func writeError(w http.ResponseWriter, code int, response string) {
	w.WriteHeader(code)
	w.Write([]byte(fmt.Sprintf(`{"code":%d,"error":"%q"}`, code, response)))
}
