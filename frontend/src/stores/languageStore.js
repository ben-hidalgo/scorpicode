import { extendObservable } from 'mobx'

class LanguageStore {

  constructor() {
    extendObservable(this, {
      current: 'en.us',
    })
  }

  decode = (msg) => {
    switch(this.current) {
      case 'en.us':
        return this.enUS(msg)
      case 'en.uk':
        return this.enUK(msg)
      default:
        return this.enUS(msg)
    }
  }

  enUS = (msg) => {

    switch(msg) {
      case 'hat.style.required':
        return 'Style is required'
      default:
        return 'An unknown error has occurred'
    }

  }

  enUK = (msg) => {

    switch(msg) {
      case 'hat.style.required':
        return 'The Style is required'
      default:
        return 'An unexpected error has occurred'
    }

  }


}

export default new LanguageStore()
