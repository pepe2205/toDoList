package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

/*
	Cria o objeto de Tarefa que contem seu identificador e uma variavel booleana

para indicar se a tarefa foi concluida
*/
type Tarefa struct {
	Nome      string `json:"nome"`
	Concluida bool   `json:"concluida"`
}

const arquivo = "tarefas.json" // Declara o arquivo onde as tarefas adicionadas serao guardadas

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: go run main.go [add/list/done/delete] [nome da tarefa]")
		return
	}
	comando := os.Args[1]

	/* A estrutura de controle abaixo é utilizada para
	selecionar a ação desejada para o manuseio da lista de tarefas,
	caso seja uma opção inválida, é retornado o texto
	"Comando desconhecido"*/
	switch comando {
	case "add":
		adicionarTarefa(os.Args[2:])
	case "list":
		listarTarefas()
	case "done":
		marcarComoConcluida(os.Args[2])
	case "delete":
		deletarTarefa(os.Args[2])
	default:
		fmt.Println("Comando desconhecido")
	}
}
func adicionarTarefa(nome []string) {
	tarefa := Tarefa{Nome: strings.Join(nome, " "), Concluida: false} // por padrao, o status de conclusao da tarefa é falso
	tarefas := carregarTarefas()
	tarefas = append(tarefas, tarefa) // adiciona a nova tarefa na lista
	salvarTarefas(tarefas)
	fmt.Println("Tarefa adicionada com sucesso!")
}
func listarTarefas() { // funcao criada para listar todas as tarefas e seu status de conclusao
	tarefas := carregarTarefas()
	for i, tarefa := range tarefas {
		status := "Pendente"
		if tarefa.Concluida {
			status = "Concluida" // se a tarefa esta concluida, o status da tarefa é exibida como "Concluida"
		}
		fmt.Printf("%d - %s (%s)\n\n", i+1, tarefa.Nome, status)
	}
}
func marcarComoConcluida(nome string) { // utilizado para alterar o status de conclusão da tarefa para verdadeira
	tarefas := carregarTarefas()
	for i, tarefa := range tarefas {
		if tarefa.Nome == nome {
			tarefas[i].Concluida = true
			salvarTarefas(tarefas)
			fmt.Println("Tarefa marcada como concluida!")
			return
		}
	}
	fmt.Println("Tarefa nao encontrada")
}
func deletarTarefa(nome string) {
	tarefas := carregarTarefas()
	for i, tarefa := range tarefas {
		if tarefa.Nome == nome {
			tarefas = append(tarefas[:i], tarefas[i+1:]...)
			salvarTarefas(tarefas)
			fmt.Println("Tarefa deletada com sucesso!")
			return
		}
	}
	fmt.Println("Tarefa nao encontrada")
}

/*
	essa função lê o arquivo "tarefas.json" e retorna uma variavel contendo as tarefas e seus status

caso não seja possível, retorna "Erro ao carregar do json: codigo de erro"
*/
func carregarTarefas() []Tarefa {
	var tarefas []Tarefa
	dados, err := ioutil.ReadFile(arquivo)
	if err != nil {
		return tarefas
	}
	err = json.Unmarshal(dados, &tarefas)
	if err != nil {
		fmt.Println("Erro ao carregar do json:", err)
	}
	return tarefas
}

// salva as modificações feitas nas tarefas no arquivo "tarefas.json"
func salvarTarefas(tarefas []Tarefa) {
	dados, err := json.Marshal(tarefas)
	if err != nil {
		fmt.Println("Erro ao salvar tarefas:", err)
		return
	}
	err = ioutil.WriteFile(arquivo, dados, 0644)
	if err != nil {
		fmt.Println("Erro ao salvar tarefas:", err)
	}
}
