import * as React from "react";
import { IPlayer } from "../../types";
import CurrentPlayerDisplayStyles from "./CurrentPlayerDisplay.module.css";
import c from "classnames";

interface ICurrentPlayerDisplayProps {
  player: IPlayer;
  blackScore: number;
  whiteScore: number;
  onClick(p: IPlayer): void;
  allowPlayerSelection: boolean;
}

const CurrentPlayerDisplay: React.FunctionComponent<
  ICurrentPlayerDisplayProps> = function CurrentPlayerDisplay({
                                                                player,
                                                                blackScore,
                                                                whiteScore,
                                                                onClick,
                                                                allowPlayerSelection
                                                              }) {
  return (
    <div>
      <button
        onClick={() => player !== IPlayer.Black && onClick(IPlayer.Black)}
        className={c(
          CurrentPlayerDisplayStyles.btn,
          CurrentPlayerDisplayStyles.btn__black,
          player === IPlayer.Black && CurrentPlayerDisplayStyles.btn__selected
        )}
        disabled={!allowPlayerSelection}
      >
        {blackScore}
      </button>
      <button
        onClick={() => player !== IPlayer.White && onClick(IPlayer.White)}
        className={c(
          CurrentPlayerDisplayStyles.btn,
          CurrentPlayerDisplayStyles.btn__white,
          player === IPlayer.White && CurrentPlayerDisplayStyles.btn__selected
        )}
        disabled={!allowPlayerSelection}
      >
        {whiteScore}
      </button>
    </div>
  );
};

export default CurrentPlayerDisplay;
