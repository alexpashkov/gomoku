import React, { Component } from "react";
// import { BrowserRouter as Router, Route, Link } from "react-router-dom";
// import GameCreator from "./components/GameCreator";
import Game from "./components/Game";

class App extends Component {
  render() {
    return (
      <Game />
      // <Router>
      //   <Route path="/new-game" component={GameCreator} />
      // </Router>
    );
  }
}

export default App;
