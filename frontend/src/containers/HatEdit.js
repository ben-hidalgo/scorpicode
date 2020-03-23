import React, { Component } from 'react'
import { extendObservable } from 'mobx'
import { observer }  from 'mobx-react'


class HatEdit extends Component {

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
      <div className="container is-white">
        {
          hatStore.error 
          && 
          <article class="message is-warning">
            <div class="message-body">
              {languageStore.decode(hatStore.error.msg)}
            </div>
          </article>
        }
        <br/>
        <HatColors hec={this}/>
        <br/>
        <HatStyles hec={this}/>
        <br/>
        <HatSizes hec={this}/>
        <br/>
        <div class="field is-grouped">
          <div class="control">
            <button onClick={() => {this.save(this, hatStore)}} className="button is-link">Save</button>
          </div>
          <div class="control">
            <button onClick={() => {this.cancel(this, hatStore)}} className="button is-link is-light">Cancel</button>
          </div>
        </div>
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
    <div class="field">
      <label class="label">Style</label>
      <div class="control">
        <div class="select">
          <select onChange={(ce) => {props.hec.style = ce.target.value}} value={props.hec.style}>
            {
              props.styles.map(style => {
                return (
                  <option key={style.value} value={style.value} >{style.text}</option>
                )
              })              
            }
          </select>
        </div>
    </div>
  </div>
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
    <div class="field">
      <label class="label">Color</label>
      <div class="control">
        <div class="select">
          <select onChange={(ce) => {props.hec.color = ce.target.value}} value={props.hec.color}>
            {
              props.colors.map(color => {
                return (
                  <option key={color.value} value={color.value} >{color.text}</option>
                )
              })              
            }
          </select>
        </div>
    </div>
  </div>
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
    <div class="field">
      <label class="label">Size</label>
      <div class="control">
        <div class="select">
          <select onChange={(ce) => {props.hec.size = ce.target.value}} value={props.hec.size}>
            {
              props.sizes.map(size => {
                return (
                  <option key={size.value} value={size.value} >{size.text}</option>
                )
              })              
            }
          </select>
        </div>
    </div>
  </div>
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


export default observer(HatEdit)
