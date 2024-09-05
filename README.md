
# Redis Clone Server

This project is a simple Redis-like server implemented in Go. It listens for incoming TCP connections on port `6379` and handles basic Redis commands. The server also supports AOF (Append-Only File) persistence for specific commands.

## Features

- **TCP Server**: Listens on port `6379` for incoming client connections.
- **RESP Protocol**: Parses incoming requests using the Redis Serialization Protocol (RESP).
- **Command Handlers**: Supports basic Redis commands with handlers for each command.
- **AOF Persistence**: Writes `SET` and `HSET` commands to an AOF file to persist data.

## Getting Started

### Prerequisites

- Go 1.18 or higher installed on your system.

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/Anshul091/RedisDB.git
   cd RedisDB
   ```

2. Run the server:

   ```bash
   go run main.go
   ```

The server will start listening on port `6379`.

## Usage

### Connecting to the Server

You can connect to the server using a Redis client.

### Supported Commands

- **SET**: Set a key-value pair.
  
  ```plaintext
  SET key value
  ```
  
- **GET**: Get the value of a key.
  
  ```plaintext
  GET key
  ```

- **HSET**: Set a field in a hash.
  
  ```plaintext
  HSET hash field value
  ```

- **HGET**: Get the value of a field in a hash.
  
  ```plaintext
  HGET hash field
  ```

Other commands may be supported as implemented in the server's command handlers.

### Persistence

The server persists `SET` and `HSET` commands to an AOF file (`database.aof`). The AOF file is created or appended to during runtime to ensure data is not lost on restart.

## Code Overview

- **`main.go`**: The entry point of the server. Handles setting up the TCP listener, accepting connections, reading requests, and delegating them to command handlers.
- **Command Handlers**: Defined for each supported command. They are used to process the incoming arguments and perform the respective operation.
- **AOF**: Implements writing operations to the `database.aof` file for persistence.

## Error Handling

The server has basic error handling and will print errors to the console. For production use, more robust error handling and logging should be implemented.

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
