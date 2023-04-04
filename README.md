# USAGE

[![Go](https://github.com/ani-social/profiles/actions/workflows/go.yml/badge.svg)](https://github.com/ani-social/profiles/actions/workflows/go.yml)

## GET all users
<code>curl http://localhost:8080/users -H "Content-Type: application/json" -H "User-Agent: Morty's client"</code>

## POST a new user
<code>curl -X POST http://localhost:8080/users -H "Content-Type: application/json" -H "User-Agent: Rick's client" \
-d '{"username":"jerry","avatar":"https://rickandmortyapi.com/api/character/avatar/5.jpeg","profile":{"name":"Jerry Smith","image":"https://rickandmortyapi.com/api/character/avatar/5.jpeg","bio":"A simple man with simple needs.","philosophy":"rock the boat.","achievements":[{"title":"Employee of the Month","description":"Once a year."}],"socialLinks":[{"name":"Twitter","icon":"fa fa-twitter","url":"https://twitter.com/jerrysmith"},{"name":"Facebook","icon":"fa fa-facebook","url":"https://facebook.com/jerrysmith"}]}}'</code>

## GET a specific user by username
<code>curl http://localhost:8080/users/morty -H "Content-Type: application/json" -H "User-Agent: Rick's client"</code>

## UPDATE an existing user by username
<code>curl -X PUT http://localhost:8080/users/morty -H "Content-Type: application/json" -H "User-Agent: Morty's client" \
-d '{"username":"morty","avatar":"https://rickandmortyapi.com/api/character/avatar/2.jpeg","profile":{"name":"Morty Smith","image":"https://rickandmortyapi.com/api/character/avatar/2.jpeg","bio":"A timid and easily-flustered boy who goes on adventures with his grandfather.","philosophy":"Just keep your head down and hope for the best.","achievements":[{"title":"Survived a gazorpazorpfield attack","description":"It was brutal."}],"socialLinks":[{"name":"Instagram","icon":"fa fa-instagram","url":"https://instagram.com/morty"}]}}'</code>

## DELETE a user by username
<code>curl -X DELETE http://localhost:8080/users/jerry -H "Content-Type: application/json" -H "User-Agent: Rick's client"</code>
