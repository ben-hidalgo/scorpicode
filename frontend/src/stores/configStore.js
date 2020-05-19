
class ConfigStore {

  getSocketHost = () => {
    if (typeof process.env.REACT_APP_SOCKET_HOST != 'undefined') {
      return process.env.REACT_APP_SOCKET_HOST
    }
    return window['REACT_APP_SOCKET_HOST']
  }

  getSocketUrl = () => {
    return `ws://${this.getSocketHost()}/ws`
  }

  getSocketDebug = () => {
    if (typeof process != 'undefined') {
      return process.env.REACT_APP_SOCKET_DEBUG
    }
    return window['REACT_APP_SOCKET_DEBUG']
  }

  getSocketDebugBool = () => {
    return this.getSocketDebug() === 'true'
  }


} // ConfigStore

export default new ConfigStore()
