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
