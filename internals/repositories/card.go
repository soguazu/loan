package repositories

import (
	"core_business/internals/core/domain"
	"core_business/internals/core/ports"
	"core_business/pkg/utils"
	"gorm.io/gorm"
)

type cardRepository struct {
	db *gorm.DB
}

// NewCardRepository creates a new instance card repository
func NewCardRepository(db *gorm.DB) ports.ICardRepository {
	return &cardRepository{
		db: db,
	}
}

func (c *cardRepository) GetByID(id string) (*domain.Card, error) {
	var card domain.Card
	if err := c.db.Where("id = ?", id).
		Preload("Transaction").
		Preload("Customer").
		First(&card).Error; err != nil {
		return nil, err
	}
	return &card, nil
}

func (c *cardRepository) GetCardByCompanyID(id string, pagination *utils.Pagination) (*utils.Pagination, error) {
	var cards []domain.Card
	if err := c.db.Scopes(utils.Paginate(cards, pagination, c.db)).
		Where("Company = ? AND Lock = ? 	AND Status = ? AND type = ?", id, false, "active", pagination.GetFilter()).
		Find(&cards).Error; err != nil {
		return nil, err
	}

	pagination.Rows = cards
	return pagination, nil
}

func (c *cardRepository) Get(pagination *utils.Pagination) (*utils.Pagination, error) {
	var cards []domain.Card
	if err := c.db.Scopes(utils.Paginate(cards, pagination, c.db)).Find(&cards).Error; err != nil {
		return nil, err
	}
	pagination.Rows = cards
	return pagination, nil
}

func (c *cardRepository) Persist(card *domain.Card) error {
	if card.ID.String() != "" {
		if err := c.db.Save(card).Error; err != nil {
			return err
		}
		return nil
	}
	if err := c.db.Create(&card).Error; err != nil {
		return err
	}
	return nil
}

func (c *cardRepository) Delete(id string) error {
	if err := c.db.Where("id = ?", id).Delete(&domain.Card{}).Error; err != nil {
		return err
	}
	return nil
}

func (c *cardRepository) DeleteAll() error {
	if err := c.db.Exec("DELETE FROM cards").Error; err != nil {
		return err
	}
	return nil
}

func (c *cardRepository) WithTx(tx *gorm.DB) ports.ICardRepository {
	return NewCardRepository(tx)
}
