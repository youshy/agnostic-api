# Agnostic-api

Aka - create the same thing few times and check how fast it is. And then refactor and refactor again.

I was curious if there's any difference between the same API implementation in different languages - therefore, I present to you **Agnostic API**! (I know, it might not be the best name ever).

## Two APIs

The initial model was only to test the Read/Write actions using the DB, but after a chat with a friend, I came to realisation, that this is rather stupid - the differences in performance won't be that big to even care (we'll see) and also - you can tweak the performance of each API to be superfast.

Therefore, small change of plans. There will be two APIs - one CRUD with DB (in this case **Mongo**), second one: **--To-Be-Updated--** probably get the request, kick off encryption on something and formulate JSON and return to the client.

## What is the model

model: **User**

```
Firstname: string
Lastname: string
Age: integer
```

Routes:

**GET** `/users` - Gets all the users

**GET** `/users/?firstname={firstname}&lastname={lastname}` - Gets user with `firstname` and `lastname`

**POST** `/users` - Create a new user

**PUT** `/users/?firstname={firstname}&lastname={lastname}` - Update a user with `firstname` and `lastname`

**DELETE** `/users/?firstname={firstname}&lastname={lastname}` - Delete a user with `firstname` and `lastname`

Routes are specifically designed to not to use `Id`. This gives an extra flavour to the task as not often you have to look for something using two parameters in the route. Probably, in **phase V**, I will implement the API with `Id` as a param, not as a query.

## Phases

* **Phase I** - Implement the API using languages from table below with MongoDB connection.

* **Phase II** - Implement `Docker` images for every language (to make the startup nicer) and `docker-compose` to run all the API's on different ports simultaneously.

* **Phase III** - Create performance checker using either `Go` or `Python` (or both actually). The idea is to shoot request to each API within 5 second intervals and to check the performance of every endpoint from API's consumer point of view.

* **Phase IV** - Implement the API using languages from table below with PostgreSQL connection.

* **Phase IV** - Reimplement Docker and tweak **phase III** if needed.

* **Phase V** - Implement second API (probably with streams to check file uploads) and do everything again. 

To get all the phases done it'll probably take some time but - that's a long-term project here as well.

There might be a case, that phases **I** - **III** will happen all at the same time - to monitor the performance from the very start. And also - phases might differ from language to language - it might happen, that writing CRUD API in one of the languages belwo might be impossible to do - then, the phases will move.

## Languages used in this repo

* Empty - not done
* `I` - in progress
* `X` - done

| Language | MongoDb | PostgreSQL | Notes |
|---|---|---|---|
| NodeJS      | X |   |   |
| Go          | X |   |   |
| Python      | X |   |   | 
| F#          | I |   | (Small issues) | 
| Java        | X |   |   |   
| C#          | I |   | (Basically done, unable to compile on my Mac) |
| C++         |   |   |   |
| Rust        | I |   |   |
| PHP         | I |   |   |
| Ruby        | I |   |   |
| Haskell     |   |   |   |
| Objective-C |   |   |   |
| Clojure     |   |   |   |
| Elixir      |   |   |   |
| Scala       |   |   |   |
| C           |   |   |   |
| Lua         |   |   |   |
| Julia         |   |   |   |

## Can I add my own submission?

Create a pull request with your language submission according to the model. 

Create a folder with your language name in the root directory, add `mongo` and `psql` folders and start digging!
