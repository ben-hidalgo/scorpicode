import React, { useEffect } from 'react'
import { Link } from 'react-router-dom'
import { observer }  from 'mobx-react'

import StoreContext from '../storeContext';

const HatList = () => {

  const {
    hatStore,
  } = React.useContext(StoreContext);

  // Similar to componentDidMount and componentDidUpdate
  useEffect(() => {
    if (hatStore.list.length === 0) {
      hatStore.listHats()
    }
  });  
  
  return (
    <div className="container is-white">
      <br/>
      <button onClick={hatStore.listHats} className="button is-primary">Refresh</button>
      <table className="table">
        <thead>
          <tr>
            <th>ID</th>
            <th>Inches</th>
            <th>Color</th>
            <th>Style</th>
            <th>Version</th>
            <th></th>
          </tr>
        </thead>
        <tbody>
        {
          hatStore.list.map(h => {
            return (
              <tr key={h.id}>
                <td><Link to={`/hatsview/${h.id}`}>{h.id}</Link></td>
                <td>{h.inches}</td>
                <td>{h.color}</td>
                <td>{h.style}</td>
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
