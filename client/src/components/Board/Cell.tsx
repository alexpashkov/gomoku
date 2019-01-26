import * as React from "react";
import { IBoardCell, ICellClickHandler, IPlayer } from "../../types";
import BoardStyles from "./Board.module.css";

interface ICellProps {
  children: IBoardCell;
  x: number;
  y: number;
  onClick: ICellClickHandler;
}

function stoneClassName(c: IBoardCell): string {
  switch (c) {
    case 0:
      return BoardStyles["stone--none"];
    case IPlayer.Black:
      return BoardStyles["stone--black"];
    case IPlayer.White:
      return BoardStyles["stone--white"];
    default:
      throw new Error("Invalid board cell value " + c);
  }
}

class Cell extends React.PureComponent<ICellProps> {
  render() {
    const { children: cell, x, y, onClick } = this.props;
    return (
      <div className={BoardStyles.cell} onClick={() => onClick({ x, y })}>
        <div className={`${BoardStyles.stone} ${stoneClassName(cell)}`} />
      </div>
    );
  }
}

export default Cell;
