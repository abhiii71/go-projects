-----

# Go Movies API

This project provides a simple RESTful API for managing a collection of movies. It's built with Go and uses the `gorilla/mux` router for handling HTTP requests.

## Features

  * **Get all movies:** Retrieve a list of all movies.
  * **Get a single movie:** Fetch details of a specific movie by its ID.
  * **Create a movie:** Add a new movie to the collection.
  * **Update a movie:** Modify an existing movie's details.
  * **Delete a movie:** Remove a movie from the collection.

## Getting Started

### Prerequisites

  * Go installed (higher version recommended)
  * `gorilla/mux` package:
    ```bash
    go get github.com/gorilla/mux
    ```

### Running the Application

1.  **Clone the repository (if applicable):**

    ```bash
    git clone https://github.com/abhiii71/go-projects.git
    cd project2/go-movies-crud
    ```

    (Note: Assuming `go-movies-crud` is your project directory)

2.  **Run the Go application:**

    ```bash
    go run main.go
    ```

    The server will start on port `9090`. You should see the following output:

    ```
    starting server at port: 9090
    ```

## API Endpoints

All API endpoints are accessible at `http://localhost:9090`.

-----

### Get All Movies

  * **URL:** `/movies`
  * **Method:** `GET`
  * **Description:** Retrieves a list of all movies currently stored.
  * **Example Request:**
    ```bash
    curl http://localhost:9090/movies
    ```
  * **Example Response:**
    ```json
    [
      {
        "id": "1",
        "isbn": "4321",
        "title": "dexter",
        "director": {
          "first_name": "dexter",
          "last_name": "morgan"
        }
      }
    ]
    ```

-----

### Get a Single Movie

  * **URL:** `/movie/{id}`
  * **Method:** `GET`
  * **Description:** Retrieves the details of a specific movie using its ID.
  * **URL Parameters:**
      * `id` (string): The ID of the movie to retrieve.
  * **Example Request:**
    ```bash
    curl http://localhost:9090/movie/1
    ```
  * **Example Response:**
    ```json
    {
      "id": "1",
      "isbn": "4321",
      "title": "dexter",
      "director": {
        "first_name": "dexter",
        "last_name": "morgan"
      }
    }
    ```

-----

### Create a Movie

  * **URL:** `/movies`
  * **Method:** `POST`
  * **Description:** Adds a new movie to the collection. A unique ID will be generated for the new movie.
  * **Request Body (JSON):**
    ```json
    {
      "isbn": "123456",
      "title": "Vampire Diaries",
      "director": {
        "first_name": "um",
        "last_name": "idk"
      }
    }
    ```
  * **Example Request:**
    ```bash
    curl -X POST http://localhost:9090/movies \
      -H "Content-Type: application/json" \
      -d '{
            "isbn": "123456",
            "title": "Vampire Diaries",
            "director": {
              "first_name": "um",
              "last_name": "idk"
            }
          }'
    ```
  * **Example Response:**
    ```json
    {
      "id": "GeneratedID",
      "isbn": "123456",
      "title": "Vampire Diaries",
      "director": {
        "first_name": "um",
        "last_name": "idk"
      }
    }
    ```
    (Note: "GeneratedID" will be a random number)

-----

### Update a Movie

  * **URL:** `/movie/{id}`
  * **Method:** `PUT`
  * **Description:** Updates the details of an existing movie identified by its ID.
  * **URL Parameters:**
      * `id` (string): The ID of the movie to update.
  * **Request Body (JSON):**
    ```json
    {
      "isbn": "updated_isbn",
      "title": "updated_title",
      "director": {
        "first_name": "updated_first_name",
        "last_name": "updated_last_name"
      }
    }
    ```
  * **Example Request:**
    ```bash
    curl -X PUT http://localhost:9090/movie/1 \
      -H "Content-Type: application/json" \
      -d '{
            "isbn": "123",
            "title": "you",
            "director": {
              "first_name": "hello",
              "last_name": "joe"
            }
          }'
    ```
  * **Example Response:**
    ```json
    {
      "id": "1",
      "isbn": "123",
      "title": "you",
      "director": {
        "first_name": "hello",
        "last_name": "joe"
      }
    }
    ```

-----

### Delete a Movie

  * **URL:** `/movie/{id}`
  * **Method:** `DELETE`
  * **Description:** Deletes a movie from the collection using its ID.
  * **URL Parameters:**
      * `id` (string): The ID of the movie to delete.
  * **Example Request:**
    ```bash
    curl -X DELETE http://localhost:9090/movie/1
    ```
  * **Example Response:**
    ```json
    [
      {
        "id": "131735",
        "isbn": "123456",
        "title": "Vampire Diaries",
        "director": {
          "first_name": "um",
          "last_name": "idk"
        }
      }
    ]
    ```
    (Note: The response will show the remaining movies after the deletion.)

-----