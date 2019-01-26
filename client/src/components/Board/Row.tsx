import * as React from "react";
import { IBoardRow, ICellClickHandler } from "../../types";
import Cell from "./Cell";
import BoardStyles from "./Board.module.css";

interface IRowProps {
  children: IBoardRow;
  y: number;
  onClick: ICellClickHandler;
}

class Row extends React.PureComponent<IRowProps> {
  render(): React.ReactNode {
    const { children: cells, y, onClick } = this.props;
    return (
      <div className={BoardStyles.row}>
        {cells.map((cell, x) => (
          <Cell x={x} y={y} onClick={onClick} key={x}>
            {cell}
          </Cell>
        ))}
      </div>
    );
  }
}

export default Row;
