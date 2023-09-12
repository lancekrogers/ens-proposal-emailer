package monitoring

import (
	"context"
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

// StartMonitoring method     Kick off monitoring every 10 seconds
func (m *Monitor) StartMonitoring(ctx context.Context) {
	log.Println("Starting ENS proposal monitor")
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			// Context was canceled or expired; stop the monitoring
			log.Println("Monitor is stopping")
			return
		case <-ticker.C:
			// Run your function
			if err := m.CheckLastBlock(ctx); err != nil {
				log.Printf("Error occurred while checking last block: %v\n", err)
			}
		}
	}
}

// CheckLastBlock method    Compares proposal block to the block in the db and sends notification if the proposal is newer
func (m *Monitor) CheckLastBlock(ctx context.Context) error {

	lastDbBlock, err := m.Store.GetLastProcessedBlock()
	if err != nil {
		if err != store.ErrNoLastProcessedBlock {
			log.Printf("Error while getting last processed block: %v", err)
			return err
		}
	}

	lastProposal, err := m.TallyApi.GetLastProposal(ctx, m.GovenorAddress)
	if err != nil {
		log.Printf("Error while getting last proposal: %v", err)
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

// SendProposalEmailNotification method     Sends the proposal email and updates the database with the latest proposals blocknumber
func (m *Monitor) SendProposalEmailNotification(blockNumber int, propId string, propTitle string) error {
	// Attempting to send email notification
	err := m.EmailClient.SendEmail(propId, propTitle)
	if err != nil {
		// if an error is experienced return the error so that the email will be retried
		return err
	}
	log.Printf("Email notification sent for %v", propId)
	// Set the current proposal block number as the last processed blocknumber
	err = m.Store.SetLastProcessedBlock(blockNumber)
	if err != nil {
		return err
	}
	return nil
}
