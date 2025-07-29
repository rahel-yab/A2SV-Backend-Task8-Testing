package domain

import (
	"context"
	"time"
)

// User represents a user in the system
type User struct {
	ID       string
	Username string
	Email    string
	Password string
	Role     string
}

// Task represents a task in the system
type Task struct {
	ID          string
	Title       string
	Description string
	DueDate     time.Time
	Status      string
}

// ITaskRepository defines the interface for task repository operations
type ITaskRepository interface {
	AddTask(ctx context.Context, task *Task) error
	GetAllTasks(ctx context.Context) ([]Task, error)
	GetTaskByID(ctx context.Context, id string) (*Task, error)
	UpdateTask(ctx context.Context, task *Task) error
	DeleteTask(ctx context.Context, id string) error
}

// IUserRepository defines the interface for user repository operations
type IUserRepository interface {
	AddUser(ctx context.Context, user *User) error
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	GetUserByUsername(ctx context.Context, username string) (*User, error)
	IsUsersCollectionEmpty(ctx context.Context) (bool, error)
	UserExistsByEmail(ctx context.Context, email string) (bool, error)
	UserExistsByUsername(ctx context.Context, username string) (bool, error)
	PromoteUserToAdmin(ctx context.Context, identifier string) error
}

// IPasswordService defines the interface for password hashing operations
type IPasswordService interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}

// IJWTService defines the interface for JWT token operations
type IJWTService interface {
	GenerateToken(user *User) (string, error)
}
