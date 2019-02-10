import { IBoard, ICommonGameState, ICoords, ISuggestion } from "./types";

const BASE_URL = process.env.REACT_APP_BASE_URL;

export function sendBoard(board: IBoard) {
  fetch(BASE_URL + "/board", {
    method: "POST",
    body: JSON.stringify(board)
  });
}

export function makeMove(state: ICommonGameState, coords: ICoords): Promise<ICommonGameState> {
  const cell = state.board[coords.y][coords.x];
  if (cell) return Promise.reject("cell is occupied");
  return fetch(BASE_URL + "/make-move", {
    method: "POST",
    body: JSON.stringify({
      state,
      coords
    })
  }).then(res => res.json());
}

export function suggestMoves(state: ICommonGameState): Promise<ISuggestion[]> {
  return fetch(BASE_URL + "/suggest-move", {
    method: "POST",
    body: JSON.stringify(state)
  }).then(res => (res.status === 200 ? res.json() : Promise.reject()));
}
