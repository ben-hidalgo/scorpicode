import React, { Component } from 'react'
import { withRouter } from 'react-router-dom'
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
      <div className="box">
        <HatError hatStore={hatStore} languageStore={languageStore} />
        <HatColors hec={this}/>
        <HatStyles hec={this}/>
        <HatSizes hec={this}/>
        <div className="field is-grouped">
          <div className="control">
            <button onClick={() => {this.save(this, hatStore)}} className="button is-link">Save</button>
          </div>
          <div className="control">
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

    hec.props.history.push('/hats')
  }


}

const HatError = observer((props) => {

  return (
    <div>
      {
        props.hatStore.error
        && 
        <article className="message is-warning">
          <div className="message-body">
            {props.languageStore.decode(props.hatStore.error.msg)}
          </div>
        </article>
      }
    </div>
  )
})

const HatStyles = observer((props) => {

  return (
    <div className="field">
      <label className="label">Style</label>
      <div className="control">
        {
          props.styles.map(style => {
            return (
              <label className="radio" key={style.value}>
                <input type="radio" name="style" value={style.value} onChange={(ce) => {props.hec.style = ce.target.value}}/>
                {style.text}
            </label>
            )
          })              
        }
    </div>
  </div>
  )
})
HatStyles.defaultProps = {
  styles: [
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
    <div className="field">
      <label className="label">Color</label>
      <div className="control">
        <div className="select">
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
    <div className="field">
      <label className="label">Size</label>
      <div className="control">
        <div className="select">
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


export default withRouter(observer(HatEdit))
