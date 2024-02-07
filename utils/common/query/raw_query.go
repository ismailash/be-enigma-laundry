package query

const (
	GetUserByUsername = `SELECT id, email, username, password, role FROM users WHERE username = $1 OR email = $1`
)
