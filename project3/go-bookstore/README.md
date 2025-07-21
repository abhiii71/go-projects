install necessary packages
gorm for interacting with db 
➜  go-bookstore git:(project3) ✗ go get github.com/jinzhu/gorm               
go: downloading github.com/jinzhu/gorm v1.9.16
go: downloading github.com/jinzhu/inflection v1.0.0
go: added github.com/jinzhu/gorm v1.9.16
go: added github.com/jinzhu/inflection v1.0.0
➜  go-bookstore git:(project3) ✗ 


and for sql
go get github.com/jinzhu/gorm/dialects/mysql                               


and mux 
 go get github.com/gorilla/mux              



📘 1. Get all books
Endpoint: GET /book/

bash
Copy
Edit
curl -X GET http://localhost:9090/book/



🆕 2. Create a new book
Endpoint: POST /book/
Request body:

json
Copy
Edit
{
  "name": "The Alchemist",
  "author": "Paulo Coelho",
  "publication": "HarperOne"
}

curl command:

bash
Copy
Edit
curl -X POST http://localhost:9090/book/ \
  -H "Content-Type: application/json" \
  -d '{"name":"The Alchemist", "author":"Paulo Coelho", "publication":"HarperOne"}'

📗 3. Get a book by ID
Endpoint: GET /book/{id}

Example for id = 1:

bash
Copy
Edit
curl -X GET http://localhost:9090/book/1


✏️ 4. Update a book
Endpoint: PUT /book/{id}
Request body: (Updated fields)

Example (update book id = 1):

json
Copy
Edit
{
  "name": "The Alchemist - Updated",
  "author": "Paulo Coelho",
  "publication": "HarperCollins"
}
curl command:

bash
Copy
Edit
curl -X PUT http://localhost:9090/book/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"The Alchemist - Updated", "author":"Paulo Coelho", "publication":"HarperCollins"}'


🗑️ 5. Delete a book
Endpoint: DELETE /book/{id}

Example for id = 1:

bash
Copy
Edit
curl -X DELETE http://localhost:9090/book/1


making request using curl
get all
/book/

create book
/book/
req - type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}
name author publication


getbyid
/book/id

update bok
/book/id
aND THE REQUEST

DELTE BOOK
/BOOK/ID
