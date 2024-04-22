# WebSocketServer for providing Realtime BTC prices
This repository contains code for a WebSocket server implemented in Go (Golang) and a client application for testing the server. Additionally, it includes an Artillery load testing script to simulate WebSocket connections and measure performance under load.
## Prerequisites

Before you begin, ensure you have the following installed:

- Go (Golang) - [Installation Guide](https://golang.org/doc/install)
- Node.js and npm - [Installation Guide](https://docs.npmjs.com/downloading-and-installing-node-js-and-npm)
- Artillery - Install using npm:
  ```bash
  npm install -g artillery

## Setup

1. Clone this repository:
   ```bash
   git clone https://github.com/your-username/your-repository.git
   cd your-repository
   
2. Install necessary Go packages for the server:
    ```bash
   go mod tidy

## Starting the WebSocket Server
1. Start the WebSocket server
   ```bash
   go run main.go

## Running the Client (Optional)
1. Open a new terminal.
2. Navigate to the test directory(Client file is present):
   ```bash
   cd test
3. Start the client to test the WebSocket server
    ```bash
   go run main.go
4. The client will connect to the server and send/receive WebSocket messages.

   ## Load Testing with Artillery
1. Open a new terminal.
2. Navigate to the root directory of the repository.
3. Edit the Artillery test script (websocket-test.yml) if necessary.
4. Run the Artillery load test:
    ```bash
   artillery run websocket-test.yml

5. Artillery will simulate WebSocket connections and measure performance metrics.

 ## Result Analysis
 After running the load test, analyze the Artillery results to evaluate the WebSocket server's performance under load. Check for metrics such as response times, message throughput, errors, and system resource usage.




