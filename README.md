# Agnostic-api

The idea about this repo is to create a lot of implementations of the same API in different languages. 

But why, you might ask? Well, learning is way more fun if you have any context to it!

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

There might be a case, that phases **I** - **III** will happen all at the same time - to monitor the performance from the very start.

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
