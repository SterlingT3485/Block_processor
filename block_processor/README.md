
# Concurrent Block Processor

## Introduction
This program is a Concurrent Block Processor for the Flow Blockchain simulation, designed to handle blocks and votes through an HTTP API. It ensures blocks are accepted in the correct sequence based on votes received.

## Environment Setup and Dependencies

### Requirements:
- **Go (Golang)**
- **Curl**: For sending HTTP requests during testing.
- **Bash (for Unix/Linux users)**: For running the shell scripts included for testing.

### Installation:
1. **Build the Program**:
   ```bash
   go build -o block_processor main.go
   ```

## Starting the Program
To start the program, run the compiled executable:
```bash
./block_processor
```
This will start the server on `localhost:8080` and begin listening for block and vote submissions.

## Interacting with the Program
You can interact with the program using HTTP POST requests to send blocks and votes:

- **Sending a Block**:
  ```bash
  curl -X POST http://localhost:8080/block -H "Content-Type: application/json" -d '{"id":"a65e9803bb37256c4a663a5c1b", "view": 1234}'
  ```
- **Sending a Vote**:
  ```bash
  curl -X POST http://localhost:8080/vote -H "Content-Type: application/json" -d '{"block_id":"a65e9803bb37256c4a663a5c1b"}'
  ```

## Running Tests
A test script `test_script.sh` is provided to automate sending requests and checking the outputs:

1. **Make the script executable**:
   ```bash
   chmod +x test_script.sh
   ```
2. **Run the script**:
   ```bash
   ./test_script.sh
   ```
   The script sends predefined block and vote data to the server, checks the server output, and validates whether the block has been accepted correctly.

### What the Test Covers
- Sending a block and corresponding vote to the server.
- Checking if the server logs that the block was accepted.

### Test Output
The script outputs whether the test passed or failed based on the server's response and the final log entries.

## Documentation and Support
For more detailed information about the API and its functions, please refer to the source code documentation in `main.go`. If you encounter any issues or have questions, email or message me anytime, thanks.
