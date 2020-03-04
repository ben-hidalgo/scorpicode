import { extendObservable } from 'mobx'

class CommonStore {

  constructor() {
    extendObservable(this, {
      counter: 0,
    })
  }

  onIncrement = () => {
    this.counter++
  }
  onDecrement = () => {
    this.counter--
  }


}

export default new CommonStore()
