# GoREST 🚀

**Because everything is ready!**

A RESTful API boilerplate built with Go, featuring:

✅ Authentication & Registration

✅ Clean Architecture with Repository Pattern

✅ Ready-to-use endpoints with tests

✅ Environment configuration

✅ Database migrations

✅ CLI tool for migrations

*Read more about [the creation of GoREST in this article](https://mostafaqanbaryan.com/how-to-write-a-backend-the-worst-way-creation-of-go-rest/).*

---

⚠️ **Attention**: This project is still in development and not ready for production use.

---

## ✨ Features

- **Framework**: 🏗️ [Echo](https://echo.labstack.com/) for high-performance routing
- **Database**: 🗄️ [SQLC](https://sqlc.dev/) for type-safe SQL queries
- **Migrations**: 🏃‍♂️ Goose with CLI support
- **Validation**: 🔍 Validating every request
- **Auth**: 🔑 Cache-based authentication
- **Testing**: 🧪 Comprehensive handler tests
- **Config**: ⚙️ .env file support
- **Structure**: 🏛️ Clean architecture with separated layers


## ⚙️ Installation

1. **Clone the repository**:
   ```bash
   git clone https://github.com/mostafaqanbaryan/go-rest.git
   cd go-rest
   ```

2. **Install dependencies**:
   ```bash
   go mod download
   ```

3. **Set up environment**:
   ```bash
   cp .env.example .env

   ```
   Edit the `.env` file with your configuration.


## 🚀 Usage

### Running the Server (For development)

This will work with air and rebuilds the image every time `go.mod` changes.

```bash
docker compose watch
```

### Using the CLI
```bash
# Run migrations (using goose)
go run cmd/cli/main.go migrate up

# Create new migration (using goose)
go run cmd/cli/main.go create migration_name (go|sql)
```

### Available Endpoints
| Method | Endpoint             | Description                  |
|--------|----------------------|------------------------------|
| POST   | `/auth/register`     | User registration            |
| POST   | `/auth/login`        | User login                   |
| GET    | `/auth/logout`       | Logout user                  |
| GET    | `/me`                | Get current user profile     |
| PATCH  | `/me`                | Update current user fullname |


## 🧪 Testing
Run all tests:
```bash
go test ./...
```

## 📝 TODO Roadmap

Here's what's coming next to make GoREST even better:

- [ ] Add Swagger/OpenAPI documentation
- [ ] Implement JWT token generation/validation
- [ ] Add Authorization (with Casbin?)
- [ ] Add Authentication middleware
- [ ] Add a Rate Limiting mechanism
- [ ] Implement better logging (Zap or Logrus?)
- [ ] Integrate Ofelia for scheduled tasks
- [ ] Add gRPC layer for microservices
- [ ] Add more RESTful endpoints

## 📄 License
MIT
