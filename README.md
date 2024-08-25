## SQLite Database with Go

### Description

This project implements a simple system that:
1. Defines a schema for an SQLite database to store actor information (id, first name, last name, and gender).
2. Populates the SQLite database with data from a CSV file (e.g., `IMDB-actors.csv`).
3. Performs queries on the database and prints the results.

Additionally, a test file (`main_test.go`) is included to validate the functionality of the system, including the schema creation, data insertion, and querying.

### Prerequisites

- **Go**: Version 1.16 or later
- **SQLite**: This project uses an in-memory SQLite database for testing, so no setup is needed for SQLite.
- **Go Modules**: Ensure you have Go modules enabled (you should have a `go.mod` file in your project directory).

### Project Structure

```plaintext
.
├── main.go            # Main application logic
├── main_test.go       # Test suite for the main application
├── IMDB-actors.csv    # Sample CSV file with actor data (must be in the same directory as main.go)
└── README.md          # Documentation for the project
```

### Setup

1. **Clone the Repository:**

   Clone the repository to your local machine:

   ```bash
   git clone https://github.com/bilguunbilegt/database_creation.git
   cd database_creation
   ```

2. **Install Dependencies:**

   Make sure Go modules are enabled. Initialize Go modules (if needed) and download dependencies:

   ```bash
   go mod init database_creation
   go mod tidy
   ```

3. **Prepare the CSV File:**

   The program expects a CSV file named `IMDB-actors.csv` in the following format:

   ```csv
   id,first_name,last_name,gender
    2,Michael,"'babeepower' Viera",M
    3,Eloy,'Chincheta',M
   ```

   Make sure this file is in the same directory as `main.go`. The sample file is provided.

### Running the Application

To run the application, use the `go run` command:

```bash
go run main.go
```

The application will:
1. Create an SQLite database with a table for actors.
2. Populate the table with data from `IMDB-actors.csv`.
3. Query the data and print the results to the console.

### Running Tests

To run the tests, simply use the `go test` command:

```bash
go test -v
```

This will:
- Create an in-memory SQLite database for testing.
- Run tests to validate schema creation, data insertion, and query functionality.

### Example Output

**Application Output (main.go):**

```plaintext
Schema created successfully.
Data inserted successfully.
ID | First Name | Last Name | Gender
-----------------------------------
2  | Michael  | 'babeepower' Viera  | M
3  | Eloy     | 'Chincheta'         | F
```

**Test Output (main_test.go):**

```plaintext
2024/08/25 14:28:16 Schema created successfully.
2024/08/25 14:28:16 Schema created successfully.
2024/08/25 14:28:16 Data inserted successfully.
2024/08/25 14:28:16 Schema created successfully.
PASS
ok  	database_creation	0.256s
```


### Author

- **Bilguun Byambadorj** (bilguunbyambadorj2026@u.northwestern.edu)

Feel free to reach out for any questions or issues regarding this project.
