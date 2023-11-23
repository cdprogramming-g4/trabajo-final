import React, { useEffect, useState } from 'react';
import '../css/Stats.css';
import { stages } from '../App';
import { getPlayerColor } from './PlayersList';

const Stats = ({setStage}) => {
    const [players, setPlayers] = useState([
        {ID: 0, characters:[0, 4, 8, 16]},
        {ID: 1, characters:[0, 1, 4]},
        {ID: 2, characters:[0, 2, 6]},
        {ID: 3, characters:[0, 3, 9]},
        {ID: 4, characters:[0, 3, 9]},
        {ID: 5, characters:[0, 3, 9]},
        {ID: 6, characters:[0, 3, 9]},
        {ID: 7, characters:[0, 3, 9]},
        {ID: 8, characters:[0, 3, 9]},
        {ID: 9, characters:[0, 3, 9]},
        {ID: 10, characters:[0, 3, 9]},
        {ID: 11, characters:[0, 3, 9]},
        {ID: 12, characters:[0, 3, 9]},
        {ID: 13, characters:[0, 3, 9]},
        {ID: 14, characters:[0, 3, 9]},
        {ID: 15, characters:[0, 3, 9]},
    ]);

    const [rankingPlayers, setRankingPlayers] = useState([]);

    useEffect(() => {
        let _rankingPlayers = [];
        players.forEach(player => {
            let sumPos = 0;
            let _rankingChars = player.characters.map((c, i)=>{return {ID:i, pos:c}});
            player.characters.forEach( (c, i) => {
                sumPos += c;
            });
            _rankingPlayers.push({
                ID: player.ID,
                sumPos: sumPos,
                chars: _rankingChars.sort((c1, c2)=> c2.pos - c1.pos),
            });
        });
        _rankingPlayers.sort((p1, p2) => p2.sumPos - p1.sumPos);
        setRankingPlayers(_rankingPlayers);
        
    }, [players]);

    const handleClick = (event)=>{
        event.preventDefault();
        setStage(stages.CREDITS);
    };

    const handleNew = (event)=>{
        event.preventDefault();
        setStage(stages.CONFIG);
    };

    return (
        <section className='stats col'>
            <h1>Estadísticas</h1>

            <div className='row highest'>
                <div className='col'>
                    <h5>3° puesto</h5>
                    <p className='circle smallest'>{rankingPlayers[2]?.ID}</p>
                </div>
                <div className='col'>
                    <h4>Mejor jugador</h4>
                    <p className='circle big'>{rankingPlayers[0]?.ID}</p>
                </div>
                <div className='col'>
                    <h5>2° puesto</h5>
                    <p className='circle small'>{rankingPlayers[1]?.ID}</p>
                </div>
            </div>

            <div className='col card merit'>
                <h5>Órdenes de mérito</h5>
                <ol>
                    {rankingPlayers.map(p =>
                    <li key={`p-${p.ID}`}>
                        <strong>{p.ID}</strong>
                        <div className='row'>
                            {p.chars.map(c =>
                            <span key={`c-${p.ID}-${c.ID}`} className='row'>
                                <img
                                    src={`images/characters/char_${getPlayerColor(p.ID)}_${c.ID+1}.png`}
                                    alt={`player ${p.ID} character ${c.ID+1}`}
                                />
                                {c.pos}
                            </span>
                            )}
                        </div>
                    </li>
                    )}
                </ol>
            </div>

            <div className='row'>
                <button onClick={handleNew}>Nuevo juego</button>
                <button onClick={handleClick}>Ver créditos</button>
            </div>
        </section>
    )
};

export default Stats;