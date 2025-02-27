# Gator CLI
A simple command-line tool for managing posts from various feeds.

## Prerequisites
Before you can run the program, ensure you have the following installed:

- PostgreSQL: Required to store and retrieve data.
- Go: Version 1.23.6 or higher is required.

## Installation
### Step 1: Install Go
To install Go, follow the instructions on the official Go website: https://golang.org/dl/.

### Step 2: Install Gator CLI
To install the Gator CLI tool, use the following go install command:

```bash
go install github.com/DNelson35/gator@latest
```

### Step 3: Set up the Configuration File
Create a .gatorconfig file in your home directory (~/.gatorconfig). The file should have the following structure:

```json
{
  "db_url": "postgresql://username:password@localhost:5432/dbname?sslmode=disable",
  "current_user_name": "your_user_name"
}
```

Make sure to replace the values with your own database connection details and the current user name.

## Available Commands
- login: Log in to the system.
- register: Register a new user.
- reset: Reset the user’s password (this will reset the database for the user).
- users: List all users.
- agg: Aggregate data.
- addfeed: Add a new feed.
- feeds: View all available feeds.
- follow: Follow a feed.
- following: View feeds you're following.
- unfollow: Unfollow a feed.
- browse: Browse posts from the feeds you follow.
