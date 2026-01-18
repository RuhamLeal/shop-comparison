# Product Comparison API

API REST para gerenciamento e comparação de produtos com especificações técnicas detalhadas.

## Tecnologias Utilizadas

| Tecnologia | Versão | Descrição |
|------------|--------|-----------|
| **Go** | 1.25.0 | Linguagem de programação |
| **Fiber** | v3 | Framework web de alta performance |
| **SQLite** | 3 | Banco de dados embarcado |
| **sqlc** | 1.29.0 | Gerador de queries SQL type-safe |
| **Goose** | v3 | Ferramenta de migrations |
| **Swagger** | - | Documentação automática da API |
| **Air** | 1.62.0 | Hot reload para desenvolvimento |

## Arquitetura

O projeto segue os princípios de **Domain-Driven Design (DDD)** com separação em três camadas:

```
┌─────────────────────────────────────────────────────────────┐
│                    Camada de Infraestrutura                 │
│  (Fiber Handlers, SQLite Repository, Config)                │
├─────────────────────────────────────────────────────────────┤
│                    Camada de Aplicação                      │
│  (Use Cases, DTOs, Services)                                │
├─────────────────────────────────────────────────────────────┤
│                    Camada de Domínio                        │
│  (Entities, Repository Interfaces, Domain Services)         │
└─────────────────────────────────────────────────────────────┘
```

### Padrões Implementados

- **Repository Pattern**: Abstração da persistência de dados
- **Use Case Pattern**: Cada operação de negócio em uma classe isolada
- **DTO Pattern**: Objetos de transferência para entrada/saída da API
- **Factory Methods**: Criação de entidades com validação
- **Value Objects**: Tipos especializados para maior segurança

## Estrutura de Pastas

```
.
├── cmd/
│   └── api/
│       └── main.go              # Entry point da aplicação
├── internal/
│   ├── application/
│   │   ├── dto/                 # Data Transfer Objects
│   │   ├── services/            # Serviços de aplicação
│   │   └── usecase/             # Casos de uso
│   ├── domain/
│   │   ├── constants/           # Constantes de domínio
│   │   ├── entity/              # Entidades do domínio
│   │   ├── exception/           # Exceções customizadas
│   │   ├── repository/          # Interfaces de repositório
│   │   ├── services/            # Serviços de domínio
│   │   └── types/               # Type aliases
│   └── infra/
│       ├── config/              # Configurações
│       ├── fiber/               # Handlers, routes, middlewares
│       └── sqlite/              # Repositórios, queries, migrations
├── pkg/
│   └── validator/               # Validadores customizados
├── test/                        # Testes unitários
├── docs/                        # Documentação Swagger gerada
├── Makefile                     # Scripts de automação
├── go.mod                       # Dependências
└── sqlc.yaml                    # Configuração do sqlc
```

## Pré-requisitos

- Go 1.25.0 ou superior
- SQLite 3
- Make (opcional, para usar os scripts do Makefile)

## Configuração

### 1. Clone o repositório

```bash
git clone <repository-url>
cd <project-folder>
```

### 2. Configure as variáveis de ambiente

Crie um arquivo `.env` na raiz do projeto baseado no `.env.example`:

```bash
cp .env.example .env
```

Edite o arquivo `.env` com suas configurações:

```properties
# Banco de Dados
SQLITE_PATH=./project.db
SQLITE_BUSY_TIMEOUT=1000

# Servidor
FIBER_HOST=localhost
FIBER_PORT=8085
FIBER_DEBUG=true
FIBER_PREFORK=false

# Swagger (autenticação básica)
SWAGGER_ROUTE_ACCESS_USER="admin"
SWAGGER_ROUTE_ACCESS_PASSWORD="sua-senha-base64"
```

### 3. Instale as dependências

```bash
# Desenvolvimento
make dev-dependencies

# Produção
make prod-dependencies
```

## Executando o Projeto

### Desenvolvimento (com hot reload)

```bash
# Setup inicial (executa migrations)
make dev-setup

# Iniciar servidor
make dev
```

### Produção

```bash
# Setup e build
make prod-setup
make prod-build

# Executar binário
./bin/api
```

### Build manual

```bash
# Build de desenvolvimento
make dev-build

# Build de produção (otimizado)
make prod-build
```

## Comandos do Makefile

