# Real-Time Chat CLI Application Achievement System

## Achievement status: 15/100

## 1. Foundation Level

### 🛠️ Achievement: Setup and Initialization (5 Points)

- [X] Create a new Golang project with proper folder structure (`cmd`, `pkg`, `internal`, etc.) 
- [X] Initialize a Go module (`go mod init`)
- [X] Install essential libraries (e.g., `gorilla/websocket`, `cobra` for CLI)

### 🛠️ Achievement: User Registration & Login (10 Points)

- [X]Implement a way to register users (username and password stored securely)
- [X]Implement login functionality (e.g., session-based or token-based authentication)
- [X]Create a mock database (e.g., JSON file, SQLite, or in-memory map)

**Total so far: 15 points**

## 2. Basic Chat Features

### 💬 Achievement: Create Rooms (10 Points)

- Allow users to create chat rooms (unique room names)
- Persist room information (e.g., room owner, participants)

### 💬 Achievement: Join/Leave Rooms (10 Points)

- Allow users to list available rooms and join a specific one
- Enable users to leave rooms gracefully

### 💬 Achievement: Send and Receive Messages (10 Points)

- Build functionality for users in the same room to send messages
- Broadcast received messages to all users in the room

**Total so far: 45 points**

## 3. Real-Time Communication

### ⚡ Achievement: WebSocket Integration (15 Points)

- Use `gorilla/websocket` to establish real-time communication
- Implement a connection handler for managing active users
- Handle edge cases: disconnected users, reconnects, etc.

### ⚡ Achievement: Message Storage (5 Points)

- Save chat history for each room in a file or database
- Allow users to fetch chat history upon joining a room

**Total so far: 65 points**

## 4. CLI User Experience

### 🖥️ Achievement: Command-Line Commands (10 Points)

- Use `cobra` to implement CLI commands (`register`, `login`, `create-room`, `join-room`, `send`, `leave`, etc.)
- Validate user inputs for these commands

### 🖥️ Achievement: Pretty CLI Output (5 Points)

- Use colors or formatting libraries (e.g., `github.com/fatih/color`) to make the CLI interactive
- Display messages with timestamps, usernames, and room names

**Total so far: 80 points**

## 5. Advanced Features

### 🛡️ Achievement: User Authentication Middleware (5 Points)

- Enforce authentication for all room actions
- Ensure only authenticated users can join/send messages

### 🛡️ Achievement: Room Ownership Management (5 Points)

- Restrict room deletion or settings changes to the room owner

### 🚀 Achievement: Notifications (5 Points)

- Notify users when someone joins or leaves the room

### 🛠️ Achievement: Graceful Shutdown (5 Points)

- Ensure the server cleans up resources (e.g., active connections, unsaved messages) on shutdown

**Total so far: 100 points**

## Bonus Achievements (Optional, No Points 🏆)

- Add support for private messaging between two users
- Enable file sharing (e.g., images or text files) in rooms
- Integrate testing frameworks to cover key features
- Add encryption for messages in transit

