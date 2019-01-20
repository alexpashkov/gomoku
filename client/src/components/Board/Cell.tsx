import * as React from "react";
import { IBoardCell, ICellClickHandler, ICoords } from "../../types";

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
    case 1:
      return "white";
    case 2:
      return "black";
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
