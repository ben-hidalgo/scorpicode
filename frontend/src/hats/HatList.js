import React, { useEffect } from 'react'
import { Link } from 'react-router-dom'
import { observer }  from 'mobx-react'

import StoreContext from '../storeContext'

const HatList = () => {

  const {
    hatStore,
  } = React.useContext(StoreContext)

  // similar to componentDidMount and componentDidUpdate
  useEffect(() => {
    if (!hatStore.listInit) {
      hatStore.listHats()
    }
  })
  
  return (
    <div className="container is-white">
      <br/>
      <button onClick={hatStore.listHats} className="button is-primary">Refresh</button>
      <table className="table">
        <thead>
          <tr>
            <th>ID</th>
            <th>Size</th>
            <th>Color</th>
            <th>Style</th>
            <th>Quantity</th>
            <th>Notes</th>
            <th>Version</th>
            <th></th>
          </tr>
        </thead>
        <tbody>
        { hatStore.list.map(h => {
            return (
              <tr key={h.id}>
                <td><Link to={`/hatsview/${h.id}`}>{h.id}</Link></td>
                <td>{h.size}</td>
                <td>{h.color}</td>
                <td>{h.style}</td>
                <td>{h.quantity}</td>
                <td>{h.notes}</td>
                <td>{h.version}</td>
                <td><button onClick={() => {hatStore.deleteHat(h.id, h.version)}} className="delete"/></td>
              </tr>
            )
          })
        }
        </tbody>
      </table>
    </div>
  )
}

export default observer(HatList)
