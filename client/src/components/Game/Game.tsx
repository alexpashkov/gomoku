import React from "react";
import Board from "../Board/Board";
import CurrentPlayer from "../CurrentPlayerDisplay";
import WinnerModal from "./WinnerModal";
import {
  GameType,
  IBoard,
  ICommonGameState,
  ICommonGameStateWithBoard,
  ICoords,
  IPlayer,
  ISuggestion
} from "../../types";

import omit from "lodash/fp/omit";
import GameStyles from "./Game.module.css";
import * as API from "../../API";
import HistoryControls from "./HistoryControls";

export interface IGameProps {
  type: GameType;
  aiPlayer?: IPlayer;
}

interface IGameState extends ICommonGameState {
  boardHistory: IBoard[];
  boardHistoryIndex: number;
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
  suggestions.forEach(({ x, y }, i) => (merged[y][x] = 3 + i));
  return merged;
}

const INITIAL_STATE: IGameState = {
  player: IPlayer.Black,
  blackScore: 0,
  whiteScore: 0,
  winner: 0,
  suggestions: [],
  boardHistory: [Board.init()],
  boardHistoryIndex: 0,
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

  getCurrentBoard = () => {
    const { boardHistory } = this.state;
    return boardHistory[boardHistory.length - 1];
  };

  currentBoardIsDisplayed = (): boolean => {
    const { boardHistory, boardHistoryIndex } = this.state;
    return boardHistoryIndex >= boardHistory.length - 1;
  };

  humanMove = (coords: ICoords) => {
    !this.state.winner &&
      API.makeMove(
        {
          ...this.state,
          board: this.getCurrentBoard()
        },
        coords
      ).then(this.setGameState, console.error);
  };

  aiFetch = () => {
    this.setState({
      aiIsThinking: true
    });
    const now = Date.now();
    return API.suggestMoves({
      ...this.state,
      board: this.getCurrentBoard()
    }).then(moves => {
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

  setGameState = (state: ICommonGameStateWithBoard) => {
    const { winner, boardHistory, boardHistoryIndex } = this.state;
    if (winner) return;
    this.setState(
      {
        boardHistory: [...boardHistory, state.board as IBoard],
        boardHistoryIndex: this.currentBoardIsDisplayed()
          ? boardHistoryIndex + 1
          : boardHistoryIndex,
        ...omit(["board"], state)
      },
      () => {
        if (state.player == this.props.aiPlayer) this.aiMove();
        else this.showSuggestions();
      }
    );
  };

  handleCellClick = (coords: ICoords) => {
    if (!this.currentBoardIsDisplayed()) return;
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
    }
  };

  render() {
    const { type } = this.props;
    const {
      player,
      blackScore,
      whiteScore,
      winner,
      suggestions,
      boardHistory,
      boardHistoryIndex,
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
            />
            {type === GameType.vsComputer && !!aiResponseTime && (
              <div className={GameStyles.responseTime}>
                AI Response Time: {aiResponseTime}ms
              </div>
            )}
          </div>
          <Board onClick={this.handleCellClick}>
            {this.currentBoardIsDisplayed()
              ? mergeBoardWithSuggestions(this.getCurrentBoard(), suggestions)
              : boardHistory[boardHistoryIndex]}
          </Board>
          <HistoryControls
            i={boardHistoryIndex}
            max={boardHistory.length - 1}
            onChange={i => {
              this.setState({
                boardHistoryIndex: i
              });
            }}
          />
        </div>
        <WinnerModal winner={winner} onRestart={this.resetGame} />
      </div>
    );
  }
}
