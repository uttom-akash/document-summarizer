internal/
│
├── domain/                      # Core business logic
│   ├── file/                    # File Aggregate
│   │   ├── entity.go            # Represents the File entity
│   │   ├── repository.go        # Interface for File repository
│   │   ├── value_objects.go     # Value objects like FilePath, Metadata
│   │   └── errors.go            # File-specific domain errors
│   │
│   └── bucket/                  # Bucket management
│       ├── entity.go            # Represents Bucket entity
│       ├── repository.go        # Interface for Bucket repository
│       └── value_objects.go     # Bucket-specific value objects
│
├── application/                 # Application Layer (Use Cases)
│   ├── file_service.go          # File-related workflows (upload, download)
│   └── bucket_service.go        # Bucket-related workflows
│
├── infrastructure/              # External systems and integrations
│   ├── persistence/             # Persistence Layer (Database)
│   │   ├── gorm/                # GORM implementations
│   │   │   ├── file_repo.go     # File repository implementation
│   │   │   └── bucket_repo.go   # Bucket repository implementation
│   │   └── migrations/          # Database schema and migration files
│   │
│   ├── storage/                 # File Storage backends
│   │   ├── local_storage.go     # Local file system implementation
│   │   └── s3_storage.go        # S3-compatible storage implementation
│   │
│   ├── http/                    # HTTP Layer (Gin Handlers)
│   │   ├── handlers/            # Gin handlers for file/bucket routes
│   │   │   ├── file_handler.go  # Handles file upload/download APIs
│   │   │   └── bucket_handler.go
│   │   └── routes/              # Gin routes configuration
│   │
│   ├── fx/                      # Fx Modules (Dependency Injection)
│   │   ├── gin.go               # Initializes Gin engine
│   │   ├── gorm.go              # Initializes GORM connection
│   │   ├── storage.go           # Initializes storage backends
│   │   └── app.go               # Root Fx module
│   │
│   └── logger/                  # Logger setup
│       └── logger.go
│
├── shared/                      # Shared utilities, errors, and responses
│   ├── errors.go                # Global error definitions
│   ├── response.go              # Response helper functions
│   └── utils.go                 # Common utilities
│
└── tests/                       # Internal tests for domain, services, etc.
    ├── file_test.go
    └── bucket_test.go
