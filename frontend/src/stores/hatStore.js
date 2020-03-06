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

  listHats = () => {
    this.isLoading = true
    agent.Hats.listHats()
      .then(({ hats }) => {
        this.hats = hats
      })
      .finally(() => { this.isLoading = false })
    this.isLoading = false  
  } // listHats

  makeHat = (color, size, units, style) => {
    
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
  } // listHats

} // HatStore

export default new HatStore()
