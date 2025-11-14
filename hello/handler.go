package function

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	var input []byte

	if r.Body != nil {
		defer r.Body.Close()

		body, _ := io.ReadAll(r.Body)

		input = body
	}

	for _, v := range r.Header {
		log.Printf("Header: %s", v)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Body: %s", string(input))))
}
