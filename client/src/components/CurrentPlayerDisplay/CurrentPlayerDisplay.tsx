import * as React from "react";
import { IPlayer } from "../../types";
import CurrentPlayerDisplayStyles from "./CurrentPlayerDisplay.module.css";
import c from "classnames";

interface ICurrentPlayerDisplayProps {
  player: IPlayer;
  onClick(p: IPlayer): void;
  allowPlayerSelection: boolean;
}

const CurrentPlayerDisplay: React.FunctionComponent<
  ICurrentPlayerDisplayProps
> = function CurrentPlayerDisplay({ player, onClick, allowPlayerSelection }) {
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
      />
      <button
        onClick={() => player !== IPlayer.White && onClick(IPlayer.White)}
        className={c(
          CurrentPlayerDisplayStyles.btn,
          CurrentPlayerDisplayStyles.btn__white,
          player === IPlayer.White && CurrentPlayerDisplayStyles.btn__selected
        )}
        disabled={!allowPlayerSelection}
      />
    </div>
  );
};

export default CurrentPlayerDisplay;
