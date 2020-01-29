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
        <p>Hat: {this.props.hat.id} {this.props.hat.name}</p>
      </div>
    );
  }
}
