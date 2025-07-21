package repository

import (
	"fmt"
	"github.com/Horronyt/marketplace"
	"github.com/jmoiron/sqlx"
)

const (
	listingTableName = "listings"
	userTableName    = "users"
)

type listingPostgres struct {
	db *sqlx.DB
}

func NewListingPostgres(db *sqlx.DB) *listingPostgres {
	return &listingPostgres{db: db}
}

func (s *listingPostgres) Create(userId int, listing marketplace.Listing) (int, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createListingQuery := fmt.Sprintf("INSERT INTO %s(user_id, title, description, img_path, price) VALUES($1,$2,$3,$4,$5) RETURNING id;", listingTableName)
	row := tx.QueryRow(createListingQuery, userId, listing.Title, listing.Description, listing.Img_path, listing.Price)
	if err = row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (s *listingPostgres) GetAll(userId int) ([]marketplace.ListingOutputFormat, error) {
	var listings []marketplace.ListingOutputFormat
	query := fmt.Sprintf("SELECT l.title, l.description,l.img_path,l.price,u.username, CASE WHEN u.id = $1 THEN TRUE ELSE FALSE END AS belonging FROM %s AS l JOIN %s AS u ON u.id = l.user_id ORDER BY l.id DESC", listingTableName, userTableName)
	err := s.db.Select(&listings, query, userId)

	return listings, err
}

func (s *listingPostgres) GetAllAnonymously() ([]marketplace.ListingOutputFormatAnon, error) {
	var listings []marketplace.ListingOutputFormatAnon
	query := fmt.Sprintf("SELECT l.title, l.description,l.img_path,l.price,u.username FROM %s AS l JOIN %s AS u ON u.id = l.user_id ORDER BY l.id DESC", listingTableName, userTableName)
	err := s.db.Select(&listings, query)

	return listings, err
}
