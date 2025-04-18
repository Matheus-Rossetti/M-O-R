package main

// TODO Make this support video files, I'll need this for the Streaming app

import (
	"fmt"
	"strings"
)

var dataBase = make(map[string]string)

//func MOR(w http.ResponseWriter, r *http.Request) {
//
//	if r.Method == http.MethodPost {
//		body, err := io.ReadAll(r.Body)
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusBadRequest)
//		}
//		defer r.Body.Close()
//
//		err = json.Unmarshal(body, &dataBase)
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusBadRequest)
//		}
//
//		fmt.Fprint(w, dataBase)
//	}
//
//	if r.Method == http.MethodGet {
//
//		body, err := io.ReadAll(r.Body)
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusBadRequest)
//		}
//		defer r.Body.Close()
//
//		fmt.Fprint(w, dataBase[string(body)])
//
//		fmt.Println(dataBase[string(body)])
//
//	}
//
//}

var acceptedOperations = []string{"SET", "GET", "DEL", "QUIT"}

func operationIsAccepted(receivedOperation *string) bool {
	for _, currentOperation := range acceptedOperations {
		if *receivedOperation == currentOperation {
			return true
		}
	}
	return false
}

func set(key *string, value *string) {
	dataBase[*key] = *value
	fmt.Println("DONE!")
	return
}
func get(key *string) {
	if dataBase[*key] == "" {
		fmt.Println("Theres no ->", *key, "<- key")
		return
	}
	fmt.Println(dataBase[*key])
}
func del(key *string) {
	_, keyExists := dataBase[*key]
	if !keyExists {
		fmt.Println("There's no ->", *key, "<-  key")
		return
	}
	delete(dataBase, *key)
	fmt.Println("DONE!")
	return
}

func main() {

	var operation, key, value string

	fmt.Println("---| WELLCOME TO MOR! |---\n\n")

	for {
		_, _ = fmt.Scanln(&operation, &key, &value)

		operation = strings.ToUpper(operation)
		if !operationIsAccepted(&operation) {
			fmt.Println("Operation needs to be either 'SET' or 'GET'")
			continue
		}

		switch operation {
		case "SET":
			set(&key, &value)
			break

		case "GET":
			get(&key)
			break

		case "DEL":
			del(&key)
			break

		case "QUIT":
			return
		}
	}
}
