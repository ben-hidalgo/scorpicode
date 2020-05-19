
class ConfigStore {

  getSocketHost = () => {
    if (typeof process != 'undefined') {
      return process.env.REACT_APP_SOCKET_HOST
    }
    return window['REACT_APP_SOCKET_HOST']
  }

  getSocketUrl = () => {
    return `ws://${this.getSocketHost()}/ws`
  }

  getSocketDebug = () => {
    if (typeof process != 'undefined') {
      return process.env.REACT_APP_SOCKET_DEBUG === 'true'
    }
    return window['REACT_APP_SOCKET_DEBUG'] === 'true'
  }


} // ConfigStore

export default new ConfigStore()
