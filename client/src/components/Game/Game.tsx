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
  aiPlayer?: IPlayer;
}

interface IGameState {
  player: IPlayer;
  board: IBoard;
  blackScore: number;
  whiteScore: number;
  suggestedMoves: ICoords[];
  aiIsThinking: boolean;
  aiResponseTime: number;
}

function getNextPlayer(p: IPlayer) {
  return p == IPlayer.Black ? IPlayer.White : IPlayer.Black;
}

export default class Game extends React.Component<IGameProps, IGameState> {
  componentDidMount(): void {
    const { type, aiPlayer } = this.props;
    const { player } = this.state;
    if (type === GameType.vsComputer && player === aiPlayer) this.aiMove();
  }

  state: IGameState = {
    player: IPlayer.Black,
    board: Board.init(),
    blackScore: 0,
    whiteScore: 0,
    suggestedMoves: [],
    aiIsThinking: false,
    aiResponseTime: 0
  };

  setPlayer = (player: IPlayer) => {
    this.props.type === GameType.debug && this.setState({ player });
  };

  humanMove = (coords: ICoords) => {
    const { player, board } = this.state;
    API.validateMove(board, player, coords).then(valid => {
      valid && this.placeStone(coords);
    });
  };

  aiMove = () => {
    this.setState({
      aiIsThinking: true
    });
    const now = Date.now();
    API.suggestMove(this.state).then(coords => {
      const aiResponseTime = Date.now() - now;
      this.placeStone(coords);
      this.setState({
        aiIsThinking: false,
        aiResponseTime
      });
    });
  };

  toggleMove = (coords: ICoords) => {
    const { player, board } = this.state;
    const cell = board[coords.y][coords.x];

    return this.setState({
      board: assocPath(
        [coords.y, coords.x],
        cell === player ? 0 : player,
        board
      )
    });
  };

  placeStone = (coords: ICoords) => {
    const { aiPlayer } = this.props;
    const { player, board } = this.state;
    const nextPlayer = getNextPlayer(player);
    this.setState(
      {
        board: assocPath([coords.y, coords.x], player, board),
        player: nextPlayer
      },
      () => nextPlayer == aiPlayer && this.aiMove()
    );
  };

  handleCellClick = (coords: ICoords) => {
    const { type, aiPlayer } = this.props;
    const { player, aiIsThinking } = this.state;

    if (aiIsThinking) {
      return;
    }
    if (
      type === GameType.vsFriend ||
      (type === GameType.vsComputer && player !== aiPlayer)
    ) {
      this.humanMove(coords);
    } else if (type === GameType.debug) {
      this.toggleMove(coords);
    }
  };

  sendBoardToServer = () => API.sendBoard(this.state.board);

  render() {
    const { type } = this.props;
    const { board, player, aiResponseTime } = this.state;
    return (
      <div className={GameStyles.container}>
        <div className={GameStyles.infoPanel}>
          <CurrentPlayer
            player={player}
            onClick={this.setPlayer}
            allowPlayerSelection={type == GameType.debug}
          />
          {type === GameType.vsComputer && !!aiResponseTime && (
            <div className={GameStyles.responseTime}>
              AI Response Time: {aiResponseTime}ms
            </div>
          )}
          {type === GameType.debug && (
            <button onClick={this.sendBoardToServer}>Send To Server</button>
          )}
        </div>
        <Board onClick={this.handleCellClick}>{board}</Board>
      </div>
    );
  }
}
