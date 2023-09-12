package monitoring

import (
	"context"
	"fmt"
	"log"
	"tally-takehome/internal/email"
	"tally-takehome/internal/store"
	"tally-takehome/internal/tally"
	"time"
)

type Monitor struct {
	GovenorAddress string
	TallyApi       tally.TallyApi
	EmailClient    *email.EmailClient
	Store          *store.BoltDBStore
}

func NewMonitor(govenorAddress string, tallyApi tally.TallyApi, emailClient *email.EmailClient, store *store.BoltDBStore) *Monitor {
	return &Monitor{
		GovenorAddress: govenorAddress,
		TallyApi:       tallyApi,
		EmailClient:    emailClient,
		Store:          store,
	}
}

func (m *Monitor) StartMonitoring(ctx context.Context) {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			// Context was canceled or expired; stop the monitoring
			return
		case <-ticker.C:
			// Run your function
			if err := m.CheckLastBlock(ctx); err != nil {
				// Handle error
				fmt.Printf("Error occurred while checking last block: %v\n", err)
			}
		}
	}
}

func (m *Monitor) CheckLastBlock(ctx context.Context) error {
	lastDbBlock, err := m.Store.GetLastProcessedBlock()
	if err != nil {
		if err != store.ErrNoLastProcessedBlock {
			return err
		}
	}

	lastProposal, err := m.TallyApi.GetLastProposal(ctx, m.GovenorAddress)
	if err != nil {
		return err
	}
	lastProposalBlock := lastProposal.Block.Number
	if lastProposalBlock > lastDbBlock {
		err := m.SendProposalEmailNotification(lastProposalBlock, lastProposal.ID, lastProposal.Title)
		if err != nil {
			log.Printf("Experienced an error during email notification for proposal %v with error %v", lastProposal.ID, err)
		}
	}

	return nil
}

func (m *Monitor) SendProposalEmailNotification(blockNumber int, propId string, propTitle string) error {
	// Attempting to send email notification
	err := m.EmailClient.SendEmail(propId, propTitle)
	if err != nil {
		// if an error is experienced return the error so that the email will be retried
		return err
	}
	// Set the current proposal block number as the last processed blocknumber
	err = m.Store.SetLastProcessedBlock(blockNumber)
	if err != nil {
		return err
	}
	return nil
}
