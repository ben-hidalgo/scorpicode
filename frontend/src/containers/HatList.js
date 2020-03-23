import React, { Component } from 'react'
import { observer }  from 'mobx-react'
// import Hat from '../components/Hat'


class HatList extends Component {

  render() {

    const {
      hats,
      listHats,
    } = this.props.stores.hatStore

    if (!hats.length) {
      return (
      <div className="container is-white">
        <br/>
        <button onClick={listHats} className="button is-primary">List Hats</button>
      </div>)
    }
    
    return (
    <div className="container is-white">
      <br/>
      <button onClick={listHats} className="button is-primary">List Hats</button>
      <table class="table">
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
          hats.map(h => {
            return (
              <tr>
                <td>{h.id}</td>
                <td>{h.inches}</td>
                <td>{h.color}</td>
                <td>{h.style}</td>
                <td>{h.version}</td>
              </tr>
            )
          })
        }

        </tbody>
      </table>
    </div>
    )
  }

}
export default observer(HatList)
