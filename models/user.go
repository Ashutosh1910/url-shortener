package models

type Hash struct {
	Code string
}

type User struct {
	ID       int
	Name     string
	Passowrd Hash
	IsAdmin  bool
	Active   bool
}

// func encrypt_password(password string) Hash{

// }

// func decrypt_hash(hash Hash) string{

// }

// func (user *User)Check_password(password string ) bool {

// }

// func (user *User) Set_password(password string){

// }

// func (user *User) Save(){

// }
// func (user *User) Delete() {

// }
// func FindUser(id int) *User{

// }
// func FindUserByName(name string)*User{

// }
