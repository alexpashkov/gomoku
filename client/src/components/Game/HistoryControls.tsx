import React from "react";

interface IHistoryControlsProps {
  col: any[];
  i: number;
  onChange(i: number): void;
}

function forwardDisabled(i: number, col: any[]) {
  return i >= col.length - 1;
}

const HistoryControls: React.FunctionComponent<IHistoryControlsProps> = ({
                                                                           i,
                                                                           col,
                                                                           onChange
                                                                         }) => (
  <div>
    <button disabled={!i} onClick={() => onChange(0)}>
      {"<<"}
    </button>
    <button disabled={!i} onClick={() => onChange(i - 1)}>
      {"<"}
    </button>
    <button disabled={forwardDisabled(i, col)} onClick={() => onChange(i + 1)}>
      {">"}
    </button>
    <button
      disabled={forwardDisabled(i, col)}
      onClick={() => onChange(col.length - 1)}
    >
      {">>"}
    </button>
  </div>
);

export default HistoryControls;
