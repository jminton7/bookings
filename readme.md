# Bookings and Reservations

The repository for [Building Modern Web Applications with Go](https://www.udemy.com/course/building-modern-web-applications-with-go/?referralCode=0415FB906223F10C6800).

- Built in Go version 1.15
- Uses the [chi router](github.com/go-chi/chi)
- Uses [alex edwards scs session management](github.com/alexedwards/scs)
- Uses [nosurf](github.com/justinas/nosurf)

# Useful commands

### Running

- running, terminal from golang-bookings and run "go run ./cmd/web/." -> with existance of run.bat -> ".\run.bat"

### Testing

- testing, terminal from golang-bookings/cmd/web and run "go test -v"
- test coverage, terminal from golang-bookings/cmd/web and run "go test -cover"
- test coverage with http page showing where tests are missing, terminal golang-bookings/cmd/web and run "go test -coverprofile=coverage.out && go tool cover -html=coverage.out"

### Soda

- To create a migration file. "soda generate fizz CreateUserTable"
- To start the migration. "soda migrate" / "soda migrate down"
- Reset the table. "soda reset"