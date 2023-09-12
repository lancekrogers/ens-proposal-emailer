package takehome

import (
	"context"
	"log"
	"tally-takehome/internal/email"
	monitor "tally-takehome/internal/monitoring"
	"tally-takehome/internal/store"
	"tally-takehome/internal/utils"
)

type ENSMonitoringService struct {
	Monitor monitor.Monitor
}

func NewENSMonitoringService(ctx context.Context, cfg *utils.Config) (*ENSMonitoringService, error) {

	db, err := store.NewBoltDBStore(utils.DB_PATH)
	if err != nil {
		return nil, err
	}

	if cfg == nil {
		log.Panic("Config files not loaded!")
	}

	monitor := monitor.NewMonitor(
		cfg.ENSGovernanceContract.Address,
		cfg.TallyApi,
		email.NewEmailClient(cfg.EmailSettings),
		db)

	return &ENSMonitoringService{
		Monitor: *monitor,
	}, nil
}

func (ens *ENSMonitoringService) Run(ctx context.Context) {
	ens.Monitor.StartMonitoring(ctx)
}
