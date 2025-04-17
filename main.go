package main

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

var acceptedOperations = []string{"SET", "GET"}

func operationIsAccepted(receivedOperation *string) bool {
	for _, currentOperation := range acceptedOperations {
		if *receivedOperation == currentOperation {
			return true
		}
	}
	return false
}

func main() {

	var operation, key, value string

	fmt.Println("---| WELLCOME TO MOR! | ---\n\n")

	for {

		_, _ = fmt.Scanln(&operation, &key, &value)

		operation = strings.ToUpper(operation)
		if !operationIsAccepted(&operation) {
			fmt.Println("Operation needs to be either 'SET' or 'GET'")
			continue
		}

		switch operation {

		case "SET":
			dataBase[key] = value
			fmt.Println("Added ", value, " to ", key)
			break

		case "GET":
			if dataBase[key] == "" {
				fmt.Println("Theres no ->", key, "<- key")
				continue
			}
			fmt.Println(dataBase[key])
			break
		}

	}
}
