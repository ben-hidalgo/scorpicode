import React, { Component } from 'react'
import { withRouter } from 'react-router-dom'
import { extendObservable } from 'mobx'
import { observer }  from 'mobx-react'


class HatCreate extends Component {

  constructor() {
    super()
    extendObservable(this, {
      color: '',
      size: '',
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
    hatStore.makeHat(hec.size, hec.color, hec.style, hec.props.history)

    // TODO: this is why a server side error resets the form...
    // How to split the "store" logic from the form updating logic...?
    hec.color = ''
    hec.size = ''
    hec.style = 'UNKNOWN_STYLE'
  }

  cancel(hec, hatStore) {
    hec.color = ''
    hec.size = ''
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
        {
          props.colors.map(color => {
            return (
              <label className="radio" key={color.value}>
                <span onClick={(ce) => {props.hec.color = color.value}} className={`bd-color-${color.value === props.hec.color} has-background-${color.c}`} ></span>
              </label>
            )
          })              
        }
    </div>
  </div>
  )
})
HatColors.defaultProps = {
  colors: [
    {value: 'RED', text: 'Red', c: 'red'},
    {value: 'BLUE', text: 'Blue', c: 'blue'},
    {value: 'GREEN', text: 'Green', c: 'green'},
    {value: 'YELLOW', text: 'Yellow', c: 'yellow'},
    {value: 'PURPLE', text: 'Purple', c: 'purple'},
    {value: 'BLACK', text: 'Black', c: 'black'},
    {value: 'GREY', text: 'Grey', c: 'grey'},
    {value: 'ORANGE', text: 'Orange', c: 'orange'},
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
    {value: "", text: 'Please select a size'},
    {value: "06000", text: '6 inches'},
    {value: "06125", text: '6 1/8 inches'},
    {value: "06250", text: '6 1/4 inches'},
    {value: "06375", text: '6 3/8 inches'},
    {value: "06500", text: '6 1/2 inches'},
    {value: "06625", text: '6 5/8 inches'},
    {value: "06750", text: '6 3/4 inches'},
    {value: "06875", text: '6 7/8 inches'},
    {value: "07000", text: '7 inches'},
  ],
}


export default withRouter(observer(HatCreate))
