import { IBoard, ICoords, IPlayer } from "./types";

export function validateMove(
  board: IBoard,
  player: IPlayer,
  coords: ICoords
): Promise<boolean> {
  const cell = board[coords.y][coords.x];
  return Promise.resolve(!cell);
}
