import React, { Component } from "react";
import "./Game.css";
import Board from "../Board/Board";
import CurrentPlayer from "./CurrentPlayer";

export default class Game extends Component {
  state = {
    player: 1,
    board: Array.from(Array(19)).map(() => Array.from(Array(19)).fill(0))
  };

  setPlayer = player => this.setState({ player });

  occupyCell = ({ x, y }) => {
    const { player, board } = this.state;
    const cell = board[y][x];
    board[y][x] = cell === player ? 0 : player;
    this.setState({
      board
    });
  };

  sendToServer = () => {
    fetch("http://localhost:4444", {
      method: "POST",
      body: JSON.stringify(this.state.board)
    })
  }


  render() {
    const { board, player } = this.state;
    return (
      <div className="game">
        <Board onClick={this.occupyCell}>{board}</Board>
        <CurrentPlayer player={player} onClick={this.setPlayer} />
        <button onClick={this.sendToServer}>Send</button>
      </div>
    );
  }
}
