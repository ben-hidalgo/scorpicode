import { extendObservable } from 'mobx'
import agent from '../agent'

class HatStore {

  constructor() {
    extendObservable(this, {
      counter: 0,
      isLoading: false,
      hats: [
        {id: "1", inches: 7, color: "RED", style: "DERBY", version: 0},
        {id: "2", inches: 7, color: "RED", style: "DERBY", version: 0}
      ],
    })
  } // constructor

  listHats = () => {
    this.isLoading = true
    agent.Hats.listHats()
      .then(({ hats }) => {
        this.hats = hats
      })
      .finally(() => { this.isLoading = false })
  } // listHats

} // HatStore

export default new HatStore()
