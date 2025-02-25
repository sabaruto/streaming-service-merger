package authorisation

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"time"

	"github.com/google/uuid"
	"github.com/sabaruto/streaming-service-merger/backend/lib/authorisation/postgres/models"
)

func (s *server) getLatestToken(ctx context.Context, customerID uuid.UUID) (*models.TokenStore, error) {
	var (
		code         string
		customer_id  string
		expire_after time.Time
		newStore     *models.TokenStore
	)
	err := s.db.QueryRowContext(
		ctx,
		"SELECT (code, customer_id, expire_after) FROM token_store WHERE expire_after > 'yesterday' AND customer_id = $1",
		customerID.String(),
	).Scan(&code, &customer_id, &expire_after)

	switch err {
	case sql.ErrNoRows:
		if newStore, err = s.newToken(ctx, customerID); err != nil {
			return nil, err
		}
	case nil:
		newStore = &models.TokenStore{
			Code:        code,
			CustomerID:  customerID,
			ExpireAfter: expire_after,
		}
	default:
		return nil, err
	}
	return newStore, nil
}

func (s *server) newToken(ctx context.Context, customerID uuid.UUID) (*models.TokenStore, error) {
	bytes := make([]byte, 64)
	_, err := rand.Read(bytes)
	if err != nil {
		return nil, err
	}

	code := hex.EncodeToString(bytes)

	newStore := &models.TokenStore{
		Code:        code,
		CustomerID:  customerID,
		ExpireAfter: time.Now().Add(7 * 24 * time.Hour),
	}

	err = newStore.Save(ctx, s.db)
	if err != nil {
		return nil, err
	}

	return newStore, nil
}
