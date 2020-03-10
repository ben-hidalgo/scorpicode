import React, { Component } from 'react'
import { extendObservable } from 'mobx'
import './HatEditContainer.css'
import { observer }  from 'mobx-react'


class HatEditContainer extends Component {

  constructor() {
    super()
    extendObservable(this, {
      color: '',
      size: 0,
      style: 'UNKNOWN_STYLE',
    })
  }

  render() {

    const {
      hatStore,
      languageStore,
    } = this.props.stores

    return (
      <div className="HatEditContainer">
        {hatStore.error && <span className="warning">{languageStore.decode(hatStore.error.msg)}</span>}
        <br/>
        <HatColors hec={this}/>
        <br/>
        <HatStyles hec={this}/>
        <br/>
        <HatSizes hec={this}/>
        <br/>
        <button onClick={() => {this.save(this, hatStore)}} type="button">Save</button>
        <button onClick={() => {this.cancel(this, hatStore)}} type="button">Cancel</button>
      </div>
    )
  }

  save(hec, hatStore) {
    hatStore.makeHat(hec.color, hec.size, hec.style)
    hec.color = ''
    hec.size = 0
    hec.style = 'UNKNOWN_STYLE'
  }

  cancel(hec, hatStore) {
    hec.color = ''
    hec.size = 0
    hec.style = 'UNKNOWN_STYLE'
    hatStore.error = null
  }


}

const HatStyles = observer((props) => {

  return (
      <label>
        Style
        <select onChange={(ce) => {props.hec.style = ce.target.value}} value={props.hec.style}>
          {
            props.styles.map(style => {
              return (
                <option key={style.value} value={style.value} >{style.text}</option>
              )
            })              
          }
        </select>
      </label>
  )
})
HatStyles.defaultProps = {
  styles: [
    {value: 'UNKNOWN_STYLE', text: 'Please select a style'},
    {value: 'BOWLER', text: 'Bowler'},
    {value: 'FEDORA', text: 'Fedora'},
    {value: 'BASEBALL', text: 'Baseball Cap'},
    {value: 'NEWSBOY', text: 'Newsboy'},
    {value: 'COWBOY', text: 'Cowboy Hat'},
    {value: 'DERBY', text: 'Derby'},
    {value: 'TOP_HAT', text: 'Top Hat'},
  ],
}

const HatColors = observer((props) => {

  return (
      <label>
        Color
        <select onChange={(ce) => {props.hec.color = ce.target.value}} value={props.hec.color}>
          {
            props.colors.map(color => {
              return (
                <option key={color.value} value={color.value} >{color.text}</option>
              )
            })              
          }
        </select>
      </label>
  )
})
HatColors.defaultProps = {
  colors: [
    {value: '', text: 'Please select a color'},
    {value: 'RED', text: 'Red'},
    {value: 'BLUE', text: 'Blue'},
    {value: 'GREEN', text: 'Green'},
    {value: 'YELLOW', text: 'Yellow'},
    {value: 'PURPLE', text: 'Purple'},
    {value: 'BLACK', text: 'Black'},
    {value: 'GREY', text: 'Grey'},
    {value: 'ORANGE', text: 'Orange'},
  ],
}

const HatSizes = observer((props) => {

  return (
      <label>
        Size
        <select onChange={(ce) => {props.hec.size = ce.target.value}} value={props.hec.size}>
          {
            props.sizes.map(size => {
              return (
                <option key={size.value} value={size.value} >{size.text}</option>
              )
            })              
          }
        </select>
      </label>
  )
})
HatSizes.defaultProps = {
  sizes: [
    {value: 0, text: 'Please select a size'},
    {value: 6, text: '6 inches'},
    {value: 7, text: '7 inches'},
    {value: 8, text: '8 inches'},
  ],
}


export default observer(HatEditContainer)
