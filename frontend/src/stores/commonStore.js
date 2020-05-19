import { extendObservable } from 'mobx'

class CommonStore {

  constructor() {
    extendObservable(this, {
    })
  }


} // CommonStore

export default new CommonStore()
