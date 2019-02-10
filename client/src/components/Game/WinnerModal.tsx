import React from "react";
import { IPlayer } from "../../types";
import GameStyles from "./Game.module.css";

const WinnerModal: React.FunctionComponent<{
  winner: IPlayer;
  onRestart(): void;
}> = ({ winner, onRestart }) =>
  !!winner && (
    <div className={GameStyles.winnerModal}>
      <div>{(winner == IPlayer.Black ? "Black" : "White") + " won!"}</div>
      <button onClick={onRestart}>Restart Game</button>
    </div>
  );

export default WinnerModal;
