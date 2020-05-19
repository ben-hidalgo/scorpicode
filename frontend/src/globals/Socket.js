import React from 'react'
import Websocket from 'react-websocket';

import StoreContext from '../storeContext'

const Socket = () => {

  let {
    hatStore,
    configStore,
  } = React.useContext(StoreContext)

  let handleData = (data) => {
    let result = JSON.parse(data);
    console.log(result)
  }
  console.log(hatStore)

  return (
    <div>
      <Websocket 
      url={configStore.getSocketUrl()}
      onMessage={handleData}
      debug={configStore.getSocketDebug()}
      />
    </div>
  )
}

export default Socket
