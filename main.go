package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func main() {
	//Toda api precisa de handlers, que são rotas que o servidor vai receber requests e responder com responses 
	http.handleFunc("/sum",func(w http.ResponseWriter,r *http.Request)
	{
		query := r.URl.Query()
		n1 := query.Get("n1")
		//query.Get volta string, temos que converter para int para realizar a conta
		n2 := query.Get("n2")
		//query.Get volta string ,  temos que connverter para int para realizar a conta
		number1, _ := strconv.Atoi(n1)
		//convertendo n1 para int
		number2, _ := strconv.Atoi(n2)
		//convertendo n2 para int
		result := number1 + number2
		response := map[string]int{}
		response["result"] = result

		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}

	//Basta criar outros handlers aqui, semelhantes ao primeiro}

	log.Print("http server listening on port 80")

	//Esse comando inicia o server http escutando na porta 80,pra parar de rodar depois que iniciar],basta apertar ctrl+c
	http.ListenAndServe(":80,nil")
}
