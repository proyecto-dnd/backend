package weapon

import "github.com/proyecto-dnd/backend/internal/domain"

type service struct {
	repo RepositoryWeapon
}

func NewServiceWeapon(repo RepositoryWeapon) ServiceWeapon {
	return &service{repo: repo}
}

// Create implements ServiceWeapon.
func (s *service) Create(weapon domain.Weapon) (domain.Weapon, error) {
	newWeapon, err := s.repo.Create(weapon)
	if err != nil {
		return domain.Weapon{}, err
	}
	return newWeapon, nil
}

// Delete implements ServiceWeapon.
func (s *service) Delete(id int64) error {
	err := s.repo.Delete(id)
    if err!= nil {
        return err
    }
    return nil
}

// GetAll implements ServiceWeapon.
func (s *service) GetAll() ([]domain.Weapon, error) {
	weapons, err := s.repo.GetAll()
    if err!= nil {
        return []domain.Weapon{}, err
    }
    return weapons, nil
}

// GetByCampaignId implements ServiceWeapon.
func (s *service) GetByCampaignId(campaignId int64) ([]domain.Weapon, error) {
	weapons, err := s.repo.GetByCampaignId(campaignId)
    if err!= nil {
        return []domain.Weapon{}, err
    }
    return weapons, nil
}

// GetById implements ServiceWeapon.
func (s *service) GetById(id int64) (domain.Weapon, error) {
	weapon, err := s.repo.GetById(id)
    if err!= nil {
        return domain.Weapon{}, err
    }
    return weapon, nil
}

// Update implements ServiceWeapon.
func (s *service) Update(weapon domain.Weapon) (domain.Weapon, error) {
	weapon, err := s.repo.Update(weapon)
    if err!= nil {
        return domain.Weapon{}, err
    }
    return weapon, nil
}

