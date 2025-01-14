package repository

// sql queries for user repository.
const (
	// signUpQuery is a query to sign up a user.
	signUpQuery = `
		INSERT INTO users (username, password, email) VALUES ($1, $2, $3) RETURNING id;`

	// getUserByEmailQuery is a query for getting user by email.
	getUserByEmailQuery = `
		SELECT 
			id, 
			email,
			username, 
			password
		FROM users
		WHERE email = $1;`
)
