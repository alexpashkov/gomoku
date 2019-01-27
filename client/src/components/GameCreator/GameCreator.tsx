import React from "react";
import { Link } from "react-router-dom";
import { GameType, IPlayer } from "../../types";

const GameCreator: React.FunctionComponent<{}> = function() {
  const gameId = Date.now() + "" + Math.floor(Math.random() * 1e10);
  return (
    <React.Fragment>
      <Link to={`/game/${gameId}/${GameType.vsFriend}`}>Play a Friend</Link>
      <br />
      <Link
        to={`/game/${gameId}/${GameType.vsComputer}?aiPlayer=${IPlayer.White}`}
      >
        Play Computer as Black
      </Link>
      <br />
      <Link
        to={`/game/${gameId}/${GameType.vsComputer}?aiPlayer=${IPlayer.Black}`}
      >
        Play Computer as White
      </Link>
      <br />
      <Link to={`/game/${gameId}/${GameType.debug}`}>Debug Mode</Link>
    </React.Fragment>
  );
};

export default GameCreator;
