
class ConfigStore {

  getSocketHost = () => {
    console.log(typeof process)
    console.log(typeof process.env)
    console.log(typeof process.env.REACT_APP_SOCKET_HOST)
    if (typeof process.env.REACT_APP_SOCKET_HOST != 'undefined') {
      console.log('returning process.env.REACT_APP_SOCKET_HOST')
      return process.env.REACT_APP_SOCKET_HOST
    }
    console.log('returning window')
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
