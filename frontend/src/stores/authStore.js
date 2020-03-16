import { extendObservable } from 'mobx'
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

  setCookie(name, value, days) {
    var d = new Date();
    d.setTime(d.getTime() + 24*60*60*1000*days);
    // not using string interpolation because we're inside a Go string literal
    document.cookie = name + "=" + value + ";path=/;expires=" + d.toGMTString();
  }
  getCookie(name) {
    var v = document.cookie.match('(^|;) ?' + name + '=([^;]*)(;|$)')
    return v ? v[2] : null
  }
  deleteCookie(name) { 
    this.setCookie(name, '', -1)
  }


  getPayload = () => {

    if (this.decoded) {
      return this.decoded.payload  
    }

    var token = this.getCookie(TokenCookieName)
    if (token) {
      this.decoded = jwt.decode(token, {complete: true});
      return this.decoded.payload
    }

    return {}
  }

  logout = () => {
    this.deleteCookie(TokenCookieName)
    this.decoded = null
    // TODO: navigate by injected variable
    window.location.href = 'http://localhost:8080'
  }

}

export default new AuthStore()
