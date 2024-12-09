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

type OnlineSubscriber struct {
	contract      *contracts.Register
	oracleAddress []common.Address
	sink          chan *contracts.RegisterOracleOnline
	subscription  event.Subscription
}

func NewOnlineSubscriber(contract *contracts.Register, oracleAddress []common.Address) *OnlineSubscriber {
	return &OnlineSubscriber{
		contract: contract,
		sink:     make(chan *contracts.RegisterOracleOnline),
	}
}

func (w *OnlineSubscriber) Start(opts *bind.WatchOpts) error {
	var err error

	w.subscription, err = w.contract.WatchOracleOnline(opts, w.sink, w.oracleAddress)
	if err != nil {
		return err
	}

	go w.handleEvents()

	return nil
}

func (w *OnlineSubscriber) Stop() {
	if w.subscription != nil {
		w.subscription.Unsubscribe()
		close(w.sink)
		log.Println("OracleWatcher stopped.")
	}
}

func (w *OnlineSubscriber) handleEvents() {
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

func (w *OnlineSubscriber) processEvent(event *contracts.RegisterOracleOnline) {
	di.RegisterService().AddOracle(event.OracleAddress)
	log.Printf("OracleOnline Event - Address: %s, Timestamp: %s",
		event.OracleAddress,
		time.Now(),
	)
}
