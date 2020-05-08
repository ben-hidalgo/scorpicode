package rabbit

// Exchanges are hyphen separated (by convention)
const (
	// ServiceMsgAction topic exchange
	ServiceMsgActionX = Exchange("service-msg-action")
)

// Keys are dot separated (by rule) and can contain asterisk or hash (* or #)
const (
	HatsOrderCreatedK = RKey("hats.order.created")
	HatsHatCreatedK   = RKey("hats.hat.created")
)

// Queues are underscore separated (by convention)
const (
	HatsOrderCreatedQ = Queue("hats_order_created")
	// this one is probably temporary... and/or will move the the web sockets service
	HatsHatCreatedQ = Queue("hats_hat_created")
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
  hats_hat_created:
    durable: true
    bindings:
      - exchange: "service-msg-action"
        key: "hats.hat.created"    
`
