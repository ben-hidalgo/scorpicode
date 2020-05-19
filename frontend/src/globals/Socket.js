import React from 'react'
import Websocket from 'react-websocket';

import StoreContext from '../storeContext'

const Socket = () => {

  let {
    hatStore,
    configStore,
  } = React.useContext(StoreContext)

  return (
    <div>
      <Websocket 
        url={configStore.getSocketUrl()}
        onMessage={hatStore.appendHat}
        debug={configStore.getSocketDebugBool()}
      />
    </div>
  )
}

export default Socket
