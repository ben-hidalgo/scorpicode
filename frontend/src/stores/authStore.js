import { extendObservable } from 'mobx'
import jwt from 'jsonwebtoken'

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

  getCookie(name) {
    var v = document.cookie.match('(^|;) ?' + name + '=([^;]*)(;|$)');
    return v ? v[2] : null;
}

  getPayload = () => {
    if (this.decoded == null) {
      var token = this.getCookie('id_token')
      this.decoded = jwt.decode(token, {complete: true});
    }

    return this.decoded.payload
  }

}

export default new AuthStore()
