import React from "react";
import Board from "../Board/Board";
import CurrentPlayer from "../CurrentPlayerDisplay";
import WinnerModal from "./WinnerModal";
import {
  GameType,
  IBoard,
  ICommonGameState,
  ICoords,
  IPlayer,
  ISuggestion
} from "../../types";
import assocPath from "lodash/fp/assocPath";
import GameStyles from "./Game.module.css";
import * as API from "../../API";
import pick from "lodash/fp/pick";
import HistoryControls from "./HIstoryControls";

export interface IGameProps {
  type: GameType;
  aiPlayer?: IPlayer;
}

interface IGameState extends ICommonGameState {
  history: ICommonGameState[];
  historyIndex: number;
  suggestions: ISuggestion[];
  aiIsThinking: boolean;
  aiResponseTime: number;
}

function mergeBoardWithSuggestions(
  board: IBoard,
  suggestions: ISuggestion[]
): IBoard {
  if (!suggestions.length) return board;
  const merged = board.map(row => row.slice()) as IBoard;
  suggestions.forEach(
    ({ x, y, evaluation }) => (merged[y][x] = 3 + evaluation)
  );
  return merged;
}

const INITIAL_STATE: IGameState = {
  player: IPlayer.Black,
  board: Board.init(),
  blackScore: 0,
  whiteScore: 0,
  winner: 0,
  suggestions: [],
  history: [],
  historyIndex: 0,
  aiIsThinking: false,
  aiResponseTime: 0
};

export default class Game extends React.Component<IGameProps, IGameState> {
  resetGame = () =>
    this.setState(INITIAL_STATE, () => {
      const { type, aiPlayer } = this.props;
      const { player } = this.state;
      if (type === GameType.vsComputer && player === aiPlayer) this.aiMove();
    });

  componentWillMount = this.resetGame;

  setPlayer = (player: IPlayer) => {
    this.props.type === GameType.debug && this.setState({ player });
  };

  humanMove = (coords: ICoords) => {
    !this.state.winner &&
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
    if (this.state.winner) return;
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
    !this.state.winner &&
    this.setState(
      oldState => ({
        history: oldState.history.concat(
          pick(
            ["board", "player", "blackScore", "whiteScore", "winner"],
            oldState
          )
        ),
        ...state
      }),
      () => {
        if (state.player == this.props.aiPlayer) this.aiMove();
        else this.showSuggestions();
      }
    );

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
    const {
      board,
      player,
      blackScore,
      whiteScore,
      winner,
      suggestions,
      aiResponseTime
    } = this.state;
    return (
      <div className={GameStyles.container}>
        <div className={GameStyles.centerWrapper}>
          <div className={GameStyles.infoPanel}>
            <CurrentPlayer
              player={player}
              blackScore={blackScore}
              whiteScore={whiteScore}
              onClick={this.setPlayer}
              allowPlayerSelection={type == GameType.debug}
            />
            {type === GameType.vsComputer && !!aiResponseTime && (
              <div className={GameStyles.responseTime}>
                AI Response Time: {aiResponseTime}ms
              </div>
            )}
            <HistoryControls/>
          </div>
          <Board onClick={this.handleCellClick}>
            {mergeBoardWithSuggestions(board, suggestions)}
          </Board>
        </div>
        <WinnerModal winner={winner} onRestart={this.resetGame} />
      </div>
    );
  }
}
