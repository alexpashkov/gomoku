import { ICommonGameStateWithBoard, ICoords, ISuggestion } from "./types";

const BASE_URL = process.env.REACT_APP_BASE_URL;

export function makeMove(
  state: ICommonGameStateWithBoard,
  coords: ICoords
): Promise<ICommonGameStateWithBoard> {
  const cell = state.board[coords.y][coords.x];
  if (cell) return Promise.reject("cell is occupied");
  return Promise.reject("test");
  // fetch(BASE_URL + "/make-move", {
  //   method: "POST",
  //   body: JSON.stringify({
  //     state,
  //     coords
  //   })
  // }).then(res => (res.status == 200 ? res.json() : Promise.reject(res.json())));
}

export function suggestMoves(
  state: ICommonGameStateWithBoard
): Promise<ISuggestion[]> {
  return fetch(BASE_URL + "/suggest-move", {
    method: "POST",
    body: JSON.stringify(state)
  }).then(res => (res.status === 200 ? res.json() : Promise.reject()));
}
