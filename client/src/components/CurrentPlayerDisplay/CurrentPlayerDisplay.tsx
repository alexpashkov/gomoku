import * as React from "react";
import { IPlayer } from "../../types";

interface ICurrentPlayerDisplayProps {
  player: IPlayer;
  onClick(p: IPlayer): any;
}

const CurrentPlayerDisplay: React.FunctionComponent<
  ICurrentPlayerDisplayProps
> = function CurrentPlayerDisplay({ player, onClick }) {
  return (
    <div className="current-player">
      <button
        onClick={() => player !== 1 && onClick(1)}
        className={
          "current-player__btn current-player__btn--white" +
          (player === 1 ? " current-player__btn--selected" : "")
        }
      />
      <button
        onClick={() => player !== 2 && onClick(2)}
        className={
          "current-player__btn current-player__btn--black" +
          (player === 2 ? " current-player__btn--selected" : "")
        }
      />
    </div>
  );
};

export default CurrentPlayerDisplay;
