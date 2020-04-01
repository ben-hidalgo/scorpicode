import React, { useEffect } from 'react'
import { withRouter } from 'react-router-dom'
import { observer }  from 'mobx-react'

import StoreContext from '../storeContext'



const HatCreate = (props) => {

  const {
    hatStore,
    languageStore,
  } = React.useContext(StoreContext)

  // similar to componentDidMount and componentDidUpdate
  useEffect(() => {
    if (hatStore.list.length === 0) {
      hatStore.initDraft()
    }
  })

  return (
    <div className="box">
      <HatError hatStore={hatStore} languageStore={languageStore} />
      <HatColors hatStore={hatStore} />
      <HatStyles hatStore={hatStore} />
      <HatSizes hatStore={hatStore} />
      <div className="field is-grouped">
        <div className="control">
          <button onClick={() => {hatStore.makeHat(props.history)}} className="button is-link">Save</button>
        </div>
        <div className="control">
          <button onClick={() => {hatStore.cancelMakeHat(props.history)}} className="button is-link is-light">Cancel</button>
        </div>
      </div>
    </div>
  )
}

const HatError = observer((props) => {

  const {
    hatStore,
    languageStore,
  } = React.useContext(StoreContext)

  return (
    <div>
      {
        props.hatStore.error
        && 
        <article className="message is-warning">
          <div className="message-body">
            {languageStore.decode(hatStore.error.msg)}
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
                <input type="radio" name="style" value={style.value} onChange={(ce) => {props.hatStore.draft.style = ce.target.value}}/>
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
                <span onClick={(ce) => {props.hatStore.draft.color = color.value}} className={`bd-color-${color.value === props.hatStore.draft.color} has-background-${color.c}`} ></span>
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
          <select onChange={(ce) => {props.hatStore.draft.size = ce.target.value}} value={props.hatStore.draft.size}>
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