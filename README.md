# Groq API Proxy - Go Backend

This Go application serves as a simple proxy for interacting with the Groq API. It forwards requests to the Groq API for chat completions and returns the responses to the client. The app also includes basic CORS support, environment configuration via `.env` file, and a health check endpoint.

The application is built using the `chi` router and can be run as a standalone Go application or containerized using Docker.

---

## Features

- **API Proxy**: Forwards chat completion requests to the Groq API and returns the results.
- **CORS Support**: Configures cross-origin resource sharing (CORS) to allow requests from different domains.
- **Health Check**: Exposes a `/healthz` endpoint for server readiness checks.
- **Environment Configuration**: Loads environment variables like the Groq API key from a `.env` file.

---

## Requirements

- **Docker** (Optional): Use Docker to containerize and run the application.
- **.env File**: This file should contain environment variables such as the API key for Groq.

---

## Getting Started

## Running the Application

1. Clone the repository:

   ```bash
   git clone https://github.com/indenigrate/iykyk.git
   cd iykyk
2. Build and run:

   ```bash
   docker compose build
   docker compose up
3. Make the .env file as per .env.example
4. Access the API
        http://localhost:5000/healthz.
        http://localhost:5000/enter your request argument.

## Basic Usage
### API endpoints
- **GET `/enter your request here`**
  - **Description:** Returns a response after processing the request
  - **Response:** Returns a response text.
