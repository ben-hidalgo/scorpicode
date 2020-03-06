import React, { Component } from 'react'
import { extendObservable } from 'mobx'
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

    return (
      <div className="HatEditContainer">
        <span>_color_ {this.color}</span>
        <br/>
        <HatColors hec={this}/>
        <br/>
        <button onClick={() => {console.log('save')}} type="button">Save</button>
        <button onClick={() => {console.log('cancel')}} type="button">Cancel</button>
      </div>
    )
  }

}

function HatColors(props) {

  return (
      <label>
        <select onChange={(ce) => {props.hec.color = ce.target.value}}>
          {
            props.colors.map(color => {
              return (
                <option key={color.value} value={color.value} >{color.text}</option>
              )
            })              
          }
        </select>
        Color
      </label>
  )
}
HatColors.defaultProps = {
  colors: [
    {value: 'RED', text: 'Red'},
    {value: 'BLUE', text: 'Blue'},
    {value: 'GREEN', text: 'Green'},
    {value: 'YELLOW', text: 'Yellow'},
    {value: 'PURPLE', text: 'Purple'},
    {value: 'BLACK', text: 'Black'},
    {value: 'GREY', text: 'Grey'},
    {value: 'ORANGE', text: 'Orange'},
  ],

};


export default observer(HatEditContainer)
