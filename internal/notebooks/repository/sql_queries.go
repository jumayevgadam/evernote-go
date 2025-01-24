package repository

// sql queries for notebooks.
const (
	// addNotebookQuery is.
	addNotebookQuery = `
		INSERT INTO notebooks (
			user_id,
			name,
			is_shared
		) VALUES (
		 	$1,
			$2,
			$3
		) RETURNING id;`

	// countNotebooksByUserQuery is.
	countNotebooksByUserQuery = `
		SELECT COUNT(user_id)
		FROM notebooks
		WHERE user_id = $1;`

	// listNotebooksQuery is.
	listNotebooksQuery = `
		SELECT 
			id,
			user_id,
			name,
			is_shared,
			created_at,
			updated_at
		FROM notebooks
		WHERE user_id = $1
		ORDER BY created_at DESC
		OFFSET $2 LIMIT $3;`
)
