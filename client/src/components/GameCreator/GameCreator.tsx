import React from "react";
import { Link } from "react-router-dom";
import { GameType } from "../Game/Game";

const GameCreator: React.FunctionComponent<{}> = function() {
  const gameId = Date.now() + "" + Math.floor(Math.random() * 1e10);
  return (
    <React.Fragment>
      <Link to={`/game/${gameId}/${GameType.vsFriend}`}>Play a Friend</Link>
      <br />
      <Link to={`/game/${gameId}/${GameType.vsComputer}`}>Play Computer</Link>
    </React.Fragment>
  );
};

export default GameCreator;
