package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// type helloWorldResponse struct {
// 	Message string
// }
type helloWorldResponse struct {
	Message string `json:"message"`
}

func main() {
	port := 8080

	http.HandleFunc("/helloworld", helloWorldHandler)

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	response := helloWorldResponse{Message: "Hello World"}
	data, err := json.Marshal(response)
	if err != nil {
		panic("Ooops")
	}
	fmt.Fprint(w, string(data))
}

// func Marshal(v interface{}) ([]byte, error) {
// 	e := newEncodeState()

// 	err := e.marshal(v, encOpts{escapeHTML: true})
// 	if err != nil {
// 		return nil, err
// 	}
// 	buf := append([]byte(nil), e.Bytes()...)

// 	e.Reset()
// 	encodeStatePool.Put(e)

// 	return buf, nil
// }

// func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
// 	b, err := Marshal(v)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var buf bytes.Buffer
// 	err = Indent(&buf, b, prefix, indent)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return buf.Bytes(), nil
// }
