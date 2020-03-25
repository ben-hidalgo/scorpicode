import React, { Component } from 'react'
import { observer }  from 'mobx-react'


class HatList extends Component {

  componentDidMount() {
    this.props.stores.hatStore.listHats()
  }

  render() {

    const {
      hats,
      listHats,
    } = this.props.stores.hatStore

    return (
    <div className="container is-white">
      <br/>
      <button onClick={listHats} className="button is-primary">Refresh</button>
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
          hats.map(h => {
            return (
              <tr key={h.id}>
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
