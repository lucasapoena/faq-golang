# 📘 Primeiro programa em Go — `hello.go`

## 1. Estrutura mínima

Todo programa Go começa com **um pacote principal** e **uma função principal**:

```go
package main       // define que este é o pacote "principal"
import "fmt"       // importa pacote de formatação/print

func main() {      // ponto de entrada do programa
    fmt.Println("Hello, Go!") // imprime no console
}
```

* `package main` → obrigatório para programas executáveis.
* `func main()` → função de entrada (Go começa a execução aqui).
* `fmt.Println` → função da biblioteca padrão que imprime no console e adiciona quebra de linha.

---

## 2. Compilando e executando

### Rodar direto (interpretação JIT pelo runtime):

```bash
go run hello.go
```

### Compilar para binário:

```bash
go build hello.go
./hello       # Linux/macOS
hello.exe     # Windows
```
---

## 3. `fmt.Printf` e formatações

Além de `fmt.Println`, muito usado para debug, Go também oferece `fmt.Printf`, que permite **formatar strings** de forma mais controlada:

```go
nome := "João das Neves"
idade := 42
fmt.Printf("Olá, %s! Você tem %d anos.\n", nome, idade)
```

Explicando:

* `fmt.Printf` → imprime com **formatação** (`Printf` = *print formatted*).
* `"Olá, %s! Você tem %d anos.\n"` → string de formatação com **verbos**:

  * `%s` → insere uma string.
  * `%d` → insere um número inteiro.
  * `\n` → quebra de linha.
* Os argumentos após a string (`nome, idade`) substituem os verbos na ordem.

👉 Saída:

```
Olá, João das Neves! Você tem 42 anos.
```

---

## 4. Exercícios práticos

1. **Alterar mensagem**: mude `"Hello, Go!"` para imprimir `"Olá, mundo!"`.
2. **Adicionar variável**: declare `nome := "Lucas"` e use `fmt.Printf("Olá, %s!\n", nome)`.
3. **Dois prints**: faça o programa imprimir duas linhas diferentes.