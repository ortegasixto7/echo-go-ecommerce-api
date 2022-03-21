package admin

type IAdminPersistence interface {
	Save(AdminUser)
	Update(AdminUser)
	GetById(id string) AdminUser
	GetByUserName(username string) AdminUser
}
