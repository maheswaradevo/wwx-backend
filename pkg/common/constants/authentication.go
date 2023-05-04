package constants

const (
	CheckUsernameQuery = "SELECT id, username, role, password FROM users WHERE username = ?;"
)

var (
	LoginSuccess     = "Login Success!"
	PasswordMismatch = "Masukkan password yang sesuai"
	NoUsernameExists = "Username tidak ditemukan"
)
