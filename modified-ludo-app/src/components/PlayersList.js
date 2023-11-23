import React, { useEffect, useState } from 'react';
import '../css/PlayersList.css';

const colors = ['red', 'violet'];
const getPlayerColor = (playerID) => colors[playerID % colors.length];;

// players {ID:number, characters:[]int, missTurn:bool,}
const PlayersList = ({players=[]}) => {
    const [playersList, setPlayersList] = useState([]);
    useEffect(() => {
        setPlayersList( players.map(p => p.ID) );
    }, [players.length > 0]);

    return (
        <article className='players-list'>
            {playersList.map(pID =>
                <div className='player-info' key={`p-${pID}`}>
                    <div className={`i-color-${getPlayerColor(pID)}`}></div>
                    <div className='i-name'>
                        Player <strong>{pID+1}</strong>
                    </div>
                </div>
            )}
        </article>
    );
}

export default PlayersList;
export {colors, getPlayerColor};