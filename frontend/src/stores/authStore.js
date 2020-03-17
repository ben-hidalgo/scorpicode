import { extendObservable } from 'mobx'
import { getCookie, deleteCookie } from '../util'
import jwt from 'jsonwebtoken'

const TokenCookieName = 'id_token'

class AuthStore {

  constructor() {
    extendObservable(this, {
      token: null,
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

    this.token = getCookie(TokenCookieName)
    if (this.token) {
      this.decoded = jwt.decode(this.token, {complete: true});
      return this.decoded.payload
    }

    // TODO: this should never happen
    return {}
  }

  getToken = () => {
    return this.token
  }

  logout = () => {
    deleteCookie(TokenCookieName)
    this.decoded = null
    // TODO: navigate by injected variable
    window.location.href = 'http://localhost:8080'
  }

}

export default new AuthStore()
