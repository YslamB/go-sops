# 🔐 Go SOPS Environment Variable Manager

A comprehensive Go application demonstrating secure environment variable management using [SOPS (Secrets OPerationS)](https://github.com/mozilla/sops) for encrypting `.env` files containing sensitive configuration data like API keys, database passwords, and OAuth secrets.

## 🌟 Features

- ✅ **Secure .env File Encryption**: Encrypt environment variables at rest using SOPS
- 🔑 **GPG Encryption**: Uses 4096-bit RSA encryption for maximum security
- 🚀 **Dual Loading Methods**: Load into structured config OR system environment variables
- 📝 **Smart Secret Masking**: Automatically masks sensitive values in logs
- 🔄 **Easy Secret Management**: Edit encrypted .env files directly with SOPS
- 🛡️ **Production Ready**: Safe storage of encrypted secrets in Git repositories
- 📊 **Comprehensive Logging**: Beautiful, organized display of loaded configuration

## 📋 Prerequisites

- Go 1.24.3 or higher
- SOPS (installed via Homebrew: `brew install sops`)
- GPG (installed via Homebrew: `brew install gnupg`)

## 🚀 Quick Start

### 1. Clone and Setup

```bash
git clone <your-repo>
cd go-sops/env
go mod tidy
```

### 2. Run the Application

```bash
go run main.go
```

The application will automatically decrypt the `config.sops.env` file and demonstrate both loading methods.

## 📁 Project Structure

```
env/
├── config.env            # Original plaintext environment file (backup)
├── config.sops.env       # SOPS-encrypted environment file
├── go.mod                # Go module dependencies  
├── main.go               # Main application with SOPS integration
├── .sops.yaml            # SOPS configuration for .env files
└── README.md             # This documentation
```

## ⚙️ Environment Variables Structure

The application manages the following categories of environment variables:

### 📊 Database Configuration
```bash
DB_HOST=localhost
DB_PORT=5432
DB_NAME=production_db
DB_USER=admin
DB_PASSWORD=super_secret_password_123  # 🔒 Encrypted
DB_MAX_CONNECTIONS=100
```

### 🔴 Redis Configuration
```bash
REDIS_URL=redis://default:password@host:6379  # 🔒 Encrypted
REDIS_PASSWORD=redis_password_456              # 🔒 Encrypted
```

### 🔐 API Keys & Secrets
```bash
JWT_SECRET=jwt_super_secret_key_789       # 🔒 Encrypted
API_KEY=sk-1234567890abcdef               # 🔒 Encrypted
STRIPE_SECRET_KEY=sk_test_abcdef123456    # 🔒 Encrypted
SENDGRID_API_KEY=SG.1234567890abcdef     # 🔒 Encrypted
```

### 🔑 OAuth Credentials
```bash
GOOGLE_CLIENT_ID=123456789-abcdef.apps.googleusercontent.com
GOOGLE_CLIENT_SECRET=GOCSPX-abcdef123456         # 🔒 Encrypted
GITHUB_CLIENT_ID=abcdef123456
GITHUB_CLIENT_SECRET=github_secret_789           # 🔒 Encrypted
```

### 🌐 External Services
```bash
WEBHOOK_URL=https://api.example.com/webhook
NOTIFICATION_SERVICE_URL=https://notifications.example.com/api
```

### ⚙️ Environment Settings
```bash
ENVIRONMENT=production
DEBUG=false
LOG_LEVEL=info
```

### 🔒 Encryption Keys
```bash
ENCRYPTION_KEY=32_byte_encryption_key_here_123   # 🔒 Encrypted
SIGNING_KEY=signing_key_for_tokens_456           # 🔒 Encrypted
```

## 💻 Application Features

### 🎯 Dual Loading Methods

#### Method 1: Structured Configuration
```go
config, err := LoadSOPSEnv("config.sops.env")
if err != nil {
    log.Fatal(err)
}

// Access via struct fields
dbPassword := config.DBPassword
jwtSecret := config.JWTSecret
```

#### Method 2: System Environment Variables
```go
err := LoadSOPSEnvToSystem("config.sops.env")
if err != nil {
    log.Fatal(err)
}

// Access via os.Getenv
dbPassword := os.Getenv("DB_PASSWORD")
jwtSecret := os.Getenv("JWT_SECRET")
```

### 🛡️ Smart Secret Masking

The application automatically detects and masks sensitive values:

```
DB_PASSWORD: su*********************23
JWT_SECRET: jw********************89
API_KEY: sk***************ef
```

Variables containing these keywords are considered sensitive:
- `PASSWORD`, `SECRET`, `KEY`, `TOKEN`, `CREDENTIAL`, `PRIVATE`

## 🔧 SOPS Operations

### View Encrypted File
```bash
cat config.sops.env
# Shows encrypted data with ENC[...] values
```

### Decrypt and View
```bash
sops -d config.sops.env
# Shows decrypted plaintext values
```

### Edit Encrypted File
```bash
sops config.sops.env
# Opens your default editor with decrypted content
# Automatically re-encrypts when you save and exit
```

### Encrypt a New .env File
```bash
sops -e config.env > config.sops.env
```

## 🔑 GPG Key Management

### Current GPG Key Fingerprint
```
14093FAD0219A1D1B52761B4A88742FB6C975643
```

### SOPS Configuration
The `.sops.yaml` file is configured specifically for .env files:

```yaml
creation_rules:
  - path_regex: \.(env|dotenv)$
    pgp: >-
      14093FAD0219A1D1B52761B4A88742FB6C975643
    unencrypted_regex: '^(#.*|[A-Z_]+)='
```

## 📊 Application Output

When you run the application, you'll see:

```
🔐 SOPS Environment Variable Manager
=====================================

📋 Method 1: Loading into structured configuration
🔓 Successfully loaded and decrypted environment configuration:
================================================================

📊 Database Configuration:
  DB_HOST: localhost
  DB_PORT: 5432
  DB_NAME: production_db
  DB_USER: admin
  DB_PASSWORD: su*********************23
  DB_MAX_CONNECTIONS: 100

🔴 Redis Configuration:
  REDIS_URL: re*****************************************************79
  REDIS_PASSWORD: re**************56

... (continued with all configuration sections)

🚀 Example Usage:
1️⃣ Using Structured Config:
   Database DSN: postgresql://admin:su*********************23@localhost:5432/production_db

2️⃣ Using System Environment Variables:
   Database DSN: postgresql://admin:su*********************23@localhost:5432/production_db
```

## 🔒 Security Best Practices

1. **Never commit plaintext .env files** to version control
2. **Keep your GPG private key secure** and backed up
3. **Use different keys** for different environments (dev/staging/prod)
4. **Rotate secrets regularly** using SOPS edit functionality
5. **Limit access** to GPG keys on production systems
6. **Use strong, unique passwords** for all services
7. **Enable secret masking** in logs and output

## 🛠️ Development Workflow

### Adding New Environment Variables

1. Edit the encrypted config:
   ```bash
   sops config.sops.env
   ```

2. Add your new environment variable in the editor that opens

3. Update Go structs in `main.go` if using structured config method

4. Update the `ourVars` slice for system environment display

### Rotating Secrets

1. Edit encrypted config: `sops config.sops.env`
2. Update the secret values
3. Deploy the updated configuration
4. Restart applications to load new values

## 🏭 Production Deployment

### Docker Example
```dockerfile
FROM golang:1.24-alpine

# Install sops and gpg
RUN apk add --no-cache sops gnupg

# Copy application
COPY . /app
WORKDIR /app

# Import GPG key (securely)
RUN echo "$GPG_PRIVATE_KEY" | gpg --import

# Run application
CMD ["go", "run", "main.go"]
```

### Kubernetes Example
```yaml
apiVersion: v1
kind: Secret
metadata:
  name: gpg-key
data:
  private.key: <base64-encoded-gpg-private-key>
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: sops-app
spec:
  template:
    spec:
      containers:
      - name: app
        image: sops-app:latest
        env:
        - name: GNUPGHOME
          value: "/tmp/.gnupg"
        volumeMounts:
        - name: gpg-key
          mountPath: "/tmp/.gnupg"
      volumes:
      - name: gpg-key
        secret:
          secretName: gpg-key
```

## 🐛 Troubleshooting

### GPG Key Issues
```bash
# Check available keys
gpg --list-secret-keys

# Import key if needed
gpg --import private-key.asc

# Test decryption
sops -d config.sops.env
```

### SOPS Configuration Issues
```bash
# Verify SOPS config
cat .sops.yaml

# Test encryption
echo "TEST=value" | sops -e /dev/stdin
```

### Go Module Issues
```bash
# Clean and rebuild
go mod tidy
go mod download
go clean -cache
```

## 📚 Dependencies

- **[github.com/joho/godotenv](https://github.com/joho/godotenv)**: For parsing .env files
- **SOPS**: For encryption/decryption operations
- **GPG**: For cryptographic operations

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test thoroughly with encrypted configs
5. Submit a pull request

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

---

**🔐 Secure your environment variables with SOPS!** 🚀

### Key Advantages:

✅ **Version Control Safe**: Store encrypted secrets in Git  
✅ **Easy Management**: Edit secrets with simple commands  
✅ **Flexible Loading**: Choose between structured config or env vars  
✅ **Production Ready**: Battle-tested encryption for enterprise use  
✅ **Developer Friendly**: Clear logging and error handling  
✅ **Security First**: Automatic secret masking and best practices
