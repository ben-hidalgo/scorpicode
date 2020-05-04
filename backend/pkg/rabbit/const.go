package rabbit

// Exchanges
const (
	// ServiceMsgtypeTx topic exchange
	ServiceMsgtypeTx = Exchange("service_msgtype_tx")
)

// Keys
const (
	HatsDotStar = RKey("hats.*")
)

// Queues
const (
	HatsQueue = Queue("hats_q")
)

const schema = `
exchanges:
  service_msgtype_tx:
    durable: true
    type: topic
queues:
  hats_q:
    durable: true
    bindings:
      - exchange: "service_msgtype_tx"
        key: "hats.*"
`
