import React, { useEffect, useState } from 'react';
import '../css/PlayersList.css';

const colors = ['red', 'violet'];
const getPlayerColor = (playerID) => colors[playerID % colors.length];;

// players {ID:number, characters:[]int, missTurn:bool,}
const PlayersList = ({players=[]}) => {
    const [playersList, setPlayersList] = useState([]);
    useEffect(() => {
        setPlayersList( [...players] );
    }, [players.length > 0]);

    return (
        <article className='players-list'>
            <h3>Players</h3>
            {playersList.map(p =>
                <div className='player-info' key={`p-${p.ID}`}>
                    <div className='row'>
                        <div className={`i-color-${getPlayerColor(p.ID)}`}></div>
                        <div className='i-name'>
                            Player <strong>{p.ID + 1}</strong>
                        </div>
                    </div>
                    <div className='row'>
                        {p.characters.map(c => <span key={`c-${c}`}>{c + 1}</span>)}
                    </div>
                </div>
            )}
        </article>
    );
}

export default PlayersList;
export {colors, getPlayerColor};