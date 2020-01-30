import Banner from './Banner';
import MainView from './MainView';
import React from 'react';
// import Tags from './Tags';
import { inject, observer } from 'mobx-react';
import { withRouter } from 'react-router-dom';

@inject('commonStore')
@withRouter
@observer
export default class Home extends React.Component {
  componentDidMount() {
    // this.props.commonStore.loadTags();
  }

  render() {
    const { token, appName } = this.props.commonStore;
    return (
      <div className="home-page">

        <Banner token={token} appName={appName} />

        <div className="container page">
          <div className="row">
            <MainView />

          </div>
        </div>

      </div>
    );
  }
}
