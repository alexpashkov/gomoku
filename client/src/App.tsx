import React, { Component } from "react";
import {
  BrowserRouter as Router,
  Link,
  Route,
  RouteComponentProps
} from "react-router-dom";
import GameCreator from "./components/GameCreator/GameCreator";
import Game from "./components/Game/Game";
import { IGameProps } from "./components/Game/Game";

class App extends Component {
  renderGame = ({
    match: {
      params: { id, type }
    }
  }: RouteComponentProps<IGameProps>) => <Game id={id} type={type} />;

  render() {
    return (
      <Router>
        <React.Fragment>
          <h1>
            <Link to="/">Gomoku</Link>
          </h1>
          <Route path="/" exact={true} component={GameCreator} />
          <Route path="/game/:id/:type" exact={true} render={this.renderGame} />
        </React.Fragment>
      </Router>
    );
  }
}

export default App;
