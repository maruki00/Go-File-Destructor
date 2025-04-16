# gofd

**gofd** is a command-line utility written in Go for securely deleting files and directories. It ensures that sensitive data is overwritten multiple times before deletion, making recovery extremely difficult.

## Features

- Secure deletion of files and directories
- Customizable number of overwrite passes
- Lightweight and fast
- Cross-platform support

## Requirements

- Go 1.16 or higher

## Installation

1. Clone the repository:
```bash
   git clone https://github.com/maruki00/gofd.git
```

2. Navigate to the project directory:
```bash
   cd gofd
```

3. Build the executable:
```bash
   go build -o gofd main.go
```

## Usage

Run the program with the following options:
```bash
gofd [options]
```
Options:

- -path: The full path of the file or directory to delete
- -type: Specify 'file' or 'dir' depending on the target
- -times: Number of times to overwrite the content before deletion

## Examples

1. To securely delete a file:
   ```bash
   gofd -path=/home/user/secret.txt -type=file -times=5
   ```

2. To securely delete a directory:
```bash
   gofd -path=/home/user/old_logs -type=dir -times=3
```

## Disclaimer

Use with caution. This tool will permanently destroy data, and it cannot be recovered.

## License

This project is licensed under the MIT License.

## Author

Developed by maruki00