| Comando | Descrição |
|---------|-----------|
| `make dev` | Inicia o servidor com hot reload |
| `make dev-setup` | Executa migrations de desenvolvimento |
| `make dev-build` | Build com símbolos de debug |
| `make dev-dependencies` | Instala dependências de dev |
| `make prod` | Build completo de produção |
| `make prod-setup` | Setup de produção |
| `make prod-build` | Build otimizado |
| `make prod-dependencies` | Instala dependências de produção |
| `make migration-up` | Executa migrations pendentes |
| `make migration-down` | Desfaz última migration |
| `make migration-reset` | Reseta o banco de dados |
| `make create-migration name="nome"` | Cria nova migration |
| `make sqlc` | Gera código das queries SQL |
| `make swag` | Gera documentação Swagger |
| `make clean` | Remove binários gerados |

## Endpoints da API

### Categorias

| Método | Endpoint | Descrição |
|--------|----------|-----------|
| GET | `/categories` | Lista todas as categorias |

### Produtos

| Método | Endpoint | Descrição |
|--------|----------|-----------|
| POST | `/products` | Cria um novo produto |
| GET | `/products` | Lista produtos (paginado) |
| GET | `/products/:public_id` | Obtém um produto |
| PUT | `/products/:public_id` | Atualiza um produto |
| DELETE | `/products/:public_id` | Remove um produto |
| GET | `/products/:public_id/specifications` | Produto com especificações |
| POST | `/products/compare` | Compara dois produtos |
| GET | `/categories/:category_public_id/products` | Produtos por categoria |

### Especificações

| Método | Endpoint | Descrição |
|--------|----------|-----------|
| GET | `/specifications` | Lista especificações |
| GET | `/specification-groups` | Lista grupos de especificações |
| POST | `/product-specifications` | Associa especificação a produto |

### Documentação

| Método | Endpoint | Descrição |
|--------|----------|-----------|
| GET | `/swagger/*` | Interface Swagger UI |

## Documentação Swagger

Após iniciar o servidor, acesse a documentação interativa:

```
http://localhost:8085/swagger/
```

**Credenciais padrão:**
- Usuário: `admin`
- Senha: configurada em `SWAGGER_ROUTE_ACCESS_PASSWORD`

## Entidades do Domínio

### Product (Produto)

Representa um produto com suas características básicas:
- Identificador público (8 caracteres)
- Nome, descrição, preço
- Avaliação (rating de 0-5 estrelas)
- Categoria associada
- Valores de especificações

### Category (Categoria)

Agrupamento de produtos por tipo/categoria.

### Specification (Especificação)

Define tipos de especificações técnicas disponíveis:
- PowerInWatts, ConsumptionKwh, CapacityLiters
- FrequencyMHz, FrequencyGHz, Threads, TDPWatts
- USBC, Waterproof (booleanos)
- NoiseDb, CaloriesKcal
- WidthCm, HeightCm, DepthCm, WeightKg, VolumeLiters

### SpecificationGroup (Grupo de Especificações)

Agrupa especificações relacionadas.

### ProductSpecificationValue (Valor de Especificação)

Associação entre produto e especificação com valor concreto (string, int ou bool).

### Insight

Resultado da comparação entre produtos, indicando:
- Se é favorável, neutro ou desfavorável
- Mensagem descritiva da comparação

## Funcionalidade de Comparação

A API permite comparar dois produtos, gerando insights sobre:

1. **Preço**: Diferença absoluta e percentual
2. **Avaliação**: Diferença de rating
3. **Especificações**: Comparação tipo-específica

Exemplo de uso:
```bash
POST /products/compare
{
  "product_a_public_id": "abc12345",
  "product_b_public_id": "xyz67890"
}
```

## Banco de Dados

O projeto utiliza SQLite com as seguintes tabelas:

- `categories` - Categorias de produtos
- `products` - Produtos
- `specification_groups` - Grupos de especificações
- `specifications` - Especificações disponíveis
- `product_specifications` - Valores de especificações por produto

As queries SQL são geradas automaticamente pelo **sqlc**, garantindo type-safety em tempo de compilação.

## Testes

Execute os testes unitários:

```bash
go test ./test/...
```

Os testes cobrem principalmente as entidades de domínio e suas validações.

## Desenvolvimento

### Gerar queries após modificar SQL

```bash
make sqlc
```

### Atualizar documentação Swagger

```bash
make swag
```

### Criar nova migration

```bash
make create-migration name="add_new_table"
```

## Licença

[Adicione sua licença aqui]
