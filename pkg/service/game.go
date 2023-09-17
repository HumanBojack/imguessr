package service

import (
	"bytes"
	"fmt"
	"image"
	"imguessr/pkg/domain"
	"time"

	_ "image/jpeg"
	_ "image/png"

	"github.com/google/uuid"
)

type gameSvc struct {
	DB domain.GameDB
}

func NewGameSvc(db domain.GameDB) domain.GameSvc {
	return gameSvc{
		DB: db,
	}
}

func (gs gameSvc) CreateGame(g *domain.Game) error {
	g.ID = uuid.New().String()
	g.DateTime = time.Now().UTC().Unix()

	return gs.DB.CreateGame(g)
}

// Check that the frequency is between 1 minute and 1 day (1440 minutes)
func (gs gameSvc) VerifyFrequency(frequency int) error {
	min := 0
	max := 1440
	if (frequency <= min) || (frequency > max) {
		return fmt.Errorf("Frequency (%v) must be between %v and %v", frequency, min, max)
	}

	return nil
}

// Check that the steps is between 3 and 30
func (gs gameSvc) VerifySteps(steps int) error {
	min := 3
	max := 30
	if (steps < min) || (steps > max) {
		return fmt.Errorf("Steps (%v) must be between %v and %v", steps, min, max)
	}

	return nil
}

// Check that the hiderType is one of the valid types:
// "pixels", "blur", "chunks", "zoom"
func (gs gameSvc) VerifyHiderType(hiderType string) error {
	validTypes := []string{"pixels", "blur", "chunks", "zoom"}

	for _, t := range validTypes {
		if t == hiderType {
			return nil
		}
	}

	return fmt.Errorf("HiderType (%v) must be one of the following: %v", hiderType, validTypes)
}

// Verify that the image is valid png or jpeg and not too large
func (gs gameSvc) VerifyImage(img []byte) error {
	// Verify that the image is not too large
	maxSize := 5000000
	if len(img) > maxSize {
		return fmt.Errorf("image is too large (must be smaller than %v): %v bytes", maxSize, len(img))
	}

	// Verify that the image is valid png or jpeg
	_, _, err := image.Decode(bytes.NewReader(img))
	if err != nil {
		return fmt.Errorf("image is not a valid png or jpeg: %v", err)
	}

	return nil
}
