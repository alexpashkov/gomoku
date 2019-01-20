export enum IPlayer {
  Black, White
}

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
