import React, { Component } from "react";
import { BrowserRouter as Router, Route } from "react-router-dom";
import GameCreator from "./components/GameCreator";
import Game from "./components/Game";

class App extends Component {
  render() {
    return (
      <Router>
        <React.Fragment>
          <h1>Gomoku</h1>
          <Route component={GameCreator} path="/" exact={true} />
          <Route component={Game} path="/game/:id/:type" exact={true} />
        </React.Fragment>
      </Router>
    );
  }
}

export default App;
