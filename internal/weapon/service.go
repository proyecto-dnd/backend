package weapon

import (
	"fmt"
	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/dto"
)

type service struct {
	weaponRepo WeaponRepository
}

func NewWeaponService(weaponRepo WeaponRepository) WeaponService {
	return &service{weaponRepo: weaponRepo}
}

func (s *service) CreateWeapon(weaponDto dto.CreateWeaponDto) (domain.Weapon, error) {
	weaponDomain := domain.Weapon{
		WeaponType:      weaponDto.WeaponType,
		Name:            weaponDto.Name,
		Weight:          weaponDto.Weight,
		Price:           weaponDto.Price,
		Category:        weaponDto.Category,
		Reach:           weaponDto.Reach,
		Description:     weaponDto.Description,
		Damage:          weaponDto.Damage,
		VersatileDamage: weaponDto.VersatileDamage,
		Ammunition:      weaponDto.Ammunition,
		DamageType:      weaponDto.DamageType,
		Basic:           weaponDto.Basic,
	}

	createdWeapon, err := s.weaponRepo.Create(weaponDomain)
	if err != nil {
		return domain.Weapon{}, err
	}

	return createdWeapon, nil
}

func (s *service) GetAllWeapons() ([]domain.Weapon, error) {
	weapons, err := s.weaponRepo.GetAllWeapons()
	if err != nil {
		return nil, err
	}

	return weapons, nil
}

func (s *service) GetWeaponById(id int) (domain.Weapon, error) {
	weapon, err := s.weaponRepo.GetWeaponById(id)
	if err != nil {
		return domain.Weapon{}, err
	}

	return weapon, nil
}

func (s *service) UpdateWeapon(weaponDto dto.CreateWeaponDto, id int) (domain.Weapon, error) {
	weaponDomain := domain.Weapon{
		WeaponType:      weaponDto.WeaponType,
		Name:            weaponDto.Name,
		Weight:          weaponDto.Weight,
		Price:           weaponDto.Price,
		Category:        weaponDto.Category,
		Reach:           weaponDto.Reach,
		Description:     weaponDto.Description,
		Damage:          weaponDto.Damage,
		VersatileDamage: weaponDto.VersatileDamage,
		Ammunition:      weaponDto.Ammunition,
		DamageType:      weaponDto.DamageType,
		Basic:           weaponDto.Basic,
	}

	updatedWeapon, err := s.weaponRepo.UpdateWeapon(weaponDomain, id)
	if err != nil {
		fmt.Println(err)
		return domain.Weapon{}, err
	}

	return updatedWeapon, nil
}

func (s *service) DeleteWeapon(id int) error {
	return s.weaponRepo.DeleteWeapon(id)
}
