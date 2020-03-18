import { extendObservable } from 'mobx'
import { getCookie, delCookie } from '../util'
import jwt from 'jsonwebtoken'

const TokenCookieName = 'id_token'

class AuthStore {

  constructor() {
    extendObservable(this, {
      token: null,
      decoded: null,
      payload: null,
      picture: null,
    })
    this.token = getCookie(TokenCookieName)
    if (this.token) {
      this.decoded = jwt.decode(this.token, {complete: true});
    }
    if (this.decoded) {
      this.payload = this.decoded.payload
    }
    if (this.payload) {
      this.picture = this.payload.picture
    }
  } // constructor


  logout = () => {
    delCookie(TokenCookieName)
    this.token = null
    this.decoded = null
    this.payload = null
    // TODO: navigate by injected variable
    window.location.href = 'http://localhost:8080'
  }

}

export default new AuthStore()
