import React, { Component } from "react";
import {
  BrowserRouter as Router,
  Link,
  Route,
  RouteComponentProps
} from "react-router-dom";
import queryString from "querystring";
import GameCreator from "./components/GameCreator/GameCreator";
import Game from "./components/Game/Game";
import { GameType } from "./types";

class App extends Component {
  renderGame = ({
    match: {
      params: { id, type }
    },
    location: { search }
  }: RouteComponentProps<{
    id: string;
    type: GameType;
  }>) => {
    const { aiPlayer } = queryString.parse(search.replace("?", ""));
    return <Game id={id} type={type} aiPlayer={+aiPlayer} />;
  };

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
