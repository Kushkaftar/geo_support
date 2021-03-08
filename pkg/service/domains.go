package service

import (
	"fmt"
	"os"
	"path"

	"github.com/Kushkaftar/geo_support/modelsstruct"
	"github.com/Kushkaftar/geo_support/pkg/repository"
	"github.com/Kushkaftar/geo_support/pkg/workingLogic/folders"
	"github.com/sirupsen/logrus"
)

// ReturnDomainsService ...
type ReturnDomainsService struct {
	repo repository.Domains
}

// NewReturnDomainsService ...
func NewReturnDomainsService(repo repository.Domains) *ReturnDomainsService {
	return &ReturnDomainsService{repo: repo}
}

// GetAllDomains ...
func (rds *ReturnDomainsService) GetAllDomains() ([]modelsstruct.Domain, error) {

	// лапша стайл(((
	path := folders.NewPath(os.Getenv("PATH_REPOSITORY"))

	domains, err := path.GetFolders()
	if err != nil {
		return nil, err
	}
	for _, value := range domains {
		//fmt.Println(value)
		req, err := rds.CheckDomain(value)
		if err != nil {
			return nil, err
		}
		if req == 0 {
			err = rds.InsertDomain(value)
			if err != nil {
				logrus.Error(err)
				return nil, err
			}
		}
		//fmt.Println(req)
	}
	return rds.repo.GetAllDomains()
}

// CheckDomain ...
func (rds ReturnDomainsService) CheckDomain(domain string) (int, error) {
	return rds.repo.CheckDomain(domain)
}

// InsertDomain ...
func (rds ReturnDomainsService) InsertDomain(domain string) error {
	return rds.repo.InsertDomain(domain)
}

// PreSetFlag ...
func (rds ReturnDomainsService) PreSetFlag(flag, id int) error {

	err := rds.repo.SetFlag(flag, id)
	if err != nil {
		return err
	}

	if flag == 1 {
		domain, err := rds.repo.GetDomain(id)
		if err != nil {
			return err
		}

		pathToFolders := folders.NewPath(path.Join(os.Getenv("PATH_REPOSITORY"), domain.Name))
		fmt.Println(pathToFolders)
		folders, err := pathToFolders.GetFolders()
		if err != nil {
			return err
		}

		for _, f := range folders {
			fmt.Println(f)
		}
	}

	return nil
}

// SetFlag ...
func (rds ReturnDomainsService) SetFlag(flag, id int) error {
	return rds.repo.SetFlag(flag, id)
}

// GetDomain ...
func (rds ReturnDomainsService) GetDomain(id int) (modelsstruct.Domain, error) {
	return rds.repo.GetDomain(id)
}
