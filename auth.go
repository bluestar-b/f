package main

func authenticateUser(username, password string) bool {
	storedPassword, ok := AuthCredentials[username]
	return ok && storedPassword == password
}
