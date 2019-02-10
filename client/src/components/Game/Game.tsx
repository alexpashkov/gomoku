import React from "react";
import Board from "../Board/Board";
import CurrentPlayer from "../CurrentPlayerDisplay";
import { GameType, IBoard, ICommonGameState, ICoords, IPlayer, ISuggestion } from "../../types";
import assocPath from "lodash/fp/assocPath";
import GameStyles from "./Game.module.css";
import * as API from "../../API";

export interface IGameProps {
  id: string;
  type: GameType;
  aiPlayer?: IPlayer;
}

interface IGameState extends ICommonGameState {
  suggestions: ISuggestion[];
  aiIsThinking: boolean;
  aiResponseTime: number;
}

function getNextPlayer(p: IPlayer) {
  return p == IPlayer.Black ? IPlayer.White : IPlayer.Black;
}

function mergeBoardWithSuggestions(
  board: IBoard,
  suggestions: ISuggestion[]
): IBoard {
  if (!suggestions.length) return board;
  const merged = board.map(row => row.slice()) as IBoard;
  suggestions.forEach(
    ({ x, y, evaluation }, i) => (merged[y][x] = 3 + evaluation)
  );
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
    blackScore: 0,
    whiteScore: 0,
    suggestions: [],
    aiIsThinking: false,
    aiResponseTime: 0
  };

  setPlayer = (player: IPlayer) => {
    this.props.type === GameType.debug && this.setState({ player });
  };

  humanMove = (coords: ICoords) => {
    API.makeMove(this.state, coords).then(this.setGameState, console.error);
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
    this.aiFetch().then(moves => {
      this.setGameState(moves[0].state);
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

  setGameState = (state: ICommonGameState) =>
    this.setState(state, () => {
      if (state.player == this.props.aiPlayer) this.aiMove();
    });

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
            <div>
              <button onClick={this.sendBoardToServer}>Send To Server</button>
              <button onClick={this.aiMove}>AI Move</button>
            </div>
          )}
        </div>
        <Board onClick={this.handleCellClick}>
          {mergeBoardWithSuggestions(board, suggestions)}
        </Board>
      </div>
    );
  }
}
