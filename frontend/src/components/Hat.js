import React from 'react';
import { observer } from 'mobx-react';

@observer
export default class Hat extends React.Component {

  render() {
    return (
      <div className="article-preview">
        <li>Id: {this.props.hat.id} Style: {this.props.hat.style} Inches: {this.props.hat.inches} Color: {this.props.hat.color}</li>
      </div>
    );
  }
}
