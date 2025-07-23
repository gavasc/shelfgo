# ShelfGo 📚

**ShelfGo** is a lightweight, self-hosted **WebDAV server written in Go**, designed as a personal ebook library for devices like **KOReader**, desktops, and mobile devices.
Optimal for hosting in your homelab.

---

## 🚀 Features
- 📖 Exposes an ebook directory via WebDAV
- 🔐 Basic Authentication support
- 🐳 Runs easily in Docker
- 🌐 Secure access via **Cloudflare Tunnel**

---

## 🛠️ Usage

### 1. Clone and Build
```bash
git clone https://github.com/gavasc/shelfgo.git
cd shelfgo
```

### 2. Environment Variables

| Variable             | Description                      | Example            |
|----------------------|----------------------------------|--------------------|
| `SHELFGO_USER`       | Username for Basic Auth          | `admin`            |
| `SHELFGO_PASS`       | Password for Basic Auth          | `secret`           |
| `SHELFGO_PORT`       | Internal port ShelfGo listens on | `8080`             |
| `SHELFGO_DIR`        | Directory to serve over WebDAV   | `/ebooks`          |
| `CF_TOKEN`  | Cloudflare Tunnel token          | *your CF token*    |

Example `.env` file:
```
SHELFGO_USER=admin
SHELFGO_PASS=secret
SHELFGO_PORT=8080
SHELFGO_DIR=/ebooks
CF_TOKEN=your_cloudflared_token
```

### 3. Config file inside /config/
```javascript
{
    "allowed_formats": [".epub", ".pdf", ...]
}
```

### 4. Run with Docker Compose
This setup includes:
- `shelfgo`: the WebDAV server
- `cloudflared`: Cloudflare Tunnel for secure public access

```bash
docker compose up --build -d
```

---

## 📱 Accessing ShelfGo
- **KOReader:** Add a **WebDAV account** pointing to your Cloudflare hostname.
- **Desktop:** Use **Cyberduck**, **WinSCP**, or OS-native WebDAV mounting.
- **CLI:** Tools like `curl` or `cadaver`.
