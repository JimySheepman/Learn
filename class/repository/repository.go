package repository

import (
	"fmt"

	"github.com/JimySheepman/class/models"
)

type adminRepository struct {
}

func NewAdminRepository() *adminRepository {
	return &adminRepository{}
}

func (r *adminRepository) FindById(id int) error {
	fmt.Printf("%v", models.Users{})
	return nil
}

func (r *adminRepository) Create(name string) error {
	return nil
}

func (r *adminRepository) Update(name string) error {
	return r.update(name)
}

func (r *adminRepository) update(name string) error {
	return nil
}
