import React from 'react'
import Websocket from 'react-websocket';

import StoreContext from '../storeContext'

const handleData = (data, hatStore) => {

  let msg = JSON.parse(data);
  
  if (msg.queue === 'soxie_hat_created') {
    hatStore.appendHat(msg.data)
  }
}

const Socket = () => {

  let {
    configStore,
    hatStore,
  } = React.useContext(StoreContext)

  return (
    <div>
      <Websocket 
        url={configStore.getSocketUrl()}
        onMessage={(data) => {handleData(data, hatStore)}}
        debug={configStore.getSocketDebugBool()}
      />
    </div>
  )
}

export default Socket
