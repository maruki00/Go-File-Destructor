# Go File Destructor

Go File Destructor is a simple and effective command-line tool written in Go for securely deleting files and directories. It works by overwriting data multiple times before deletion, ensuring that sensitive content cannot be recovered.

## Features

- Securely delete files or entire directories
- Customizable number of overwrite passes
- Lightweight and fast
- Written in Go — easy to build and run on any platform

## Installation

### Requirements

- Go 1.16 or higher

### Build

1. Clone the repository:

   ```bash
   git clone https://github.com/maruki00/Go-File-Destructor.git
   cd Go-File-Destructor
   ```
2. Build the Project

Make sure you are inside the project directory:

```bash
cd Go-File-Destructor
go build -o file-destructor main.go
