# books-list-golang

This is the book APIs.

## Sample

### AllBooks

```bash
GET : http://localhost:8000/books
```

### GetBook

```bash
GET : http://localhost:8000/books/3
```

### AllBooks

```bash
GET : http://localhost:8000/books
```

### AddBook

```bash
POST : http://localhost:8000/books

{
  "title": "C++ is old",
  "author": "Mr. C++",
  "year": "2024"
}
```

### UpdateBook

```bash
PUT : http://localhost:8000/books

{
  "id": 3,
  "title": "Golang routers",
  "author": "Mr. Router2",
  "year": "2012"
}
```

### RemoveBook
```bash
DELETE : http://localhost:8000/books/8
```
