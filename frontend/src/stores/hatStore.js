import { extendObservable } from 'mobx'
import agent from '../agent'

class HatStore {

  constructor() {
    extendObservable(this, {
      counter: 0,
      isLoading: false,
      hats: [],
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
        this.hats = hats
      })
      .catch(err => {
        this.handleCatch(err)
      })
      .finally(() => { this.isLoading = false })
    
    this.isLoading = false
  } // listHats

  // saves a new hat
  makeHat = (color, size, style) => {
    
    this.isLoading = true

    // size is inches
    agent.Hats.makeHat(size, color, style)
      .then(({ hat }) => {
        
        var temp = [hat]

        // for some reason adding an element isn't triggering render
        this.hats.forEach(h => temp.push(h))
        this.hats = temp
        this.error = null
      })
      .catch(err => {
        this.handleCatch(err)
      })
      .finally(() => { 
        this.isLoading = false
      })
    
  } // makeHat

  // returns all hats
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
