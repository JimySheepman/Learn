package services

import "github.com/JimySheepman/class/dto"

type adminRepository interface {
	FindById(id int) error
	Create(name string) error
	Update(name string) error
}

type adminServices struct {
	adminRepository adminRepository
}

func NewAdminServices(adminRepository adminRepository) *adminServices {
	return &adminServices{
		adminRepository: adminRepository,
	}
}

func (s *adminServices) Read(req *dto.RestRequest) error {
	if err := s.adminRepository.FindById(req.Id); err != nil {
		return err
	}

	return nil
}

func (s *adminServices) Write(req *dto.RestRequest) (*dto.RestResponse, error) {
	if err := s.adminRepository.Create(req.Name); err != nil {
		return nil, err
	}

	return &dto.RestResponse{}, nil
}

func (s *adminServices) AsyncRead(req *dto.RestRequest2) error {
	if err := s.adminRepository.FindById(req.Id); err != nil {
		return err
	}

	return nil
}

func (s *adminServices) AsyncWrite(req *dto.RestRequest2) (*dto.RestResponse2, error) {
	if err := s.adminRepository.Create(req.Name); err != nil {
		return nil, err
	}

	return &dto.RestResponse2{}, nil
}
