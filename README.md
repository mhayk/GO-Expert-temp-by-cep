<p align="center">
	<img src="https://cdn-icons-png.flaticon.com/512/6218/6218295.png" width="100" alt="Go weather API icon" />
</p>

<h1 align="center">GO Expert - Temperature by Brazilian ZIP Code (CEP)</h1>

<p align="center">
	A production-oriented Go microservice that validates Brazilian CEP input, orchestrates
	multiple external integrations (ViaCEP and WeatherAPI), and exposes a clear HTTP contract
	for temperature conversion in Celsius, Fahrenheit, and Kelvin.
</p>

<p align="center">
	<img src="https://img.shields.io/github/last-commit/mhayk/GO-Expert-temp-by-cep?style=flat&logo=git&logoColor=white&color=0080ff" alt="last-commit">
	<img src="https://img.shields.io/github/languages/top/mhayk/GO-Expert-temp-by-cep?style=flat&color=0080ff" alt="repo-top-language">
	<img src="https://img.shields.io/github/languages/count/mhayk/GO-Expert-temp-by-cep?style=flat&color=0080ff" alt="repo-language-count">
	<img src="https://img.shields.io/badge/Go-1.22+-00ADD8.svg?style=flat&logo=Go&logoColor=white" alt="Go">
	<img src="https://img.shields.io/badge/Docker-Enabled-2496ED.svg?style=flat&logo=Docker&logoColor=white" alt="Docker">
</p>

---

## Engineering Highlights

This service was designed to demonstrate core backend concerns beyond basic CRUD:

- Contract-first HTTP behaviour with deterministic status codes (`200`, `422`, `404`, `500`).
- Layered design separating transport logic from external provider integrations.
- Input validation and explicit failure-path handling.
- Integration testing with upstream HTTP mocking to ensure repeatability.
- Container-ready workflows for local development and production packaging.

For technical evaluation, this project evidences competency in API design, boundary management, external dependency integration, and service reliability in Go.

## Execution Flow

Given a `zipcode`, the request pipeline is:

1. Validate CEP format using `brdoc`.
2. Resolve location metadata via ViaCEP.
3. Extract city and request current weather from WeatherAPI.
4. Compute Kelvin from Celsius and serialise response JSON.

Returned fields:

- `temp_C`
- `temp_F`
- `temp_K`

## Architecture

```text
main.go
	-> api/handler/weather.go
			-> integration/address/address.go   (ViaCEP)
			-> integration/weather/weather.go   (WeatherAPI)
```

High-level structure:

```text
.
|- api/handler/weather.go
|- integration/address/address.go
|- integration/weather/weather.go
|- test/integration/weather_test.go
|- main.go
|- docker-compose.yml
|- Dockerfile.dev
|- Dockerfile.prod
```

## API Endpoint

### Request

`GET /temp-by-cep?zipcode={CEP}`

Example:

```bash
curl "http://localhost:8080/temp-by-cep?zipcode=69304350"
```

### Success Response (`200`)

```json
{
	"temp_C": 20,
	"temp_F": 68,
	"temp_K": 293
}
```

### Error Responses

- `422 Unprocessable Entity`: `invalid zipcode`
- `404 Not Found`: `can not find zipcode`
- `500 Internal Server Error`: upstream/integration failure

## Design Decisions

- Standard library HTTP stack (`net/http`) to keep runtime and dependency surface minimal.
- `http.NewServeMux` for clear and explicit routing.
- Dedicated integration packages (`integration/address`, `integration/weather`) to isolate provider concerns.
- Environment-driven configuration (`WEATHER_API_KEY`, `ENVIRONMENT`) to support multiple runtime contexts.

## Trade-offs and Rationale

- Simplicity over framework abstraction:
	Using the Go standard library keeps startup time, dependency count, and maintenance overhead low.
- Fast delivery over advanced resiliency primitives:
	The current version prioritises clarity and correctness; retries, circuit breaking, and timeout policies are planned improvements.
- Readability over premature optimisation:
	The request pipeline is intentionally explicit to make behaviour easy to reason about and review.
