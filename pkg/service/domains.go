package service

import (
	"github.com/Kushkaftar/geo_support/modelsStruct"
	"github.com/Kushkaftar/geo_support/pkg/repository"
	"github.com/Kushkaftar/geo_support/pkg/workingLogic/folders"
	"github.com/sirupsen/logrus"
	"os"
)

type ReturnDomainsService struct {
	repo repository.Domains
}

func NewReturnDomainsService(repo repository.Domains) *ReturnDomainsService {
	return &ReturnDomainsService{repo: repo}
}

func (rds *ReturnDomainsService) GetAllDomains() ([]modelsStruct.Domain, error) {

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

func (rds ReturnDomainsService) CheckDomain(domain string) (int, error) {
	return rds.repo.CheckDomain(domain)
}

func (rds ReturnDomainsService) InsertDomain(domain string) error {
	return rds.repo.InsertDomain(domain)
}

func (rds ReturnDomainsService) SetFlag(flag, id int) error {
	return rds.repo.SetFlag(flag, id)
}
