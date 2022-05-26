curl -X POST http://localhost:8080/query  \
-H 'Content-Type: application/json' \
-d '{ "query": "mutation createTodo {createTodo(input: { text: \"todo\", userId: \"1\" }) {user {id}text done}}" }'