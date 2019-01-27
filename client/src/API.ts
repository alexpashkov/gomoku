import { IBoard, ICoords, IPlayer, IScores } from "./types";

const BASE_URL = "http://localhost:4444";

export function validateMove(
  board: IBoard,
  player: IPlayer,
  coords: ICoords
): Promise<boolean> {
  const cell = board[coords.y][coords.x];
  return Promise.resolve(!cell);
}

export function suggestMove(params: {
  board: IBoard;
  player: IPlayer;
  blackScore: number;
  whiteScore: number;
}): Promise<ICoords> {
  return fetch(BASE_URL + "/suggest-move", {
    method: "POST",
    body: JSON.stringify(params)
  }).then(res => (res.status === 200 ? res.json() : Promise.reject()));
}
