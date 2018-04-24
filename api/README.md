Customer Review API:
====================

*JSON* is in use as a format.
A response has the following format:
```
{
  "success": true | false,
  "data": {} | null | [{}],
  "errors": string | null,
  "meta": {} | null
}
```

#### Create a new review

*AUTH*: anonymous

Creates a new review, but score calculation will be done asynchronously.
 
```bash
POST /reviews
Status: 201 Created
```

JSON-in:
```json
{
	"name":"bob",
	"email":"bob@mail.com",
	"content":"I like reading."
}
```

JSON-out:
```json
{
    "success": true,
    "data": {
        "id": "41a4640c-3280-4555-9f51-01a794644b89",
        "name": "bob",
        "email": "bob@mail.com",
        "content": "I like reading.",
        "published": false,
        "score": null,
        "category": null,
        "created_at": "2018-04-24T11:31:38",
        "updated_at": "2018-04-24T11:31:38"
    },
    "error": null,
    "meta": null
}
```

Validation:
- *name* - required | max:80 characters
- *email* - required | valid email | max:80 characters
- *content* - required | max:2000 characters


#### Update a review

*AUTH*: owner withing the same session.

Updates just limited number of field of a review.
```bash
PUT /reviews/{id}
Status 200 OK
``` 

JSON-in:
```json
{
	"name":"Alice",
	"published": true,
	"content": "some content"
}
```

Validation:
- *name* - optional | max:80 characters
- *content* - optional | max:2000 characters
- *published* - optional

#### Get a review

*AUTH*: anonymous

Returns a review by ID.
```bash
GET /reviews/{id}
Status: 200 OK
```

JSON-out:
```json
{
    "success": true,
    "data": {
        "id": "41a4640c-3280-4555-9f51-01a794644b89",
        "name": "bob",
        "email": "bob@mail.com",
        "content": "I like reading.",
        "published": false,
        "score": 88,
        "category": "positive",
        "created_at": "2018-04-24T06:41:18",
        "updated_at": "2018-04-24T06:47:45"
    },
    "error": null,
    "meta": null
}
```

#### Get list of reviews:

*AUTH*: anonymous

Returns a list of reviews. May apply some criteria:

```bash
GET /reviews?limit={limit}&offset={offset}&category={positive|negative}&published={true|false}
Status 200 OK
```

JSON-out:
```json
{
    "success": true,
    "data": [
        {
            "id": "41a4640c-3280-4555-9f51-01a794644b89",
            "name": "bob",
            "email": "bob@mail.com",
            "content": "I like reading.",
            "published": false,
            "score": 88,
            "category": "positive",
            "created_at": "2018-04-24T06:41:18",
            "updated_at": "2018-04-24T06:47:45"
        },
        {
            "id": "41a4640c-3280-4555-9f51-01a794644b86",
            "name": "alice",
            "email": "alice@mail.com",
            "content": "I don't like reading.",
            "published": true,
            "score": 44,
            "category": "negative",
            "created_at": "2018-04-24T06:41:18",
            "updated_at": "2018-04-24T06:47:45"
        },
    ],
    "error": null,
    "meta": {
      "total": 22,
      "count": 2,
      "limit": 10,
      "offset": 20
    }
}
```

Filter params:
- *limit* - optional | default:10
- *offset* - optional | default:0
- *category* - optional | in:positive,negative
- *published* - optional | in:true,false
