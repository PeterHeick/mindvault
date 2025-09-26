# MindVault 🧠

**A personal and collaborative knowledge engine, built for clarity and flow.**

MindVault is an open-source tool designed to combat cognitive overload. It acts as a central hub that can intelligently collect, index, and retrieve information from all your digital sources. The system is built on a modern, scalable architecture that supports both individual users and team collaboration.

---
## ✨ Core Features

* **Centralized Knowledge Base:** Gather all your notes, documents, and ideas in one secure place.
* **Intelligent Search:** A lightning-fast and typo-tolerant search experience powered by Typesense, so you can always find what you're looking for.
* **Automated Collection:** Use n8n to build workflows that automatically fetch information from emails, websites, other apps, and more.
* **API-First Design:** A robust REST API core that enables the creation of specialized clients (web, mobile, desktop).
* **Self-hosted & Open Source:** Full control over your own data. Run it on your own server and customize it to your needs.
* **Ready for Collaboration:** The architecture is designed to support shared knowledge spaces for teams and projects.

---
## 🚀 Technology Stack

MindVault is built with a modern and efficient technology stack:

* **Backend:** [Go (Golang)](https://go.dev/) with the [Gin](https://gin-gonic.com/) framework.
* **Database:** [PostgreSQL](https://www.postgresql.org/) for robust and relational data storage.
* **Search Engine:** [Typesense](https://typesense.org/) for a lightning-fast search experience.
* **Automation Engine:** [n8n](https://n8n.io/) for building integration workflows.
* **Frontend (Planned):** [Vue 3](https://vuejs.org/) for a modern and reactive user interface.
* **Containerization:** Everything is orchestrated with [Docker](https://www.docker.com/) and Docker Compose.



---
## 🏁 Getting Started

Follow these steps to get MindVault up and running locally.

### Prerequisites

* [Docker](https://www.docker.com/) and Docker Compose
* [Go](https://go.dev/doc/install) (for local development and dependency management)
* [Git](https://git-scm.com/)
* The `migrate` tool (see installation guide below)

### Installation

1.  **Clone the project:**
    ```sh
    git clone [https://github.com/your-username/mindvault.git](https://github.com/your-username/mindvault.git)
    cd mindvault
    ```

2.  **Configure your environment:**
    Rename `.env.example` to `.env` (or create a new `.env` file) and fill in the necessary variables, especially `POSTGRES_PASSWORD` and `TYPESENSE_API_KEY`.

3.  **Start all services:**
    This command builds the backend image and starts all containers.
    ```sh
    docker compose up --build -d
    ```

4.  **Run database migrations:**
    This creates the necessary tables in your database.
    ```sh
    # Install the migrate tool if you don't have it
    go install -tags 'postgres' [github.com/golang-migrate/migrate/v4/cmd/migrate@latest](https://github.com/golang-migrate/migrate/v4/cmd/migrate@latest)

    # Navigate to the backend directory
    cd backend

    # Run the migration (remember to replace with your password)
    migrate -path db/migrations -database 'postgres://mindvault:YOUR_PASSWORD@localhost:5432/mindvault_db?sslmode=disable' up
    ```

Your MindVault API is now available at `http://localhost:8080`.

---
## API Endpoints

* `POST /api/v1/notes`: Creates a new note.
* `GET /api/v1/notes`: Retrieves all notes.
* `GET /api/v1/notes/:id`: Retrieves a specific note.
* `PUT /api/v1/notes/:id`: Updates a note.
* `DELETE /api/v1/notes/:id`: Deletes a note.
* `GET /ping`: A simple health check that returns `{"message":"pong"}`.

---
## 🛣️ Project Status & Roadmap

The project is under active development.

* [x] **Phase 1: Core Foundation** - Setup of Docker, Go backend with a complete CRUD API for notes, database connection, and migration setup.
* [ ] **Phase 2: User Interface** - Development of a Vue 3 frontend to interact with the API.
* [ ] **Phase 3: Intelligent Collectors** - Integration with n8n to build the first automated data workflows.
* [ ] **Phase 4: Search Integration** - Activation of Typesense to provide a superior search experience.