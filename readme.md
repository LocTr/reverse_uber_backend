

# reverse-uber-be

Backend project using Go.

## Checklist

- [ ] Set up flavors
- [x] Setting up database
- [ ] Add models
- [ ] Seed data
- [ ] Configurate pipeline
- [ ] Deploy


cmd/web: main package, which contains the main entry file, routes, and middleware.

driver: connects to MongoDB Atlas.

handlers: processes user requests.

modules: additional packages for specific functionality.

auth: generates and parses jwt(Json Web Token) tokens.

config: reusable struct fields.

database: interacts with collections and injects results to handlers.

encrypt: handles hashing and verification of user passwords.

model: holds user details and is crucial for the application.
