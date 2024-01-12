# GoMongoMVCAuthAPI

GoMongoMVCAuthAPI is a simple Go (Golang) project that demonstrates the implementation of a CRUD REST API with MongoDB, following the MVC (Model-View-Controller) architectural pattern. The project includes a basic authentication layer and uses a custom HTTP router.

## Project Structure

- **controllers**: Contains the controllers responsible for handling HTTP requests and defining CRUD operations.
- **models**: Defines the data models and interacts with the MongoDB database.
- **routes**: Implements the custom HTTP router and sets up the API routes.
- **middleware**: Includes the authentication middleware.
- **main.go**: The main entry point of the application.

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/)
- [MongoDB](https://docs.mongodb.com/manual/installation/)
- [Docker](https://www.docker.com/get-started)

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/amro-alasri/GoMongoMVCAuthAPI.git
   cd GoMongoMVCAuthAPI
   ```

2. Initialize Go modules:
   ```bash
   go mod init GoMongoMVCAuthAPI
   ```
3. Install dependencies:

   ```bash
    go get -u
   ```

### Usage

```bash
  go run main.go
```

### API Endpoints

API Endpoints
GET /user/{id}: Retrieve a specific user by ID.
POST /user: Create a new user (protected route).
DELETE /user/{id}: Delete a user (protected route).
