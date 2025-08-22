# Instalação e configuração do Go

Este guia cobre Windows, Linux e macOS. Ao final, você terá o Go configurado e saberá criar e rodar um projeto com módulos.

## 1) Instalar o Go
Baixe o instalador em https://go.dev/dl/ e siga o assistente do seu sistema.

- **Windows**: normalmente em `C:\Program Files\Go`
- **Linux/macOS**: extrair em `/usr/local/go`

## 2) Variáveis de ambiente

Certifique-se de que o `go` e seu diretório de binários do usuário estão no PATH.

**Linux/macOS (adicione ao `~/.bashrc` ou `~/.zshrc`):**
```bash
export PATH=$PATH:/usr/local/go/bin
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

**Windows (Variáveis do Sistema):**
- `GOROOT = C:\Program Files\Go`
- `GOPATH = C:\Users\SEU_USUARIO\go`
- Inclua no `PATH`: `C:\Program Files\Go\bin` e `%GOPATH%\bin`

**Verifique:**
```bash
go version
go env
```

## 3) Extensões úteis (VS Code)
- Extensão oficial: [Go](https://code.visualstudio.com/docs/languages/go)
- Ferramentas:
  ```bash
  go install golang.org/x/tools/cmd/goimports@latest
  go install golang.org/x/lint/golint@latest
  ```