import React from 'react'
import Websocket from 'react-websocket';

const handleData = (data) => {
  let result = JSON.parse(data);
  console.log(result)
}

const Socket = () => {

  return (
    <div>
      <Websocket url={`ws://${process.env.REACT_APP_SOCKET_HOST}/ws`} onMessage={handleData}/>
    </div>
  )
}

export default Socket
