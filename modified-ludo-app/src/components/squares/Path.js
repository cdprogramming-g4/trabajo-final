import React, { useEffect, useState } from 'react';
import { getPlayerColor } from '../PlayersList';

// players {ID:number, characters:[]int, missTurn:bool,}
const Path = ({players=[], w=0, x=0, y=0}) => {
    const [playersInCell, setPlayersInCell] = useState([]);
    const [percCharSize, setPercCharSize] = useState(0);

    const isCharInPath = (pos = 0) => {
        if (!pos || !w) { return false; }
        const _x = ((pos + 1) % w ) - 1;
        const _y = parseInt(pos / w);
        // console.log('pos', pos, ':', _x, _y);
        return _x === x && _y === y;
    };

    useEffect(()=>{
        let _playersInCell = [];
        let _numCharsInCell = 0;

        players.forEach(p => {
            // Getting characters IDs
            const charsInCell = p.characters
                        .map((c, i) => isCharInPath(c) ? i : undefined)
                        .filter(i => i !== undefined);
            
            if (charsInCell.length > 0) {
                _numCharsInCell += charsInCell.length;
                console.log('player', p, charsInCell, '|', x, y);
                _playersInCell.push(
                    {ID: p.ID, missTurn: p.missTurn,
                    characters:charsInCell, color: getPlayerColor(p.ID)}
                );
            }
        });
        
        setPercCharSize((_numCharsInCell > 0) ? 1/_numCharsInCell : _numCharsInCell);
        setPlayersInCell(_playersInCell);
        
    }, [players.length > 0, w > 0, x, y]);
    
    return (
        <div className='board-cell path'>
            {playersInCell.map(p =>
                {return p.characters.map(c => 
                    <img key={`p${p.ID}-char${c}`}
                        className={`character ${p.missTurn?'disabled':''}`}
                        style={{
                            width: `calc(${percCharSize} * 100%)`,
                            margin: `calc(${percCharSize} * -20%)
                                     calc(${percCharSize} * -33%)`
                        }}
                        src={`images/characters/char_${p.color}_${c+1}.png`}
                        alt={`player ${p.ID} character ${c+1}`}
                    />
                )}
            )}
        </div>
    )
};

export default Path;