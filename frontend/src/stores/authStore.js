import { extendObservable } from 'mobx'
import { getCookie, deleteCookie } from '../util'
import jwt from 'jsonwebtoken'

const TokenCookieName = 'id_token'

class AuthStore {

  constructor() {
    extendObservable(this, {
      decoded: null,
    })
  } // constructor

  hasRole = (name) => {
    // TODO: implement
    return true
  }

  getPayload = () => {

    if (this.decoded) {
      return this.decoded.payload  
    }

    var token = getCookie(TokenCookieName)
    if (token) {
      this.decoded = jwt.decode(token, {complete: true});
      return this.decoded.payload
    }

    return {}
  }

  logout = () => {
    deleteCookie(TokenCookieName)
    this.decoded = null
    // TODO: navigate by injected variable
    window.location.href = 'http://localhost:8080'
  }

}

export default new AuthStore()
