import React from "react";
import Board from "../Board/Board";
import CurrentPlayer from "../CurrentPlayerDisplay";
import { GameType, IBoard, ICoords, IPlayer, IScores } from "../../types";
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
  scores: IScores;
  suggestions: ICoords[];
  aiIsThinking: boolean;
  aiResponseTime: number;
}

function getNextPlayer(p: IPlayer) {
  return p == IPlayer.Black ? IPlayer.White : IPlayer.Black;
}

function mergeBoardWithSuggestions(
  board: IBoard,
  suggestions: ICoords[]
): IBoard {
  if (!suggestions.length) return board;
  const merged = board.map(row => row.slice()) as IBoard;
  suggestions.forEach(({ x, y }, i) => (merged[y][x] = 3 + i));
  return merged;
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
    scores: [undefined, 0, 0],
    suggestions: [],
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

  aiFetch = () => {
    this.setState({
      aiIsThinking: true
    });
    const now = Date.now();
    return API.suggestMoves(this.state).then(moves => {
      this.setState({
        aiIsThinking: false,
        aiResponseTime: Date.now() - now
      });
      return moves;
    });
  };

  aiMove = () => {
    this.setState({
      suggestions: []
    });
    this.aiFetch().then((moves: ICoords[]) => {
      this.placeStone(moves[0]);
    });
  };

  showSuggestions = () => {
    this.aiFetch().then(suggestions => {
      this.setState({
        suggestions
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
      () => {
        if (nextPlayer == aiPlayer) return this.aiMove();
        this.showSuggestions();
      }
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
    const { board, player, suggestions, aiResponseTime } = this.state;
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
        <Board onClick={this.handleCellClick}>
          {mergeBoardWithSuggestions(board, suggestions)}
        </Board>
      </div>
    );
  }
}
