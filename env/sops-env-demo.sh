#!/bin/bash

echo "🔐 SOPS Environment Variable Manager Demo"
echo "========================================="

echo "📁 Current directory: $(pwd)"
echo

echo "1️⃣ View encrypted .env file:"
echo "cat config.sops.env | head -10"
echo "---"
head -10 config.sops.env
echo "..."
echo

echo "2️⃣ Decrypt and view .env file:"
echo "sops -d config.sops.env"
echo "---"
sops -d config.sops.env
echo

echo "3️⃣ Edit encrypted .env file (opens editor):"
echo "sops config.sops.env"
echo "Note: This will open your default editor with decrypted content"
echo

echo "4️⃣ Re-encrypt a plain .env file:"
echo "sops -e config.env.backup > config.sops.env"
echo

echo "5️⃣ Run the Go application (both loading methods):"
echo "go run main.go"
echo "---"
go run main.go
echo

echo "6️⃣ Test specific environment variable access:"
echo "sops -d config.sops.env | grep JWT_SECRET"
echo "---"
sops -d config.sops.env | grep JWT_SECRET
echo

echo "7️⃣ Verify SOPS configuration:"
echo "cat .sops.yaml"
echo "---"
cat .sops.yaml
echo

echo "✅ SOPS .env integration demo complete!"
echo "Your environment variables are now encrypted at rest and decrypted on-demand."
echo ""
echo "🔑 Key Benefits:"
echo "  • Environment variables encrypted with GPG"
echo "  • Safe to store in version control"
echo "  • Two loading methods: structured config + system env vars"
echo "  • Automatic secret masking in logs"
echo "  • Easy secret rotation with 'sops config.sops.env'"
