import { extendObservable } from 'mobx'

class CommonStore {

  constructor() {
    extendObservable(this, {
    })
  }

}

export default new CommonStore()
