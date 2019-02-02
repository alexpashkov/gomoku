import { IBoard, ICoords, IPlayer, IScores } from "./types";

const BASE_URL = process.env.REACT_APP_BASE_URL;

export function sendBoard(board: IBoard) {
  fetch(BASE_URL + "/board", {
    method: "POST",
    body: JSON.stringify(board)
  });
}

export function validateMove(
  board: IBoard,
  player: IPlayer,
  coords: ICoords
): Promise<boolean> {
  const cell = board[coords.y][coords.x];
  return Promise.resolve(!cell);
}

export function suggestMoves(params: {
  board: IBoard;
  player: IPlayer;
  scores: IScores
}): Promise<ICoords[]> {
  return fetch(BASE_URL + "/suggest-move", {
    method: "POST",
    body: JSON.stringify(params)
  }).then(res => (res.status === 200 ? res.json() : Promise.reject()));
}
