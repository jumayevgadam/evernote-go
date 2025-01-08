package repository

// sql queries for user repository.
const (
	// signUpQuery is a query to sign up a user.
	signUpQuery = `
		INSERT INTO users (username, password, email) VALUES ($1, $2, $3) RETURNING id;`
)