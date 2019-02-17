import React from "react";
import c from "classnames";
import GameStyles from "./Game.module.css";
import HistoryControlsStyles from "./HistoryControls.module.css";

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
    <button
      disabled={!i}
      onClick={() => onChange(0)}
      className={c(
        HistoryControlsStyles.btn,
        HistoryControlsStyles.btn__doubleLeft
      )}
    />
    <button
      disabled={!i}
      onClick={() => onChange(i - 1)}
      className={c(HistoryControlsStyles.btn, HistoryControlsStyles.btn__left)}
    />
    <button
      disabled={i >= max}
      onClick={() => onChange(i + 1)}
      className={c(HistoryControlsStyles.btn, HistoryControlsStyles.btn__right)}
    />
    <button
      disabled={i >= max}
      onClick={() => onChange(max)}
      className={c(
        HistoryControlsStyles.btn,
        HistoryControlsStyles.btn__doubleRight
      )}
    />
  </div>
);

export default HistoryControls;
