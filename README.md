# GoCRUD

Este projeto é uma API RESTful em Go projetada para gerenciar usuários através das operações CRUD básicas (Criar, Ler, Atualizar, Deletar), armazenando os dados em memória usando um `map` local protegido por `sync.RWMutex`.

Este é um projeto prático projetado para exercitar conceitos básicos de Go, incluindo o roteamento nativo (`http.NewServeMux`) da recém introduzida biblioteca `net/http` do Go 1.22+ e conversão de JSON.

## 🚀 Como iniciar e rodar o projeto

1. **Pré-requisitos**:
   Certifique-se de ter o Go (versão 1.22 ou superior) instalado em sua máquina.

2. **Clone ou baixe o repositório**
   Entre no diretório principal e baixe a dependência utilizada para gerar IDs únicos (UUID):
   ```bash
   go mod tidy
   ```

3. **Inicie o servidor HTTP**
   Execute o seguinte comando na raiz do projeto para inicializar:
   ```bash
   go run main.go
   ```

   O console deve exibir a mensagem:
   `Server listening on :8080`

## 📡 Endpoints da API (Rotas)

A porta de escuta base é a **8080**.
URL base: `http://localhost:8080/api/users`

A estrutura de Usuário subjacente que as respostas JSON fornecerão se baseia em:
```json
{
  "id": "UUID_AQUI",            // Gerado de forma automática
  "first_name": "Jane",         // min 2 a 20 caracteres
  "last_name": "Doe",           // min 2 a 20 caracteres
  "biography": "Lorem Ipsum"    // min 20 a 450 caracteres
}
```

### 1. Criar um novo usuário (POST)
**Endpoint**: `POST /api/users`

**Corpo (Body)** JSON:
```json
{
  "first_name": "Jane",
  "last_name": "Doe",
  "biography": "Fascinada por Go e o ecossistema Web como um todo."
}
```
**Respostas Esperadas:**
- `201 Created`: Retorna o JSON contendo os dados inseridos em uníssono ao `id`.
- `400 Bad Request`: Se os campos excederem ou ficarem abaixo do tamanho limite.

---

### 2. Listar todos os usuários (GET)
**Endpoint**: `GET /api/users`

**Respostas Esperadas:**
- `200 OK`: Lista (Array) JSON de usuários do banco em memória.

---

### 3. Buscar um usuário pelo ID (GET)
**Endpoint**: `GET /api/users/{id}`

- Forneça o `{id}` gerado nos passos anteriores para realizar a operação.

**Respostas Esperadas:**
- `200 OK`: JSON contendo o usuário específico.
- `404 Not Found`: Se o usuário contendo o ID requisitado não for encontrado.

---

### 4. Atualizar um usuário (PUT)
**Endpoint**: `PUT /api/users/{id}`

Este método atualiza toda as informações (não é uma mesclagem parcial), logo, todo o corpo JSON válido contendo `first_name`, `last_name` e `biography` deve ser enviado novamente.

**Corpo (Body)** JSON:
```json
{
  "first_name": "John",
  "last_name": "Doe",
  "biography": "Um nome diferente e atualizado agora que conheço um pouco de Rust também!"
}
```

**Respostas Esperadas:**
- `200 OK`: Usuário contendo o `{id}` foi atualizado com as novas informações.
- `400 Bad Request`: Nome, Sobrenome ou Biografia são inválidos.
- `404 Not Found`: Usuário não encontrado para atualizar.

---

### 5. Deletar um usuário (DELETE)
**Endpoint**: `DELETE /api/users/{id}`

- Remove o recurso da memória para todo o sempre.

**Respostas Esperadas:**
- `200 OK`: O objeto que foi deletado nos milissegundos passados.
- `404 Not Found`: Se o ID fornecido nem existia em memória.
