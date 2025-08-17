#!/bin/bash

echo "🔐 SOPS (Secrets OPerationS) Demo Script"
echo "========================================"

# Navigate to yaml directory
cd yaml

echo "📁 Current directory: $(pwd)"
echo

echo "1️⃣ View encrypted config file:"
echo "sops config.sops.yaml"
echo "---"
head -10 config.sops.yaml
echo "..."
echo

echo "2️⃣ Decrypt and view config file:"
echo "sops -d config.sops.yaml"
echo "---"
sops -d config.sops.yaml
echo

echo "3️⃣ Edit encrypted file (opens editor):"
echo "sops config.sops.yaml"
echo "Note: This will open your default editor with decrypted content"
echo

echo "4️⃣ Re-encrypt a plain file:"
echo "sops -e config.yaml.backup > config.sops.yaml"
echo

echo "5️⃣ Run the Go application:"
echo "go run main.go"
echo "---"
go run main.go
echo

echo "✅ SOPS integration complete!"
echo "Your secrets are now encrypted at rest and decrypted on-demand."
