curl http://localhost:8080/query \
-H "Content-Type: application/json" \
-d '{ "query":"query findTodos {todos {text done user {name}}}"}' \