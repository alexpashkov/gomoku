import React from "react";
import Board from "../Board/Board";
import CurrentPlayer from "../CurrentPlayerDisplay";
import { GameType, IBoard, ICoords, IPlayer } from "../../types";
import assocPath from "lodash/fp/assocPath";
import GameStyles from "./Game.module.css";
import * as API from "../../API";

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

  occupyCell = (coords: ICoords) => {
    const { type } = this.props;
    const { player, board } = this.state;
    const cell = board[coords.y][coords.x];

    switch (type) {
      case GameType.vsFriend:
        return API.validateMove(board, player, coords).then(valid => {
          valid
            ? this.setState({
                board: assocPath(
                  [coords.y, coords.x],
                  cell === player ? 0 : player,
                  board
                ),
                player: getNextPlayer(player)
              })
            : console.warn("invalid move");
        });
      case GameType.vsComputer:
        return console.warn("not implemented");
      case GameType.debug:
        return this.setState({
          board: assocPath(
            [coords.y, coords.x],
            cell === player ? 0 : player,
            board
          )
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
