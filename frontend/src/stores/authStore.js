import { extendObservable } from 'mobx'
import { getCookie, delCookie } from '../util'
import jwt from 'jsonwebtoken'

const TokenCookieName = 'id_token'

class AuthStore {

  constructor() {
    extendObservable(this, {
      token: null,
      decoded: null,
    })
    this.token = getCookie(TokenCookieName)
    if (this.token) {
      this.decoded = jwt.decode(this.token, {complete: true});
    }
  } // constructor

  hasRole = (name) => {
    // TODO: implement
    return true
  }

  getPayload = () => {
    return this.decoded.payload
  }

  getToken = () => {
    return this.token
  }

  logout = () => {
    delCookie(TokenCookieName)
    this.decoded = null
    // TODO: navigate by injected variable
    window.location.href = 'http://localhost:8080'
  }

}

export default new AuthStore()
