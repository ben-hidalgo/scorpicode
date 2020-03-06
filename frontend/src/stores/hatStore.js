import { extendObservable } from 'mobx'
import agent from '../agent'

class HatStore {

  constructor() {
    extendObservable(this, {
      counter: 0,
      isLoading: false,
      hats: [],
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
  makeHat = (color, size, units, style) => {
    
    this.isLoading = false

    let inches = 0

    switch(units) {
      case 'CM':
        inches = size * 2.54
        break;
      case 'INCHES':
        inches = size
        break;
      default:
        inches = size
    }
    agent.Hats.makeHat(inches, color, style)
    .then(({ hat }) => {
      
      var tempHats = []
      // add to the front of the list
      tempHats.push(hat)
      
      // for some reason added an element doesn't trigger a render
      this.hats.forEach(h => tempHats.push(h))
      this.hats = tempHats
    })
    .finally(() => { this.isLoading = false })

  } // makeHat

} // HatStore

export default new HatStore()
