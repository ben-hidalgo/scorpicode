import React, { Component } from 'react'
import { extendObservable, action } from 'mobx'
import './HatEditContainer.css'
import { observer }  from 'mobx-react'


class HatEditContainer extends Component {

  constructor() {
    super()
    extendObservable(this, {
      color: 'RED',
    })
  }

  render() {

    const colors = [
      {value: 'RED', text: 'Red'},
      {value: 'BLUE', text: 'Blue'},
      {value: 'GREEN', text: 'Green'},
      {value: 'YELLOW', text: 'Yellow'},
      {value: 'PURPLE', text: 'Purple'},
      {value: 'BLACK', text: 'Black'},
      {value: 'GREY', text: 'Grey'},
      {value: 'ORANGE', text: 'Orange'},
    ]
    
    return (
      <div className="HatEditContainer">
        <label>
          <select onChange={this.handleColorChange}>
            {
              colors.map(color => {
                return (
                  <option key={color.value} value={color.value} >{color.text}</option>
                )
              })              
            }
          </select>
          Color
        </label>
        
        <button onClick={() => {console.log('save')}} type="button">Save</button>
        <button onClick={() => {console.log('cancel')}} type="button">Cancel</button>
      </div>
    )
  }

  handleColorChange = (changeEvent) => {
    this.color = changeEvent.target.value
  }

}
export default observer(HatEditContainer)
