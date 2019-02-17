import React from "react";
import GameStyles from "./Game.module.css";

interface IHistoryControlsProps {
  i: number;
  max: number;
  onChange(i: number): void;
}

const HistoryControls: React.FunctionComponent<IHistoryControlsProps> = ({
  i,
  max,
  onChange
}) => (
  <div className={GameStyles.historyControls}>
    <button disabled={!i} onClick={() => onChange(0)}>
      {"<<"}
    </button>
    <button disabled={!i} onClick={() => onChange(i - 1)}>
      {"<"}
    </button>
    <button disabled={i >= max} onClick={() => onChange(i + 1)}>
      {">"}
    </button>
    <button disabled={i >= max} onClick={() => onChange(max)}>
      {">>"}
    </button>
  </div>
);

export default HistoryControls;
