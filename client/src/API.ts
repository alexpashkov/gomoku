import axios from "axios";
import { ICommonGameStateWithBoard, ICoords, ISuggestion } from "./types";

const BASE_URL = process.env.REACT_APP_BASE_URL;

export function makeMove(
  state: ICommonGameStateWithBoard,
  coords: ICoords
): Promise<ICommonGameStateWithBoard> {
  const cell = state.board[coords.y][coords.x];
  if (cell) return Promise.reject("cell is occupied");
  return axios
    .post(BASE_URL + "/make-move", {
      state,
      coords
    })
    .then(({ data }) => data, ({ response }) => Promise.reject(response.data));
}

export function suggestMoves(
  state: ICommonGameStateWithBoard,
  difficulty: number
): Promise<ISuggestion[]> {
  return axios
    .post(`${BASE_URL}/suggest-moves?difficulty=${difficulty}`, state)
    .then(({ data }) => data, ({ response }) => Promise.reject(response.data));
}
