package common

const (
	UserCreate              = `INSERT INTO users (id, username, password, is_active, role, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	UserListPaging          = `SELECT id, username, is_active, role, created_at, updated_at FROM users ORDER BY created_at DESC LIMIT $1 OFFSET $2`
	UserList                = `SELECT id, username, is_active, role, created_at, updated_at FROM users ORDER BY created_at`
	UserGet                 = `SELECT id, username, is_active, role, created_at, updated_at FROM users WHERE id = $1`
	UserGetUsernamePassword = `SELECT id, username, password, role, is_active FROM users WHERE username = $1`
	UserUpdate              = `UPDATE users SET username = $2, password = $3, is_active = $4, role =$5, updated_at = $6 WHERE id = $1`
	UserActivate            = `UPDATE users SET is_active = true WHERE reset_token = $1`
	UserResetToken          = `UPDATE users SET reset_token = $2, updated_at = $3 WHERE id = $1`
	UserDelete              = `DELETE FROM users WHERE id = $1`
)
