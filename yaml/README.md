# 🔐 Go SOPS Secret Management

A Go application demonstrating secure configuration management using [SOPS (Secrets OPerationS)](https://github.com/mozilla/sops) for encrypting sensitive data like database passwords, API keys, and JWT secrets.

## 🌟 Features

- ✅ **Secure Secret Storage**: Encrypt sensitive configuration data at rest
- 🔑 **GPG Encryption**: Uses 4096-bit RSA encryption for maximum security
- 🚀 **Runtime Decryption**: Automatically decrypt secrets when the application starts
- 📝 **Type Safety**: Go structs provide compile-time safety for configuration
- 🔄 **Easy Secret Management**: Edit encrypted files directly with SOPS
- 🛡️ **Version Control Safe**: Store encrypted secrets safely in Git repositories

## 📋 Prerequisites

- Go 1.24.3 or higher
- SOPS (installed via Homebrew: `brew install sops`)
- GPG (installed via Homebrew: `brew install gnupg`)

## 🚀 Quick Start

### 1. Clone and Setup

```bash
git clone <your-repo>
cd go-sops/yaml
go mod tidy
```

### 2. Run the Application

```bash
go run main.go
```

The application will automatically decrypt the `config.sops.yaml` file and display the configuration.

## 📁 Project Structure

```
yaml/
├── config.yaml           # Original plaintext config (backup)
├── config.sops.yaml      # SOPS-encrypted configuration
├── go.mod                # Go module dependencies
├── main.go               # Main application with SOPS integration
└── README.md             # This file

../
├── .sops.yaml            # SOPS configuration file
└── sops-demo.sh          # Demo script for SOPS operations
```

## ⚙️ Configuration Structure

The application uses the following configuration structure:

```yaml
storage:
  psql:
    host: 127.0.0.1
    port: 5432
    database: testtt
    username: postgres
    password: 12345  # 🔒 Encrypted with SOPS
    pg_pool_max_conn: 400
  redis:
    addr: localhost
    port: 6379
    username: user
    password: secu4e  # 🔒 Encrypted with SOPS
    db: 0
jwt:
  auth: SECRET_KEY_8899    # 🔒 Encrypted with SOPS
```

## 🔧 SOPS Operations

### View Encrypted File
```bash
cat config.sops.yaml
# Shows encrypted data with ENC[...] values
```

### Decrypt and View
```bash
sops -d config.sops.yaml
# Shows decrypted plaintext values
```

### Edit Encrypted File
```bash
sops config.sops.yaml
# Opens your default editor with decrypted content
# Automatically re-encrypts when you save and exit
```

### Encrypt a New File
```bash
sops -e config.yaml > config.sops.yaml
```

## 🔑 GPG Key Management

### List Available Keys
```bash
gpg --list-secret-keys --keyid-format LONG
```

### Current GPG Key Fingerprint
```
14093FAD0219A1D1B52761B4A88742FB6C975643
```

*Note: This key is configured in `../.sops.yaml` for automatic encryption/decryption.*

## 📊 Application Output

When you run the application, you'll see:

```
🔓 Successfully loaded and decrypted configuration:
====================================================
📊 Database Configuration:
  Host: 127.0.0.1
  Port: 5432
  Database: testtt
  Username: postgres
  Password: 12345
  Max Connections: 400

🔴 Redis Configuration:
  Address: localhost
  Port: 6379
  Username: user
  Password: secu4e
  Database: 0

🔐 JWT Configuration:
  Auth Key: SECRET_KEY_8899

======================================================
🚀 Example Usage:
PostgreSQL DSN: postgresql://postgres:12345@127.0.0.1:5432/testtt
Redis URL: redis://user:secu4e@localhost:6379/0
```

## 🔒 Security Best Practices

1. **Never commit plaintext secrets** to version control
2. **Keep your GPG private key secure** and backed up
3. **Use different keys** for different environments (dev/staging/prod)
4. **Rotate secrets regularly** using SOPS edit functionality
5. **Limit access** to GPG keys on production systems

## 🛠️ Development Workflow

### Adding New Secrets

1. Edit the encrypted config:
   ```bash
   sops config.sops.yaml
   ```

2. Add your new secret in the editor that opens

3. Save and exit - SOPS automatically re-encrypts

4. Update Go structs in `main.go` if needed

### Rotating Secrets

1. Edit encrypted config: `sops config.sops.yaml`
2. Update the secret values
3. Deploy the updated configuration

## 🐛 Troubleshooting

### GPG Key Issues
```bash
# If you get GPG errors, ensure your key is properly set up:
gpg --list-secret-keys

# Import GPG key if needed:
gpg --import your-private-key.asc
```

### SOPS Decryption Errors
```bash
# Verify SOPS can decrypt:
sops -d config.sops.yaml

# Check SOPS configuration:
cat ../.sops.yaml
```

### Go Module Issues
```bash
# Clean and rebuild dependencies:
go mod tidy
go mod download
```

## 📚 Learn More

- [SOPS Documentation](https://github.com/mozilla/sops)
- [GPG Documentation](https://gnupg.org/documentation/)
- [Go YAML Package](https://gopkg.in/yaml.v3)

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test thoroughly
5. Submit a pull request

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

---

**🔐 Keep your secrets safe with SOPS!** 🚀
