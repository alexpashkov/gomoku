import * as React from "react";
import { IBoardCell, ICellClickHandler, IPlayer } from "../../types";

interface ICellProps {
  children: IBoardCell;
  x: number;
  y: number;
  onClick: ICellClickHandler;
}

function stoneModifier(c: IBoardCell) {
  switch (c) {
    case 0:
      return "none";
    case IPlayer.Black:
      return "black";
    case IPlayer.White:
      return "white";
  }
}

const Cell: React.FunctionComponent<ICellProps> = function Cell({
  children: cell,
  x,
  y,
  onClick
}) {
  return (
    <div className="board__cell" onClick={() => onClick({ x, y })}>
      <div className={`board__stone board__stone--${stoneModifier(cell)}`} />
    </div>
  );
};

export default Cell;
