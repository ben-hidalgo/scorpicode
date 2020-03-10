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
      case 'hat.color.required':
        return 'Color is required'
      case 'hat.inches.required':
        return 'Size is required'
      default:
        return 'An unknown error has occurred'
    }
  }

  enUK = (msg) => {

    switch(msg) {
      case 'hat.style.required':
        return 'Style is required (UK)'
      case 'hat.color.required':
        return 'Colour is required'
      case 'hat.inches.required':
        return 'Size is required (UK)'
      default:
        return 'An unexpected error has occurred'
    }
  }

}

export default new LanguageStore()
