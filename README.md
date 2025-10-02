<div align="center">

<img width="100%" src="https://capsule-render.vercel.app/api?type=waving&color=00ADD8&height=200&section=header&text=Task%20Tracker&fontSize=50&fontColor=fff&animation=twinkling&fontAlignY=40&desc=Go%20|%20CLI%20|%20JSON%20|%20Task%20Management&descAlignY=60&descSize=18">

<p align="center">
  <i>✅ A simple and efficient command-line task tracker application built with Go, featuring JSON-based storage and intuitive task management.</i>
</p>

<p align="center">
  <i>✅ Uma aplicação de linha de comando simples e eficiente para rastreamento de tarefas construída com Go, com armazenamento baseado em JSON e gerenciamento intuitivo de tarefas.</i>
</p>

---

### 📚 Practice Project | Projeto de Prática

<div align="center">

**Language:** Go (Golang) | **Linguagem:** Go (Golang)  
**Objective:** Learn Go fundamentals through a practical CLI application | **Objetivo:** Aprender os fundamentos de Go através de uma aplicação CLI prática

</div>

### 🌟 Features | Funcionalidades

<div align="center">

| Feature | Description | Descrição |
|:---:|:---|:---|
| ➕ | Add Tasks | Adicionar Tarefas |
| 📋 | List All Tasks | Listar Todas as Tarefas |
| ✏️ | Update Task Status | Atualizar Status de Tarefas |
| 🗑️ | Delete Tasks | Deletar Tarefas |
| 💾 | JSON Storage | Armazenamento em JSON |
| ⚡ | Fast CLI Interface | Interface CLI Rápida |
| 🔄 | Status Tracking | Rastreamento de Status |
| 📅 | Timestamp Management | Gerenciamento de Timestamps |

</div>

### 🛠️ Technologies | Tecnologias

<div align="center">
  <a href="https://skillicons.dev">
    <img src="https://skillicons.dev/icons?i=go&theme=dark" />
  </a>
</div>

### 🏗️ Architecture | Arquitetura

```
┌─────────────────────────────┐
│   Task Tracker CLI App      │
│   (Command Line Interface)  │
└──────────────┬──────────────┘
               │
               ▼
┌─────────────────────────────┐
│   JSON File Storage         │
│   (tasks.json)              │
└─────────────────────────────┘
```

### 📋 Task Status | Status de Tarefas

- **todo**: Task not started yet | Tarefa ainda não iniciada
- **in-progress**: Task currently being worked on | Tarefa em andamento
- **done**: Task completed | Tarefa concluída

### 🚀 Getting Started | Começando

#### Prerequisites | Pré-requisitos

- Go 1.25.1 or higher installed | Go 1.25.1 ou superior instalado
- Basic command line knowledge | Conhecimento básico de linha de comando

#### Installation | Instalação

```bash
# Clone the repository | Clone o repositório
git clone https://github.com/matheussricardoo/TaskTracker.git

# Navigate to project directory | Navegue até o diretório do projeto
cd TaskTracker

# Build the application | Compile a aplicação
go build -o taskTracker.exe main.go

# Or run directly | Ou execute diretamente
go run main.go
```

#### Usage | Uso

```bash
# Show help | Mostrar ajuda
./taskTracker.exe

# Add a new task | Adicionar uma nova tarefa
./taskTracker.exe add "Complete the project documentation"

# List all tasks | Listar todas as tarefas
./taskTracker.exe list

# Update task status | Atualizar status da tarefa
./taskTracker.exe update 1 in-progress
./taskTracker.exe update 1 done

# Delete a task | Deletar uma tarefa
./taskTrackerGo.exe delete 1
```

### 📁 Project Structure | Estrutura do Projeto

```
TaskTracker/
├── main.go              # Main application file with all functionality
├── go.mod              # Go module definition
├── tasks.json          # JSON file for task storage (auto-generated)
├── taskTracker.exe   # Compiled executable
└── README.md           # Project documentation
```

### 📦 Task Structure | Estrutura de Tarefa

```json
{
  "id": 1,
  "description": "Complete the project documentation",
  "status": "todo",
  "createdAt": "2025-10-01T10:00:00Z",
  "updatedAt": "2025-10-01T10:00:00Z"
}
```

### 💡 Key Features Implementation | Implementação dos Recursos Principais

- **JSON Persistence:** Tasks are automatically saved to `tasks.json` | **Persistência JSON:** Tarefas são automaticamente salvas em `tasks.json`
- **Auto-incrementing IDs:** Each task gets a unique identifier | **IDs auto-incrementais:** Cada tarefa recebe um identificador único
- **Timestamp Tracking:** Creation and update times are recorded | **Rastreamento de Timestamps:** Horários de criação e atualização são registrados
- **Status Validation:** Only valid status values are accepted | **Validação de Status:** Apenas valores de status válidos são aceitos
- **Error Handling:** Comprehensive error messages for user guidance | **Tratamento de Erros:** Mensagens de erro abrangentes para orientação do usuário

### 🎯 Learning Objectives | Objetivos de Aprendizado

- [x] Go syntax and basic constructs | Sintaxe Go e construções básicas
- [x] File I/O operations in Go | Operações de I/O de arquivo em Go
- [x] JSON marshaling and unmarshaling | Serialização e desserialização JSON
- [x] Command-line argument parsing | Análise de argumentos de linha de comando
- [x] Struct definitions and methods | Definições de struct e métodos
- [x] Error handling patterns | Padrões de tratamento de erros
- [x] Time package usage | Uso do pacote time

### 📋 Available Commands | Comandos Disponíveis

| Command | Arguments | Description | Descrição |
|:---:|:---|:---|:---|
| `list` | - | Display all tasks | Exibir todas as tarefas |
| `add` | `"description"` | Add a new task | Adicionar uma nova tarefa |
| `update` | `<ID> <status>` | Update task status | Atualizar status da tarefa |
| `delete` | `<ID>` | Delete a task | Deletar uma tarefa |

### 👤 Author | Autor

<div align="center">
  <a href="https://github.com/matheussricardoo" target="_blank">
    <img src="https://skillicons.dev/icons?i=github" alt="GitHub"/>
  </a>
  <a href="https://www.linkedin.com/in/matheus-ricardo-426452266/" target="_blank">
    <img src="https://skillicons.dev/icons?i=linkedin" alt="LinkedIn"/>
  </a>
</div>

### 📜 License | Licença

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

<img width="100%" src="https://capsule-render.vercel.app/api?type=waving&color=00ADD8&height=120&section=footer"/>

</div>
