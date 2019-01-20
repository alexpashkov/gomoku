import * as React from "react";
import { IPlayer } from "../../types";

interface ICurrentPlayerDisplayProps {
  player: IPlayer;
  onClick(p: IPlayer): void;
}

const CurrentPlayerDisplay: React.FunctionComponent<
  ICurrentPlayerDisplayProps
> = function CurrentPlayerDisplay({ player, onClick }) {
  return (
    <div className="current-player">
      <button
        onClick={() => player !== IPlayer.Black && onClick(IPlayer.Black)}
        className={
          "current-player__btn current-player__btn--black" +
          (player === IPlayer.Black ? " current-player__btn--selected" : "")
        }
      />
      <button
        onClick={() => player !== IPlayer.White && onClick(IPlayer.White)}
        className={
          "current-player__btn current-player__btn--white" +
          (player === IPlayer.White ? " current-player__btn--selected" : "")
        }
      />
    </div>
  );
};

export default CurrentPlayerDisplay;
