# Toronto Time API
A simple API built using Go (Golang) and MySQL to log and retrieve time data. This project provides endpoints to fetch the current time and log previous time entries.
Features
* Current Time: Returns the current date and time.
* Logged Times: Retrieves all previously logged times from the MySQL database.
* MySQL Integration: Stores logged times in a MySQL database.
## Table of Contents
1. Technologies
2. Getting Started
3. Running the Application
4. API Endpoints
5. Dockerization
6. Contributing


## Technologies
This project uses the following technologies:
* Go (Golang): The backend language used to build the API.
* MySQL: Used to store and manage time data.
* Docker: For containerizing the application and MySQL database.
* Docker Compose: To manage multi-container applications.
* Golang MySQL Driver: github.com/go-sql-driver/mysql for MySQL database interaction.

## Getting Started
Prerequisites
* Install Docker
* Install Go (version 1.23 or above)
Environment Variables
The project uses the following environment variables to configure the MySQL database connection:
* MYSQL_USER: The MySQL username (e.g., root).
* MYSQL_PASSWORD: The MySQL password (e.g., root).
* MYSQL_HOST: The hostname of the MySQL database (e.g., mysql).
* MYSQL_DB: The name of the MySQL database (e.g., toronto_time).

## Running the Application
Docker Compose
To run the application and MySQL container together using Docker Compose:
1. Make sure Docker is running on your machine.
2. Clone the repository:
      git clone https://github.com/yourusername/toronto_time_api.git
   cd toronto_time_api
   
3. Build and start the application and MySQL container: bash Copy code   docker-compose up --build
4.    This will:
    * Build the Go application.
    * Set up the MySQL database container.
    * Start both containers.
5. The API will be available at http://localhost:8080.

## API Endpoints
### 1. GET /current-time
Fetches the current time in the Toronto timezone.
Example Response:

{
  "current_time": "2024-11-28 14:30:00"
}
### 2. GET /logged-times
Retrieves all logged times from the database.



## Dockerization
This project is dockerized for easy deployment. It uses a Dockerfile to build the Go application and a docker-compose.yml file to manage the application and MySQL database containers.
Build and Run Docker Containers
1. Build the Docker images: bash Copy code   docker-compose build
  
2. Start the containers: bash Copy code   docker-compose up
   
This will start both the API and MySQL containers. By default, the API will be available at http://localhost:8080, and the MySQL database will be accessible inside the container.
Stopping Containers
To stop the running containers, use:

docker-compose down

## Contributing
We welcome contributions to this project! If you'd like to contribute, please follow these steps:
1. Fork this repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes.
4. Test your changes.
5. Submit a pull request with a description of your changes.

# RESULT 
## 1. Set Up MySQL Database:

1.Install MySQL and create a new database.<br/>
2.Create a table named time_log with at least two fields: id (primary key) and timestamp.<br/>
![image](https://github.com/user-attachments/assets/6a85b69a-160a-4219-95fe-5cabe8b9b47d)

## 2. API Development:

1.Write a Go application with a web server.<br/>
2.Create an API endpoint /current-time that returns the current time in Toronto.<br/>

```
func main() {
	// Open a log file to store logs
	logFile, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	// Set log output to the log file
	log.SetOutput(logFile)

	// Read MySQL connection details from environment variables
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	dbname := os.Getenv("MYSQL_DB")

	// Create connection string (DSN)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", user, password, host, dbname)

	// Initialize the database connection
	err = db.InitDB(dsn) // Pass dsn as an argument
	if err != nil {
		log.Fatalf("Error initializing the database: %v", err)
	}
	defer db.DB.Close()

	// Register handlers
	http.HandleFunc("/current-time", handlers.CurrentTimeHandler)
	http.HandleFunc("/logged-times", handlers.GetLoggedTimesHandler) // New handler for logged times

	// Start the server
	log.Println("Server started at :8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
```
## 3.Time Zone Conversion:

1.Use Go's time package to handle the time zone conversion to Toronto's local time.
![image](https://github.com/user-attachments/assets/e9bdf71f-5f2b-4e32-9024-0747b59c880b)

## 4.Database Connection:

1.Connect to your MySQL database from your Go application.<br/>
2.On each API call, insert the current time into the time_log table.<br/>

```
dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", user, password, host, dbname)
```

## 5.Return Time in JSON:

Format the response from the /current-time endpoint in JSON.
![image](https://github.com/user-attachments/assets/879f583d-b735-49e2-bbe0-86781485a930)

## 6. Implemented log file
Implement logging in your Go application to log events and errors.<br/>
Create an additional endpoint to retrieve all logged times from the database.
![image](https://github.com/user-attachments/assets/7c62a4bd-d59a-4960-9fb0-e40d6d7e967e)


## 7. Dockerization
Dockerize your Go application and the MySQL database for easy deployment.
![image](https://github.com/user-attachments/assets/9dfbc37c-2c30-45ff-ba0d-577a19f5e9d4)











