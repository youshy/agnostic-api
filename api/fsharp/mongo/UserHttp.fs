namespace Todos.Http

open Giraffe
open Microsoft.AspNetCore.Http

module UserHttp =
  let handlers : HttpFunc -> HttpContext -> HttpFuncResult =
    choose [
      GET >=> route "/users" >=>
        fun next context ->
          text "Read" next context

      GET >=> route
