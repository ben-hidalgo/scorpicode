import { extendObservable } from 'mobx'
import agent from '../agent'

class HatStore {

  // {size: '', color: '', style: '', quantity: 0, notes: ''}

  constructor() {
    extendObservable(this, {
      isLoading: false,
      current: null,
      draft: null,
      list: [],
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
      style: 'UNKNOWN_STYLE',
      notes: '',
      quantity: 0,
    }
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


  // returns all hats
  listHats = () => {
    this.isLoading = true
    agent.Hats.listHats()
      .then(({ hats }) => {
        this.list = hats
      })
      .catch(err => {
        this.handleCatch(err)
      })
      .finally(() => { this.isLoading = false })
    
    this.isLoading = false
  } // listHats

  // creates a new hat
  makeHat = (history) => {
    
    this.isLoading = true

    // size is inches
    agent.Hats.makeHat(this.draft)
      .then(({ hat }) => {        
        this.error = null
        this.current = hat
        history.push(`/hatsview/${hat.id}`)
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
        this.hats = this.hats.filter((v, i, a) => {return v.id !== id})
      })
      .catch(err => {
        this.handleCatch(err)
      })
      .finally(() => { this.isLoading = false })

  } // deleteHat

} // HatStore

export default new HatStore()
