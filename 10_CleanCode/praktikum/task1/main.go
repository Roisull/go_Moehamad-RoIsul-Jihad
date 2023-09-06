package main

/** ada 4 kekurangan
	1. Penamaan variabel
	2. penamaan Field Struct
	3. Penamaan Method
	4. Pemberian komentar 
	5. Tidak mendeklarasikan main method
	*/ 

	// penamaan field struct harus diawali dengan huruf besar, karena mewakili infomasi pengguna
type user struct {

id int

username int

password int

}

	//  penamaan field struct harus diawali dengan huruf besar, jika ada 2 suku kata maka kata berikutnya diawali dengan huruf besar juga
	// penulisan ini biasa disebut dengan CamelCase
type userservice struct {

t []user

}

	// function getAllUsers harus menggunakan CamelCase karena menunjuk pada method
func (u userservice) getallusers() []user {

return u.t

}

	// sama dengan getAllUsers, getUserById juga harus menggunakan CamelCase karena menunjuk pada method
func (u userservice) getuserbyid(id int) user {

for _, r := range u.t {

if id == r.id {

return r

}

}


return user{}

}