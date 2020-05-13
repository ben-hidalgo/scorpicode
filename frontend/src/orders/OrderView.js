import React from 'react'
import Websocket from 'react-websocket';
import { useParams } from 'react-router-dom'
import { observer }  from 'mobx-react'

import StoreContext from '../storeContext'

const OrderView = () => {

  let {
    orderStore,
  } = React.useContext(StoreContext)

  let { id } = useParams()

  let order = orderStore.fetchOrder(id)

  if (!order) {
    return <NotFound id={id} />
  }

  let wsUrl = `ws://${process.env.REACT_APP_SOCKET_HOST}/ws?target=order:${order.id}`

  let handleData = (data) => {
    let result = JSON.parse(data);
    console.log(result)
  }  

  return (
    <div className="container">
      <Websocket url={wsUrl} onMessage={handleData}/>
      <div className="card">
        <header className="card-header">
          <p className="card-header-title">
            Order: {order.id}
          </p>
        </header>
        <div className="card-content">
          <div className="content">
            <div className="columns">

              <div className="column">
              <div className="field">
                  <label className="label">Size:</label>
                  {order.size}
                </div>
                <div className="field">
                  <label className="label">Quantity:</label>
                  {order.quantity}
                </div>
                <div className="field">
                  <label className="label">Notes:</label>
                  {order.notes}
                </div>
              </div>

              <div className="column">
                <div className="field">
                  <label className="label">Color:</label>
                  {order.color}
                </div>
                <div className="field">
                  <label className="label">Style:</label>
                  {order.style}
                </div>
                <div className="field">
                  <label className="label">Version:</label>
                  {order.version}
                </div>
              </div>

            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

const NotFound = (props) => {
  return (
    <div className="container">
      <div className="card">
        <header className="card-header is-danger">
          <p className="card-header-title">
            Order: {props.id} not found
          </p>
        </header>
      </div>
    </div>
  )
} // NotFound

export default observer(OrderView)
