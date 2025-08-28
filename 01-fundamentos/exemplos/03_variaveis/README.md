# 🔹 Variáveis e Constantes

## 1. Declarando variáveis

Go tem três formas principais:

```go
package main

import "fmt"

func main() {
    // 1) declaração explícita com tipo
    var a int = 10
    fmt.Println("a =", a)

    // 2) declaração implícita (Go infere o tipo)
    var b = 20
    fmt.Println("b =", b)

    // 3) declaração curta (somente dentro de funções)
    c := 30
    fmt.Println("c =", c)
}
```

⚠️ Regra: fora de funções você **não pode** usar `:=`.

---

## 2. Valores zero

Quando você declara sem inicializar, o Go atribui um valor padrão:

```go
var s string   // "" (string vazia)
var i int      // 0
var f float64  // 0
var ok bool    // false
```

---

## 3. Constantes

Constantes são imutáveis e conhecidas em tempo de compilação:

```go
const Pi = 3.14159
const Mensagem = "Olá, Go!"
```

Também dá pra criar blocos de constantes:

```go
const (
    A = 1
    B = 2
    C = 3
)
```

---

## 4. `iota` (contador automático)

`iota` é um identificador especial que gera números sequenciais:

```go
const (
    Domingo = iota // 0
    Segunda        // 1
    Terca          // 2
)
```

Você pode usar para enums, flags etc.

---

## 5. Exemplos práticos

📂 [./exercicios/example.go](./exercicios/example.go)

---

## 6. Exercícios sugeridos

1. **Conversor de temperatura**

   * Crie uma variável em °C e converta para °F.
     Fórmula: `F = (C * 9/5) + 32`.

2. **Enumeração com `iota`**

   * Crie constantes representando os dias da semana e imprima-os.

3. **Calculadora simples**

   * Use `var` e `:=` para criar duas variáveis numéricas.
   * Imprima a soma, subtração, multiplicação e divisão.