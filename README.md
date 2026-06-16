# Subscription API

A RESTful API for subscription management with user authentication and payment processing built with Go and Fiber.

## Features

- 🔐 **User Authentication**: Register and login with JWT tokens
- 👤 **User Management**: CRUD operations for users
- 📦 **Subscription Management**: Create and manage subscriptions
- 💳 **Payment Processing**: Simulated payment processing with transaction tracking
- 🗄️ **Database Migrations**: Auto-migrate database schemas
- 🚀 **Multi-Environment Support**: Development, staging, and production configurations
- 🛡️ **Security**: Password hashing with bcrypt, JWT authentication

## Tech Stack

- **Framework**: [Fiber](https://gofiber.io/) - Fast HTTP framework for Go
- **ORM**: [GORM](https://gorm.io/) - Object Relational Mapping
- **Database**: PostgreSQL
- **Authentication**: JWT (JSON Web Tokens)
- **Hashing**: bcrypt
- **Configuration**: JSON-based config files

## Project Structure

```
belajar-fiber/
├── config/                 # Configuration files and package
│   ├── config.go
│   ├── config.example.json
│   ├── config.development.json
│   ├── config.staging.json
│   └── config.production.json
├── database/               # Database connection and migrations
│   ├── database.go
│   └── migrate.go
├── handlers/               # Request handlers
│   ├── auth_handler.go
│   ├── user_handler.go
│   ├── subscription_handler.go
│   └── payment_handler.go
├── middleware/             # Middleware (auth, etc.)
│   └── auth.go
├── models/                 # Database models
│   ├── user.go
│   └── subscription.go
├── routes/                 # Route definitions
│   └── routes.go
├── .gitignore
├── go.mod
├── go.sum
├── main.go                 # Application entry point
├── Makefile                # Makefile for Linux/Mac
└── run.ps1                 # PowerShell script for Windows
```

## Getting Started

### Prerequisites

- Go 1.26 or higher
- PostgreSQL 13 or higher
- Git

### Installation

1. **Clone the repository**:
   ```bash
   git clone <repository-url>
   cd belajar-fiber
   ```

2. **Install dependencies**:
   ```bash
   go mod download
   ```

3. **Set up configuration**:
   ```bash
   # Copy example config to development config
   cp config/config.example.json config/config.development.json
   ```

4. **Update configuration**:
   Edit `config/config.development.json` with your database credentials:
   ```json
   {
     "app": {
       "name": "Subscription API",
       "port": "3000",
       "environment": "development"
     },
     "database": {
       "host": "localhost",
       "port": "5432",
       "user": "your_db_user",
       "password": "your_db_password",
       "name": "subrekDB_dev",
       "sslmode": "disable"
     },
     "jwt": {
       "secret": "your-super-secret-key-change-this-in-production",
       "expiry": "24h"
     }
   }
   ```

5. **Create database**:
   ```sql
   CREATE DATABASE subrekDB_dev;
   ```

6. **Run migrations**:
   ```bash
   # Windows
   .\run.ps1 migrate

   # Linux/Mac
   make migrate
   ```

7. **Start the server**:
   ```bash
   # Windows
   .\run.ps1 dev

   # Linux/Mac
   make dev
   ```

The server will start on `http://localhost:3000`

## Configuration

### Environment Files

The application uses JSON configuration files located in the `config/` directory:

- `config.example.json` - Example configuration (commit to Git)
- `config.development.json` - Development environment (do not commit)
- `config.staging.json` - Staging environment (do not commit)
- `config.production.json` - Production environment (do not commit)

### Configuration Options

| Section | Key | Description |
|---------|-----|-------------|
| `app` | `name` | Application name |
| `app` | `port` | Server port |
| `app` | `environment` | Environment name |
| `database` | `host` | Database host |
| `database` | `port` | Database port |
| `database` | `user` | Database user |
| `database` | `password` | Database password |
| `database` | `name` | Database name |
| `database` | `sslmode` | SSL mode (disable/require) |
| `jwt` | `secret` | JWT secret key |
| `jwt` | `expiry` | JWT token expiry (e.g., 24h, 7d) |

## Usage

### Available Commands

#### Windows (PowerShell)
```powershell
.\run.ps1 dev       # Start development server
.\run.ps1 stage     # Start staging server
.\run.ps1 prod      # Start production server
.\run.ps1 migrate   # Run database migrations
.\run.ps1 reset     # Reset database (drop all tables and migrate)
.\run.ps1 build     # Build the application
.\run.ps1 clean     # Remove binary files
.\run.ps1 help      # Show help
```

#### Linux/Mac (Makefile)
```bash
make dev            # Start development server
make stage          # Start staging server
make prod           # Start production server
make migrate        # Run database migrations
make reset          # Reset database (drop all tables and migrate)
make build          # Build the application
make clean          # Remove binary files
make help           # Show help
```

#### Direct Flags
```bash
# Run with specific environment
go run main.go --env=staging

# Only run migrations
go run main.go --env=development --migrate

# Reset database
go run main.go --env=development --reset
```

## API Documentation

### Base URL
```
http://localhost:3000/api
```

### Authentication Endpoints

#### Register User
```http
POST /auth/register
Content-Type: application/json

{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "password123"
}
```

#### Login User
```http
POST /auth/login
Content-Type: application/json

{
  "email": "john@example.com",
  "password": "password123"
}
```

**Response**:
```json
{
  "message": "Login successful",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com"
  }
}
```

### User Endpoints

#### Get All Users
```http
GET /users
```

#### Get User by ID
```http
GET /users/:id
```

#### Update User
```http
PUT /users/:id
Content-Type: application/json

{
  "name": "John Doe Updated",
  "email": "john.updated@example.com"
}
```

#### Delete User
```http
DELETE /users/:id
```

### Subscription Endpoints (Requires Auth)

All subscription endpoints require a valid JWT token in the `Authorization` header:
```
Authorization: Bearer <your-token>
```

#### Create Subscription
```http
POST /subscriptions
Content-Type: application/json

{
  "package_id": "premium",
  "status": "pending",
  "amount": 99.99
}
```

#### Get User's Subscriptions
```http
GET /subscriptions
```

#### Get All Subscriptions
```http
GET /subscriptions/all
```

### Payment Endpoints (Requires Auth)

#### Process Payment
```http
POST /payments
Content-Type: application/json

{
  "subscription_id": 1,
  "method": "credit_card"
}
```

#### Get User's Payments
```http
GET /payments
```

#### Get Payment Status
```http
GET /payments/:transaction_id
```

### Health Check
```http
GET /
```

**Response**:
```json
{
  "message": "Welcome to Subscription API",
  "version": "1.0.0",
  "env": "development"
}
```

## Database Models

### User
```go
type User struct {
  gorm.Model
  Name     string `json:"name"`
  Email    string `json:"email" gorm:"uniqueIndex"`
  Password string `json:"-"`
}
```

### Subscription
```go
type Subscription struct {
  gorm.Model
  UserID    uint    `json:"user_id"`
  PackageID string  `json:"package_id"`
  Status    string  `json:"status"`
  Amount    float64 `json:"amount"`
}
```

### Payment
```go
type Payment struct {
  gorm.Model
  SubscriptionID uint   `json:"subscription_id"`
  Method         string `json:"method"`
  Status         string `json:"status"`
  TransactionID  string `json:"transaction_id"`
}
```

## Security

- **Passwords**: Hashed using bcrypt before storage
- **JWT Tokens**: Used for stateless authentication
- **Environment Variables**: Sensitive data stored in config files (not committed to Git)
- **CORS**: Enabled for cross-origin requests

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Support

For support, email your-email@example.com or create an issue in the repository.

---

**Happy Coding!** 🚀
