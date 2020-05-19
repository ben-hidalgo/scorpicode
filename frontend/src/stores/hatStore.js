import { extendObservable } from 'mobx'
import agent from '../agent'

class HatStore {

  // {id: '', size: '', color: '', style: '', quantity: 0, notes: '', version: 0}

  constructor() {
    extendObservable(this, {
      isLoading: false,
      current: null,
      draft: null,
      list: [],
      listInit: false,
      error: null,
    })
    this.initDraft()
  } // constructor

  handleCatch(err) {

    this.error = {code: err.response.statusCode, msg: '', status: err.response.status}
    if (err.response.body) {
      this.error.msg = err.response.body.msg
    }
  }

  initDraft = () => {
    this.draft = {
      color: '',
      size: '',
      style: 'UNKNOWN_STYLE', // protobuf enum
      notes: '',
      quantity: 1,
    }
    this.error = null
  }

  cancelMakeHat = (history) => {
    this.initDraft()
    history.goBack()
  }

  fetchHat = (id) => {
    this.isLoading = true

    if (this.current && this.current.id === id) {
      return this.current
    }

    agent.Hats.fetchHat(id)
      .then(({ hat }) => {
        this.current = hat
      })
      .catch(err => {
        this.handleCatch(err)
        this.current = null
      })
      .finally(() => { this.isLoading = false })

    return this.current    
  } // fetchHat


  // TODO: rename to fetchHatList()
  // retrieves hats from the server
  listHats = () => {
    this.isLoading = true
    agent.Hats.listHats()
      .then(({ hats }) => {
        // Twirp gRPC doesn't return the empty array
        if (hats) {
          this.list = hats
        }
        
        this.listInit = true
      })
      .catch(err => {
        this.handleCatch(err)
      })
      .finally(() => { 
        this.isLoading = false 
      })
    
  } // listHats

  // creates a new hat
  makeHat = (history) => {
    
    this.isLoading = true

    // size is inches
    agent.Hats.makeHat(this.draft)
      // TODO: this needs to handle order response...
      .then(({ order }) => {        
        this.error = null
        this.current = order
        history.push(`/ordersview/${order.id}`)
      })
      .catch(err => {
        this.handleCatch(err)
      })
      .finally(() => { 
        this.isLoading = false
      })
    
  } // makeHat

  // deletes an existing hat
  deleteHat = (id, version) => {
    this.isLoading = true
    agent.Hats.deleteHat(id, version)
      .then(() => {
        this.list = this.list.filter((v, i, a) => {return v.id !== id})
      })
      .catch(err => {
        console.log(err)
        // this.handleCatch(err)
      })
      .finally(() => { this.isLoading = false })

  } // deleteHat

  // allows web socket to append hats
  appendHat = (hat) => {
    this.list.push(hat)
  } // appendHat


} // HatStore

export default new HatStore()
