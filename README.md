# MindVault 🧠

**En personlig og kollaborativ vidensmotor, bygget til at skabe klarhed og flow.**

MindVault er et open source-værktøj designet til at bekæmpe kognitiv overbelastning. Det fungerer som en central hub, der intelligent kan indsamle, indeksere og genfinde information fra alle dine digitale kilder. Systemet er bygget på en moderne, skalerbar arkitektur, der understøtter både den enkelte bruger og teamsamarbejde.

---
## ✨ Kernefunktioner

* **Centraliseret Vidensbase:** Saml alle dine noter, dokumenter og ideer ét sikkert sted.
* **Intelligent Søgning:** En lynhurtig og fejltolerant søgeoplevelse drevet af Typesense, så du altid kan finde, hvad du leder efter.
* **Automatiseret Indsamling:** Brug n8n til at bygge workflows, der automatisk henter information fra e-mails, hjemmesider, andre apps og meget mere.
* **API-Først Design:** En robust REST API-kerne, der gør det muligt at bygge specialiserede klienter (web, mobil, desktop).
* **Selv-hostet & Open Source:** Fuld kontrol over dine egne data. Kør det på din egen server og tilpas det efter dine behov.
* **Klar til Samarbejde:** Arkitekturen er forberedt til at understøtte delte vidensrum for teams og projekter.

---
## 🚀 Teknologi-stak

MindVault er bygget med en moderne og effektiv teknologi-stak:

* **Backend:** [Go (Golang)](https://go.dev/) med [Gin](https://gin-gonic.com/) frameworket.
* **Database:** [PostgreSQL](https://www.postgresql.org/) for robust og relationel datalagring.
* **Søgemaskine:** [Typesense](https://typesense.org/) for en lynhurtig søgeoplevelse.
* **Automations-motor:** [n8n](https://n8n.io/) til at bygge integrations-workflows.
* **Frontend (Planlagt):** [Vue 3](https://vuejs.org/) for en moderne og reaktiv brugerflade.
* **Containerisering:** Alt er orkestreret med [Docker](https://www.docker.com/) og Docker Compose.



---
## 🏁 Kom i gang

Følg disse trin for at få MindVault op at køre lokalt.

### Forudsætninger

* [Docker](https://www.docker.com/) og Docker Compose
* [Go](https://go.dev/doc/install) (til lokal udvikling og dependency management)
* [Git](https://git-scm.com/)
* `migrate` værktøjet (se installations-guide nedenfor)

### Installation

1.  **Klon projektet:**
    ```sh
    git clone [https://github.com/dit-brugernavn/mindvault.git](https://github.com/dit-brugernavn/mindvault.git)
    cd mindvault
    ```

2.  **Konfigurer dit miljø:**
    Omdøb `.env.example` til `.env` (eller opret en ny `.env`-fil) og udfyld de nødvendige variabler, især `POSTGRES_PASSWORD` og `TYPESENSE_API_KEY`.

3.  **Start alle services:**
    Denne kommando bygger backend-imaget og starter alle containere.
    ```sh
    docker compose up --build -d
    ```

4.  **Kør database-migrationer:**
    Dette opretter de nødvendige tabeller i din database.
    ```sh
    # Installer migrate-værktøjet, hvis du ikke har det
    go install -tags 'postgres' [github.com/golang-migrate/migrate/v4/cmd/migrate@latest](https://github.com/golang-migrate/migrate/v4/cmd/migrate@latest)

    # Naviger til backend-mappen
    cd backend

    # Kør migrationen (husk at erstatte med dit kodeord)
    migrate -path db/migrations -database 'postgres://mindvault:DIT_KODEORD@localhost:5432/mindvault_db?sslmode=disable' up
    ```

Dit MindVault API er nu tilgængeligt på `http://localhost:8080`.

---
## Основні кінцеві точки API

* `POST /api/v1/notes`: Opretter en ny note.
* `GET /api/v1/notes`: Henter alle noter.
* `GET /api/v1/notes/:id`: Henter en specifik note.
* `PUT /api/v1/notes/:id`: Opdaterer en note.
* `DELETE /api/v1/notes/:id`: Sletter en note.
* `GET /ping`: En simpel health-check, der returnerer `{"message":"pong"}`.

---
## 🛣️ Projektstatus & Køreplan

Projektet er i aktiv udvikling.

* [x] **Fase 1: Kernens Fundament** - Opsætning af Docker, Go-backend med komplet CRUD API til noter, databaseforbindelse og migrations-setup.
* [ ] **Fase 2: Brugerflade** - Udvikling af en Vue 3 frontend til at interagere med API'et.
* [ ] **Fase 3: Intelligente Indsamlere** - Integration med n8n for at bygge de første automatiserede data-workflows.
* [ ] **Fase 4: Søge-integration** - Aktivering af Typesense for at levere en overlegen søgeoplevelse.