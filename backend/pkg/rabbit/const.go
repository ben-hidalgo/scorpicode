package rabbit

// Exchanges are hyphen separated (by convention)
const (
	// ServiceMsgAction topic exchange
	ServiceMsgAction = Exchange("service-msg-action")
)

// Keys are dot separated (by rule) and can contain asterisk or hash (* or #)
const (
	HatsOrderCreated = RKey("hats.order.created")
)

// Queues are underscore separated (by convention)
const (
	HatsQueue = Queue("hats_order_created")
)

const schema = `
exchanges:
  service-msg-action:
    durable: true
    type: topic
queues:
  hats_order_created:
    durable: true
    bindings:
      - exchange: "service-msg-action"
        key: "hats.order.created"
`
