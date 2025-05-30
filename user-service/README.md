Here's the `README.md` in English, tailored for your `user-service` microservice:

---

````markdown
# 👤 User Service — MooDu Microservice

A microservice responsible for user management within the MooDu platform: registration, authentication, and role-based profiles for students and instructors.

---

## 🚀 Features

- gRPC API for user operations
- User registration and login
- Support for student and instructor profiles
- JWT-based authentication
- Redis caching of profiles
- MongoDB as primary database
- NATS integration for event publishing
- Clean architecture (adapter → usecase → domain layers)

---

## 🛠️ Tech Stack

- Go (1.21+)
- MongoDB
- Redis
- NATS
- gRPC
- JWT
- Postman (for testing)
- Docker (optional)

---

## 📦 Environment Variables (.env)

```env
VERSION=1.0.0

# gRPC
GRPC_PORT=50051

# MongoDB
MONGO_DB=moodu_users
MONGO_DB_URI=mongodb://localhost:27017
MONGO_USERNAME=empty
MONGO_PWD=empty
MONGO_DB_REPLICA_SET=rs0
MONGO_WRITE_CONCERN=majority
MONGO_TLS_FILE_PATH=/path/to/cert.pem
MONGO_TLS_ENABLE=false

# Redis
REDIS_HOST=localhost:6379
REDIS_DB=0

# NATS
NATS_HOST=nats://localhost:4222

# JWT
JWT_SECRET=your_super_secret_key
ACCESS_TOKEN_TTL_MIN=15
REFRESH_TOKEN_TTL_HR=24
````

---

## 📡 gRPC Methods

```proto
service UserService {
  rpc Register(RegisterRequest) returns (RegisterResponse);
  rpc Login(LoginRequest) returns (LoginResponse);
  rpc GetProfile(GetProfileRequest) returns (GetProfileResponse);
}
```

Example Postman JSON request (Register):

```json
{
  "full_name": "Miras",
  "email": "miras@aitu.edu.kz",
  "password": "12345678",
  "role": "student"
}
```

---

## 🧪 Run

```bash
go run cmd/main.go
```

---

## 📌 Components

* `/cmd/main.go` — application entrypoint
* `/config` — environment config parsing
* `/internal/adapter` — gRPC, Mongo, Redis, NATS handlers
* `/internal/usecase` — business logic
* `/internal/domain` — domain models and interfaces
* `/pkg` — utility packages: JWT, password manager, mongo, etc.

---

## 🧠 Authorization

Include your JWT token in metadata:

```
authorization: Bearer <your_token>
```

---

## 🧑‍💻 Contributing

Pull requests are welcome. Feel free to suggest improvements or fixes!

---

## 📃 License

MIT License

```

Would you like me to also generate example gRPC requests using `grpcurl` or provide `Makefile` or Docker configs next? 🔧
```
