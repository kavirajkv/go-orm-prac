package db

import "github.com/kavirajkv/go-orm-prac/pkg/users"

var (
	Userservice users.Service = nil
)

func InitServices() {
	db := GetDB()

	userrepo := users.NewRepo(db)
	Userservice = users.Newservice(userrepo)

}
