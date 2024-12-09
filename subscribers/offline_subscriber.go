package subscribers

import (
	"log"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"

	"oracle/contracts"
	"oracle/di"
)

type OfflineSubscriber struct {
	contract      *contracts.Register
	oracleAddress []common.Address
	sink          chan *contracts.RegisterOracleOffline
	subscription  event.Subscription
}

func NewOfflineSubscriber(contract *contracts.Register, oracleAddress []common.Address) *OfflineSubscriber {
	return &OfflineSubscriber{
		contract: contract,
		sink:     make(chan *contracts.RegisterOracleOffline),
	}
}

func (w *OfflineSubscriber) Start(opts *bind.WatchOpts) error {
	var err error

	w.subscription, err = w.contract.WatchOracleOffline(opts, w.sink, w.oracleAddress)
	if err != nil {
		return err
	}

	go w.handleEvents()

	return nil
}

func (w *OfflineSubscriber) Stop() {
	if w.subscription != nil {
		w.subscription.Unsubscribe()
		close(w.sink)
		log.Println("OracleWatcher stopped.")
	}
}

func (w *OfflineSubscriber) handleEvents() {
	for {
		select {
		case event := <-w.sink:
			w.processEvent(event)
		case err := <-w.subscription.Err():
			log.Printf("Error in subscription: %v", err)
			return
		}
	}
}

func (w *OfflineSubscriber) processEvent(event *contracts.RegisterOracleOffline) {
	di.RegisterService().RemoveOracle(event.OracleAddress)
	log.Printf("OracleOnline Event - Address: %s, Timestamp: %s",
		event.OracleAddress,
		time.Now(),
	)
}
