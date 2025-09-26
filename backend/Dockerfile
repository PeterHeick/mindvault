# backend/Dockerfile

# --- Fase 1: Build ---
# Brug et Go-image til at bygge applikationen
FROM golang:1.24-alpine AS builder

# Sæt working directory
WORKDIR /app

# Kopier go.mod og go.sum for at hente afhængigheder
COPY go.mod ./
# TÆND for go modules og hent dependencies
RUN go mod download

# Kopier resten af kildekoden
COPY . .

# Byg applikationen.
# CGO_ENABLED=0 er vigtigt for at bygge en statisk binær fil, der virker i et minimalt image.
# -o /server bygger outputtet til en fil ved navn 'server' i roden.
RUN CGO_ENABLED=0 go build -o /server .

# --- Fase 2: Final ---
# Brug et minimalt image til det endelige produkt for at reducere størrelse og angrebsflade
FROM alpine:latest

# Kopier den byggede binære fil fra 'builder'-fasen
COPY --from=builder /server /server

# Eksponér den port, som vores applikation lytter på (fra .env-filen)
EXPOSE 8080

# Kommandoen til at køre applikationen, når containeren starter
CMD [ "/server" ]