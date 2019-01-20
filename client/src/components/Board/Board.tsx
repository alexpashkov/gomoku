import * as React from "react";
import Row from "./Row";
import { IBoard, ICellClickHandler } from "../../types";

interface IBoardProps {
  children: IBoard;
  onClick: ICellClickHandler;
}

const Board: React.FunctionComponent<IBoardProps> = function Board({
  children: rows,
  onClick
}) {
  return (
    <div className="board">
      {rows.map((row, y) => (
        <Row y={y} onClick={onClick} key={y}>
          {row}
        </Row>
      ))}
    </div>
  );
};

export default Board;
