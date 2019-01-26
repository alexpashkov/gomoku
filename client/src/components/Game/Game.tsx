import React from "react";
import Board from "../Board/Board";
import CurrentPlayer from "../CurrentPlayerDisplay";
import { GameType, IBoard, ICoords, IPlayer } from "../../types";
import assocPath from "lodash/fp/assocPath";
import GameStyles from "./Game.module.css";

export interface IGameProps {
  id: string;
  type: GameType;
}

interface IGameState {
  player: IPlayer;
  board: IBoard;
}

function getNextPlayer(p: IPlayer) {
  return p == IPlayer.Black ? IPlayer.White : IPlayer.Black;
}

export default class Game extends React.Component<IGameProps, IGameState> {
  state = {
    player: IPlayer.Black,
    board: Board.init()
  };

  setPlayer = (player: IPlayer) => {
    this.setState({ player });
  };

  occupyCell = ({ x, y }: ICoords) => {
    const { type } = this.props;
    const { player, board } = this.state;
    const cell = board[y][x];

    switch (type) {
      case GameType.vsFriend:
        return (
          !cell &&
          this.setState({
            board: assocPath([y, x], cell === player ? 0 : player, board),
            player: getNextPlayer(player)
          })
        );
      case GameType.vsComputer:
        return console.warn("Not implemented");
      case GameType.debug:
        return this.setState({
          board: assocPath([y, x], cell === player ? 0 : player, board)
        });
    }
  };

  sendBoardToServer = () => {
    fetch("/board", {
      method: "POST",
      body: JSON.stringify(this.state.board)
    });
  };

  render() {
    const { type } = this.props;
    const { board, player } = this.state;
    return (
      <div className={GameStyles.container}>
        <Board onClick={this.occupyCell}>{board}</Board>
        <div className={GameStyles.sidebar}>
          <CurrentPlayer
            player={player}
            onClick={this.setPlayer}
            allowPlayerSelection={type == GameType.debug}
          />
          {type === GameType.debug && (
            <button onClick={this.sendBoardToServer}>Send To Server</button>
          )}
        </div>
      </div>
    );
  }
}