- Strong API contract over silent fallback behaviour:
	Distinct status codes and clear error messages make failure modes predictable for API consumers.

## Testing Strategy

The suite under `test/integration` validates both happy path and failure paths at HTTP boundary level:

- Success scenario with complete CEP -> city -> weather chain.
- Invalid CEP input validation.
- Unknown CEP lookup behaviour.
- Weather provider failure propagation.

External calls are mocked with `gock`, ensuring deterministic and fast feedback loops during CI and local runs.

## Running Locally

### 1. Clone the repository

```bash
git clone https://github.com/mhayk/GO-Expert-temp-by-cep
cd GO-Expert-temp-by-cep
```

### 2. Configure environment variables

```bash
cp .env.example .env
```

Set your WeatherAPI key in `.env`:

```env
WEATHER_API_KEY=your_weather_api_key
PORT=8080
ENVIRONMENT=development
```

### 3. Install dependencies

```bash
go mod tidy
```

### 4. Run with Go

```bash
go run main.go
```

The API will be available at `http://localhost:8080`.

## Local Verification

Smoke test:

```bash
curl -i "http://localhost:8080/temp-by-cep?zipcode=69304350"
```

## Running with Docker (Development)

```bash
docker compose up -d
docker compose exec goapp sh
go run main.go
```

## Tests

Run all tests:

```bash
go test ./...
```

The integration tests cover:

- valid CEP -> correct temperature conversion response
- invalid CEP -> `422`
- non-existent CEP -> `404`
- external weather provider failure -> `500`

## Production Build

The repository includes a multi-stage production image in `Dockerfile.prod`.

Build and run:

```bash
docker build -f Dockerfile.prod -t temp-by-cep:prod .
docker run --rm -p 8080:8080 --env WEATHER_API_KEY=your_weather_api_key temp-by-cep:prod
```

## Operational Notes

- In non-production mode, `.env` is loaded automatically.
- In production mode, environment variables must be injected externally.
- The server listens on port `8080`.

## Scalability and Reliability Approach

- Horizontal scale:
	The service is stateless and suitable for scale-out behind a load balancer.
- Throughput model:
	Runtime bottlenecks are primarily external I/O (ViaCEP/WeatherAPI), not CPU.
- Production hardening roadmap:
	Add per-request context timeouts, bounded retries with jitter, and circuit breaker behaviour.
- Resilience posture:
	Keep upstream concerns isolated in integration packages so resiliency policies can be added without affecting handlers.

## Incident Handling Model

- Detection:
	Monitor request error rates, latency percentiles, and upstream provider failure spikes.
- Triage:
	Classify incidents by category (input validation, ViaCEP dependency, WeatherAPI dependency, platform/runtime).
- Mitigation:
	Degrade gracefully for upstream failures with explicit `500` responses and actionable logs.
- Recovery:
	Use deployment rollback and feature toggles (where available) to restore stable behaviour quickly.
- Prevention:
	Convert incident learnings into tests and guardrails (timeouts, retries, observability, alerting thresholds).

## Interview Discussion Points

If asked to evolve this service in a senior backend interview, I would prioritise:

1. Request-scoped timeouts and transport-level tuning in HTTP clients.
2. Structured logs with correlation IDs and request metadata.
3. Metrics and traces (latency, dependency error rates, saturation signals).
4. Explicit service-level objectives and alert strategy.
5. Cache strategy for repeated CEP lookups to reduce upstream dependency pressure.

## Known Improvements (Next Iteration)

- Add request timeouts and retry/backoff policies for upstream calls.
- Introduce structured logging and correlation IDs.
- Add contract tests and benchmark tests.
- Add observability primitives (metrics/tracing) for production diagnostics.

## Tech Stack

- Go
- Docker / Docker Compose
- ViaCEP API
- WeatherAPI
- Testify + Gock

---

Built by Mhayk Whandson as part of the GO Expert postgraduate challenge, with focus on maintainable service design, integration robustness, and production-oriented Go engineering.
