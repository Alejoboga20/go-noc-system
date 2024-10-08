# Go NOC System

A simple **Network Operations Center (NOC) system** built in Go using **Clean Architecture** principles. This system monitors web services (e.g., pings a web page) and writes logs into various data sources such as a file system or databases.

## Features

- **Ping Web Services**: Monitor the status and response time of web services.
- **Logging**: Store logs in multiple backends, including the file system, databases (PostgreSQL, MongoDB), etc.
- **Clean Architecture**: Modular and scalable project structure with separation of concerns.
- **Retry Mechanism**: Automatic retries for failed pings with configurable thresholds.
- **CLI and Web Interface** (Optional): View logs and service status through a simple CLI or web interface.
- **Notifications** (Optional): Alerts users of service failures via email, SMS, or messaging services.

---

## Project Structure

The project follows **Clean Architecture**, ensuring a scalable and testable codebase:

- Entry point of the application `/domain`
- Core business logic (Entities) `/entities`
- Domain entities (Service, PingResult) `/repository`
- Interfaces for data sources `/usecase`
- Application logic (PingService, LogResults) `/infrastructure`
- External systems (HTTP clients, databases, file system) `/interfaces`
- API layers (CLI, HTTP routes) `/tests`

## Requirements

- **Go** (>= 1.20)
- **PostgreSQL** or **MongoDB** (optional, for database logging)
- **Docker** (for containerization and deployment)
- **Golang-migrate** (optional, for database migrations)

## Setup

1. Clone the Repository.
   ```bash
   git clone https://github.com/yourusername/go-noc-system.git
   cd go-noc-system
   ```
2. Install Dependencies: Ensure you have Go installed. Then run:
   ```bash
   go mod tidy
   ```
3. Set Environment Variables: Create a `.env` file in the root directory and add the following:
   ```env
   PING_INTERVAL=
   FILE_PATH=
   ```
4. Run the Application:
   ```bash
    go run main.go
   ```

## Contributing

Contributions are welcome! Feel free to submit a pull request or open an issue.

## License

This project is licensed under the MIT License

## Contact

For any questions or issues, feel free to reach out:

alejoboga19@gmail.com
GitHub: @alejoboga20
