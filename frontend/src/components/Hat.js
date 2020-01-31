import React from 'react';
// import { Link } from 'react-router-dom';
// import { inject, observer } from 'mobx-react';
import { observer } from 'mobx-react';



//@inject('articlesStore')
@observer
export default class Hat extends React.Component {

  render() {

    return (
      <div className="article-preview">
        <li>id: {this.props.hat.id} name: {this.props.hat.name} inches: {this.props.hat.inches} color: {this.props.hat.color}</li>
      </div>
    );
  }
}
