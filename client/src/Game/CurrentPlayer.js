import React from "react";
import "./CurrentPlayer.css";

export default function CurrentPlayer({ player, onClick }) {
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
}
