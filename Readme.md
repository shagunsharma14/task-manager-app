`README.md`

# Task Manager Application

A full-stack task management web application with user authentication, task CRUD operations, and task filtering capabilities.  
Built with **Golang** using **Gin** framework for the backend API, and **React** for the frontend UI.

---

## Table of Contents

- [Project Overview](#project-overview) 
- [Project Demo](#project-demo)
- [Features](#features)  
- [Technology Stack](#technology-stack)  
- [Backend Architecture](#backend-architecture)  
- [Frontend Architecture](#frontend-architecture)  
- [Setup and Installation](#setup-and-installation)  
- [API Endpoints](#api-endpoints)  
- [Usage](#usage)  
- [Future Enhancements](#future-enhancements)  
- [Testing](#testing)  
- [Contact](#contact)  

---

## Project Overview

This Task Manager App enables users to register and authenticate securely using JWT tokens. Authenticated users can create, view, update, and delete their personal tasks. Tasks support fields such as title, description, status, and due date. The frontend provides an intuitive React interface with live filtering of tasks by status and real-time updates.

This project demonstrates ability to build scalable,
secure full-stack applications with clean code architecture, modular design, and modern best practices.

---
---

## Project Demo

[![Watch the video](https://img.youtube.com/vi/-cuphB-7Ss8/0.jpg)](https://youtu.be/-cuphB-7Ss8)

Watch this 7-minute video where I walk through the features, architecture, and usage of the Task Manager App.

## Features

- User registration & secure login with password hashing and JWT authentication.
- User-specific task CRUD functionality:
  - Create tasks with title, description, status, and optional due date.
  - View a list of all your tasks.
  - Edit task details.
  - Delete tasks.
- Filtering tasks dynamically by status.
- Responsive and accessible React frontend.
- Centralized API client managing JWT tokens for authenticated calls.
- Persistent login via token stored in local storage.
- Loading and error state management.
- Well-structured and documented codebase.
- Comprehensive unit and integration tests.

---

## Technology Stack

- **Backend**:
  - Go (Golang)
  - Gin Web Framework
  - GORM ORM
  - SQLite (development database)
  - JSON Web Tokens (JWT) for authentication
  - bcrypt for password hashing

- **Frontend**:
  - React
  - Axios for API calls
  - Context API for auth state management
  - HTML5, CSS3

- **Development Tools**:
  - Postman (API testing)
  - Go testing package + httptest
  - Create React App

---

## Backend Architecture

- **Models**:
  - `User`: Stores user info with hashed password.
  - `Task`: Stores tasks linked to users (`UserID` foreign key).
- **Authentication**:
  - Secure password storage with bcrypt.
  - JWT tokens issued at login, validated via middleware on protected routes.
- **Routing**:
  - Public endpoints: `/register`, `/login`.
  - Protected endpoints (with JWT middleware): `/tasks` CRUD and `/health`.
- **Services**:
  - Business logic layer interacts with database via GORM.
  - User-specific task filtering is enforced in services.
- **Middleware**:
  - JWT verification extracts username from token and sets context for handlers.
- **Database**:
  - SQLite for ease of development, migrations handled by GORM.

---

## Frontend Architecture

- **Authentication Context**:
  - Stores authenticated user and JWT token.
  - Provides `login` and `logout` functions.
  - Token stored and retrieved from `localStorage`.
- **Components**:
  - `Login` and `Register`: Manage user authentication flows.
  - `TaskForm`: Handles task creation and editing with validation.
  - `TaskList`: Displays filtered list of tasks.
  - `TaskItem`: Represents single task with edit/delete controls.
- **API Client**:
  - Axios instance configured with interceptor to attach JWT token to all API requests.
- **State Management**:
  - React `useState` and `useEffect` for fetching & managing tasks.
  - Error and loading states are handled gracefully.
- **Filtering**:
  - UI dropdown to filter tasks dynamically by status with immediate feedback.

---

## Setup and Installation

### Backend

1. Install Go (version 1.18+ recommended) from [golang.org](https://golang.org/dl/).
2. Clone this repository.
3. Navigate to the backend directory.
4. Run:
   ```bash
   go mod tidy
   go run cmd/main.go
   ```
5. Server listens on port `8080`.

### Frontend

1. Node.js (version 14+) and npm/yarn installed.
2. Navigate to frontend directory.
3. Run:
   ```bash
   npm install
   npm start
   ```
4. The app will open on `http://localhost:3000` and proxy API requests to backend.

---

## API Endpoints

| Method | Endpoint        | Description                   | Auth Required |
|--------|-----------------|-------------------------------|--------------|
| POST   | `/register`     | Register a new user            | No           |
| POST   | `/login`        | Log in and receive JWT token   | No           |
| GET    | `/tasks`        | Get all tasks for authenticated user | Yes          |
| POST   | `/tasks`        | Create a new task              | Yes          |
| PUT    | `/tasks/:id`    | Update a task by ID            | Yes          |
| DELETE | `/tasks/:id`    | Delete a task by ID            | Yes          |
| GET    | `/health`       | Health check endpoint          | No           |

---

## Usage

- Register a new user.
- Log in with credentials.
- Use the task management UI to add, edit, delete, and filter your tasks.
- Tasks are saved per user and secured with JWT authentication.
- Logout to clear the session.

---

## Future Enhancements

- Add filtering by due date and search.
- Add task priorities and categories.
- Implement real-time updates with WebSockets.
- Support file attachments to tasks.
- Improve UI/UX with animations and accessibility improvements.
- Add role-based access control and admin features.
- Implement mobile-friendly responsive design.

---

## Testing

- Backend tests cover user registration, login, authorization middleware, and CRUD operations.
- To run backend tests:

  ```bash
  go test ./internal/handlers -v
  ```

- Frontend testing can be added with Jest and React Testing Library.

---

## Contact

For questions or feedback, please contact:

**Your Name**  
Email: shagunsharmadev@gmail.com
GitHub: [Shagun Sharma](https://github.com/shagunsharma14)

---

Thank you for reviewing my Task Manager Application! I look forward to your feedback.
