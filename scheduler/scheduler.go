package scheduler

import (
	"context"
	"fmt"
	"time"

	"github.com/brevis-network/prover-network-bidder/client"
	"github.com/brevis-network/prover-network-bidder/config"
	"github.com/brevis-network/prover-network-bidder/dal"
	"github.com/brevis-network/prover-network-bidder/onchain"
	"github.com/celer-network/goutils/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/viper"
)

type Scheduler struct {
	*dal.DAL
	*onchain.ChainClient
	*client.ProverNetworkClient
	ruleConfig *config.RuleConfig
}

func NewScheduler(db *dal.DAL, c *onchain.ChainClient, p *client.ProverNetworkClient) (*Scheduler, error) {
	var ruleConfig *config.RuleConfig
	err := viper.UnmarshalKey(config.KeyRule, c)
	if err != nil {
		return nil, fmt.Errorf("UnmarshalKey err: %w", err)
	}
	return &Scheduler{db, c, p, ruleConfig}, nil
}

func (s *Scheduler) Start() {
	go s.scheduleAppRegister()
	go s.scheduleBid()
	go s.scheduleProve()
}

func (s *Scheduler) scheduleAppRegister() {
	for {
		time.Sleep(5 * time.Second)
		apps, err := s.FindNotRegisteredApps(context.Background())
		if err != nil {
			log.Errorf("FindNotRegisteredApps err: %s", err)
			continue
		}
		var elf []byte
		// TODO: retrieve elf from img url
		for _, app := range apps {
			err = s.RegisterApp(app.AppID, "", elf)
			if err != nil {
				log.Errorf("RegisterApp %s err: %s", app.AppID, err)
				continue
			}

			err = s.UpdateAppAsRegistered(context.Background(), app.AppID)
			if err != nil {
				log.Errorf("UpdateAppAsRegistered %s err: %s", app.AppID, err)
				continue
			}
		}
	}
}

func (s *Scheduler) scheduleBid() {
	for {
		time.Sleep(5 * time.Second)
		reqs, err := s.FindNotProcessedProofRequests(context.Background())
		if err != nil {
			log.Errorf("FindNotProcessedProofRequests err: %s", err)
			continue
		}
		var inputs [][]byte
		// TODO: get inputs from req.InputData or req.InputUrl
		for _, req := range reqs {
			_, pvDigest, err := s.EstimateCost(req.AppID, inputs)
			if err != nil {
				log.Errorf("EstimateCost %s err: %s", req.ReqID, err)
				continue
			}
			if common.BytesToHash(pvDigest) != common.HexToHash(req.PublicValuesDigest) {
				log.Errorf("pvDigest not match, %x vs %s", pvDigest, req.PublicValuesDigest)
				err = s.UpdateRequestAsProcessed(context.Background(), req.ReqID)
				if err != nil {
					log.Errorf("UpdateRequestAsProcessed %s err: %s", req.ReqID, err)
				}
				continue
			}
			// TODO compare cost with rule
			// TODO submit bid if pass rule

			err = s.UpdateRequestAsProcessed(context.Background(), req.ReqID)
			if err != nil {
				log.Errorf("UpdateRequestAsProcessed %s err: %s", req.ReqID, err)
			}
		}
	}
}

func (s *Scheduler) scheduleProve() {
	for {

		time.Sleep(5 * time.Second)
		// TODO
	}
}
