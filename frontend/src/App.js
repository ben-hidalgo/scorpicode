import React from "react";
import { Switch, Route, withRouter } from "react-router-dom";
import { observer } from "mobx-react";

import HatList from "components/HatList";


@withRouter
@observer
export default class App extends React.Component {

  render() {
      return (
        <div>
          <Switch>
            <Route path="/" component={HatList} />
          </Switch>
        </div>
      );
  }
}
