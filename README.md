# Agnostic-api

The idea about this repo is to create a lot of implementations of the same API in different languages. 

But why, you might ask? Well, learning is way more fun if you have any context to it!

## What is the model

Let's start simple:

```
model: User

Firstname: string
Lastname: string
Age: integer
```

Routes:

**GET** `/users` - Gets all the users

**GET** `/user/?firstname={firstname}&lastname={lastname}` - Gets user with `firstname` and `lastname`

**POST** `/user` - Create a new user

**PUT** `/user/?firstname={firstname}&lastname={lastname}` - Update a user with `firstname` and `lastname`

**DELETE** `/user/?firstname={firstname}&lastname={lastname}` - Delete a user with `firstname` and `lastname`

I know the routes are a bit sketchy now, but let's keep them like this. The goal of this repo is to cover as much ground as possible.

Database - use any - ideally, there will be an implementation for MongoDB and PostgreSQL in all examples.

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
| Java        |   |   |   |   
| C#          |   |   |   |
| C++         |   |   |   |
| Rust        |   |   |   |
| PHP         |   |   |   |
| Ruby        |   |   |   |
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

All APIs should be using port `3000`.
