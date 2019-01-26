import React from "react";
import Board from "../Board";
import CurrentPlayer from "../CurrentPlayerDisplay";
import { IBoard, ICoords, IPlayer } from "../../types";
import assocPath from "lodash/fp/assocPath";
import GameStyles from "./Game.module.css";

interface IGameProps {
}

interface IGameState {
  player: IPlayer;
  board: IBoard;
}

export default class Game extends React.Component<{}, IGameState> {
  state = {
    player: IPlayer.Black,
    board: Array.from(Array(19)).map(() =>
      Array.from(Array(19)).fill(0)
    ) as IBoard
  };

  setPlayer = (player: IPlayer) => {
    this.setState({ player });
  };

  occupyCell = ({ x, y }: ICoords) => {
    const { player, board } = this.state;
    const cell = board[y][x];
    this.setState({
      board: assocPath([y, x], cell === player ? 0 : player, board)
    });
  };

  sendToServer = () => {
    fetch("/board", {
      method: "POST",
      body: JSON.stringify(this.state.board)
    });
  };

  render() {
    console.log(this.props)
    const { board, player } = this.state;
    return (
      <div className={GameStyles.container}>
        <Board onClick={this.occupyCell}>{board}</Board>
        <CurrentPlayer player={player} onClick={this.setPlayer} />
        <button onClick={this.sendToServer}>Send</button>
      </div>
    );
  }
}
