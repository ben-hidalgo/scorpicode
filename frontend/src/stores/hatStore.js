import { extendObservable } from 'mobx'
import agent from '../agent'

class HatStore {

  constructor() {
    extendObservable(this, {
      counter: 0,
      isLoading: false,
      hats: [],
      error: null,
    })
  } // constructor

  // returns all hats
  listHats = () => {
    this.isLoading = true
    agent.Hats.listHats()
      .then(({ hats }) => {
        this.hats = hats
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
      })
      .catch(err => {
        this.error = {code: err.response.body.code, msg: err.response.body.msg, status: err.response.status}
      })
      .finally(() => { 
        this.isLoading = false
        this.error = null
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
        this.error = {code: err.response.body.code, msg: err.response.body.msg, status: err.response.status}
      })
      .finally(() => { this.isLoading = false })

  } // deleteHat

} // HatStore

export default new HatStore()
