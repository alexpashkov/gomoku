import * as React from "react";
import { IBoardRow, ICellClickHandler } from "../../types";
import Cell from "./Cell";

interface IRowProps {
  children: IBoardRow;
  y: number;
  onClick: ICellClickHandler;
}

const Row: React.FunctionComponent<IRowProps> = function Row({
  children: cells,
  y,
  onClick
}) {
  return (
    <div className="board__row">
      {cells.map((cell, x) => (
        <Cell x={x} y={y} onClick={onClick} key={x}>
          {cell}
        </Cell>
      ))}
    </div>
  );
};

export default Row;
