# tchat

> A lightweight, terminal-based chat application for direct machine-to-machine communication over TCP.

---

## Features

- Direct peer-to-peer messaging over TCP
- Works on Windows, Linux, and macOS
- No account, server, or internet required (same network)
- Timestamped messages with custom usernames
- Type `/quit` to exit cleanly

---

## Installation

### Windows

Open **PowerShell as Administrator** and run:

```powershell
mkdir C:\Tools
Invoke-WebRequest -Uri "https://github.com/Tomdabom27/tchat/releases/download/v1.0/tchat.exe" -OutFile "C:\Tools\tchat.exe"
setx PATH "$env:PATH;C:\Tools"
```

Then **restart your terminal** and verify the install:

```powershell
tchat
```

---

### Linux

**AMD64 (most desktops/servers):**

```bash
curl -L -O https://github.com/Tomdabom27/tchat/releases/download/v1.0/tchat-linux-amd64.tar.gz
tar -xzf tchat-linux-amd64.tar.gz
```

**ARM64 (Raspberry Pi, Apple Silicon VMs, etc.):**

```bash
curl -L -O https://github.com/Tomdabom27/tchat/releases/download/v1.0/tchat-linux-arm64.tar.gz
tar -xzf tchat-linux-arm64.tar.gz
```

**Then for both architectures:**

```bash
chmod +x tchat
sudo mv tchat /usr/local/bin/
```

Verify the install:

```bash
tchat
```

---

## Usage

### 1. Host a session

Run this on the machine that will be hosting:

```bash
tchat host 8080
```

Find your local IP address:
- **Windows:** run `ipconfig` and look for `IPv4 Address`
- **Linux/macOS:** run `ip a` or `ifconfig`

### 2. Join a session

Run this on the machine that is connecting (replace with the host's IP):

```bash
tchat join 192.168.1.x:8080
```

Both users will be prompted for a username. Once connected, just start typing.

---

## Commands

| Command | Description |
|---------|-------------|
| `tchat host <port>` | Start hosting on the given port |
| `tchat join <ip:port>` | Join a host at the given address |
| `/quit` | Leave the chat session |

---

## Building from Source

Requires [Go 1.21+](https://go.dev/dl/)

```bash
git clone https://github.com/Tomdabom27/tchat.git
cd tchat
go build -o tchat .
```

---

## Notes

- Both machines must be on the **same local network** for the default setup to work
- For connections across the internet, use [Tailscale](https://tailscale.com) to create a virtual local network between machines — no code changes needed
- Only **one client** can connect to a host at a time
