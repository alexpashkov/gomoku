import React from "react";
import "./Board.css";

function resolveMarkModifier(mark) {
  switch (mark) {
    case 0:
      return "empty";
    case 1:
      return "white";
    case 2:
      return "black";
  }
}

function Cell({ children: mark, x, y, onClick }) {
  return (
    <div className="board__cell" onClick={() => onClick({ x, y })}>
      <div
        className={`board__mark board__mark--${resolveMarkModifier(mark)}`}
      />
    </div>
  );
}

function Row({ children: cells, y, onClick }) {
  return (
    <div className="board__row">
      {cells.map((cell, x) => (
        <Cell x={x} y={y} onClick={onClick} key={x}>
          {cell}
        </Cell>
      ))}
    </div>
  );
}

export default function Board({ children: rows, onClick }) {
  return (
    <div className="board">
      {rows.map((row, y) => (
        <Row y={y} onClick={onClick} key={y}>
          {row}
        </Row>
      ))}
    </div>
  );
}
