import React from "react";
import { Link } from "react-router-dom";

const GameCreator: React.FunctionComponent<{}> = function() {
  const gameId = Date.now() + "" + Math.floor(Math.random() * 1e10);
  return (
    <React.Fragment>
      <Link to={`/game/${gameId}/vsFriend`}>Play a Friend</Link>
      <Link to={`/game/${gameId}/vsComp`}>Play Computer</Link>
    </React.Fragment>
  );
};

export default GameCreator;
