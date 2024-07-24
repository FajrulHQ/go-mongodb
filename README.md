# Go-MongoDB API

### Description

This project demonstrates a simple CRUD (Create, Read, Update, Delete) application built with Golang and the official MongoDB driver. It allows you to manage a portfolio of projects, including their details, links, and start/end dates.

### Models:
Represents a project with its details

```golang
type Projects struct {
  ID          primitive.ObjectID `json:"_id" bson:"_id"`
  Name        string             `json:"name"`
  Description string             `json:"description"`
  Link        string             `json:"link"`
  Start       *time.Time         `json:"start"`
  End         *time.Time         `json:"end"`
  Created     time.Time          `json:"created"`
  Updated     time.Time          `json:"updated"`
}
```

### Prerequisites:
- Golang (version 1.16 or later)
- MongoDB server

### Installation:
- Clone this repository.
- Install dependencies using `go mod download`.
- Add file `.env` with these variables
```golang
PORT=8080
MONGODB_URI="..."
```

### Running the Application:
- Compile and run the application: `go run main.go`

### API Endpoints:

- `GET /projects`: Retrieves all projects.
- `GET /projects/{id}`: Retrieves a specific project by ID.
- `POST /projects`: Creates a new project. (Payload should be a JSON object representing the Projects struct)
- `PUT /projects/{id}`: Updates an existing project. (Payload should be a JSON object with updated project information)
- `DELETE /projects/{id}`: Deletes an existing project.