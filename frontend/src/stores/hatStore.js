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
        
        var tempHats = []
        // add to the front of the list
        tempHats.push(hat)

        // for some reason adding an element isn't triggering render
        this.hats.forEach(h => tempHats.push(h))
        this.hats = tempHats
      })
      .catch(err => {
        this.error = {code: err.response.body.code, msg: err.response.body.msg, status: err.response.status}
      })
      .finally(() => { 
        this.isLoading = false 
      })
    
  } // makeHat

} // HatStore

export default new HatStore()
