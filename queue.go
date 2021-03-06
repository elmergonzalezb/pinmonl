package main

import (
	"time"

	"github.com/pinmonl/pinmonl/config"
	"github.com/pinmonl/pinmonl/monl"
	"github.com/pinmonl/pinmonl/queue"
)

func initQueueManager(cfg *config.Config, ss stores, monl *monl.Monl) *queue.Manager {
	itv, _ := time.ParseDuration(cfg.Queue.Interval)

	return queue.NewManager(queue.ManagerOpts{
		Parallel: cfg.Queue.Parallel,
		Interval: itv,
		Monl:     monl,

		Store:    ss.store,
		Jobs:     ss.jobs,
		Pkgs:     ss.pkgs,
		Monls:    ss.monls,
		Pinls:    ss.pinls,
		Stats:    ss.stats,
		Substats: ss.substats,
	})
}
