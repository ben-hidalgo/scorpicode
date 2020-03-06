import React, { Component } from 'react'
import { extendObservable } from 'mobx'
import './HatEditContainer.css'
import { observer }  from 'mobx-react'


class HatEditContainer extends Component {

  constructor() {
    super()
    extendObservable(this, {
      color: 'RED',
      size: '06_00',
    })
  }

  render() {

    return (
      <div className="HatEditContainer">
        <span>_color_ {this.color}, _size_ {this.size}</span>
        <br/>
        <HatColors hec={this}/>
        <HatSizes hec={this}/>
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

function HatSizes(props) {

  return (
      <label>
        <select onChange={(ce) => {props.hec.size = ce.target.value}}>
          {
            props.sizes.map(size => {
              return (
                <option key={size.value} value={size.value} >{size.text}</option>
              )
            })              
          }
        </select>
        Size
      </label>
  )
}
HatSizes.defaultProps = {
  sizes: [
    {value: '06_00', text: '6 inches'},
    {value: '06_25', text: '6 1/4 inches'},
    {value: '06_50', text: '6 1/2 inches'},
  ],

};

export default observer(HatEditContainer)
