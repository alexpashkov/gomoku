import * as React from "react";
import { IBoardCell, ICellClickHandler, IPlayer } from "../../types";

interface ICellProps {
  children: IBoardCell;
  x: number;
  y: number;
  onClick: ICellClickHandler;
}

function stoneModifier(c: IBoardCell): string {
  switch (c) {
    case 0:
      return "none";
    case IPlayer.Black:
      return "black";
    case IPlayer.White:
      return "white";
    default:
      throw new Error("Invalid board cell value " + c);
  }
}

class Cell extends React.PureComponent<ICellProps> {
  render() {
    const { children: cell, x, y, onClick } = this.props;
    return (
      <div className="board__cell" onClick={() => onClick({ x, y })}>
        <div className={`board__stone board__stone--${stoneModifier(cell)}`} />
      </div>
    );
  }
}

export default Cell;
