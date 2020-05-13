import { extendObservable } from 'mobx'
import agent from '../agent'

class OrderStore {

  // {id: '', size: '', color: '', style: '', quantity: 0, notes: '', version: 0}

  constructor() {
    extendObservable(this, {
      isLoading: false,
      current: null,
      error: null,
    })
  } // constructor

  handleCatch(err) {

    this.error = {code: err.response.statusCode, msg: '', status: err.response.status}
    if (err.response.body) {
      this.error.msg = err.response.body.msg
    }
  }

  fetchOrder = (id) => {
    this.isLoading = true

    if (this.current && this.current.id === id) {
      return this.current
    }

    agent.Orders.fetchOrder(id)
      .then(({ order }) => {
        this.current = order
      })
      .catch(err => {
        this.handleCatch(err)
        this.current = null
      })
      .finally(() => { this.isLoading = false })

    return this.current    
  } // fetchHat




} // HatStore

export default new OrderStore()
