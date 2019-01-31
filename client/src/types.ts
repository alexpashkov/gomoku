export enum GameType {
  vsFriend = "vsFriend",
  vsComputer = "vsComputer",
  debug = "debug"
}

export enum IPlayer {
  Black = 1,
  White = 2
}

export type IScores = [number, number];

export type IBoardCell = 0 | IPlayer;

export type IBoardRow = [
  IBoardCell,
  IBoardCell,
  IBoardCell,
  IBoardCell,
  IBoardCell,
  IBoardCell,
  IBoardCell,
  IBoardCell,
  IBoardCell,
  IBoardCell,
  IBoardCell,
  IBoardCell,
  IBoardCell,
  IBoardCell,
  IBoardCell,
  IBoardCell,
  IBoardCell,
  IBoardCell,
  IBoardCell
];

export type IBoard = [
  IBoardRow,
  IBoardRow,
  IBoardRow,
  IBoardRow,
  IBoardRow,
  IBoardRow,
  IBoardRow,
  IBoardRow,
  IBoardRow,
  IBoardRow,
  IBoardRow,
  IBoardRow,
  IBoardRow,
  IBoardRow,
  IBoardRow,
  IBoardRow,
  IBoardRow,
  IBoardRow,
  IBoardRow
];

export interface ICoords {
  x: number;
  y: number;
}

export type ICellClickHandler = (coords: ICoords) => void;
