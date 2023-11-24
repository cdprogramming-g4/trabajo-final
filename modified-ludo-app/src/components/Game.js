import React, { useEffect, useState } from 'react';
import PlayersList from './PlayersList';
import Board from './Board';
import { stages } from '../App';

const Game = ({setStage, gameData}) => {
    // board es la matriz final
    const [board, setBoard] = useState([]);
    // gameData.Board es lo que se lee desde servidor (para reconocer obstaculos)
    // opcional: esa configuración también podría ser enviada desde react.
    //     [
    //     1,1,1,0,0,0,1,1,
    //     0,0,0,0,0,0,0,0,
    //     0,0,2,0,0,0,0,0,
    //     0,0,0,3,1,0,0,1,
    //     1,0,0,0,0,0,0,0,
    //     1,0,0,0,0,1,1,1,
    // ]
        
    const [players, setPlayers] = useState([]);
    // {ID: 0, characters:[0, 4, 8, 16], missTurn:true,},
    // {ID: 1, characters:[0, 1, 4], missTurn:false,},
    // {ID: 2, characters:[0, 2, 6], missTurn:true,},
    // {ID: 3, characters:[0, 3, 9], missTurn:false,},
    // {ID: 4, characters:[0, 3, 9], missTurn:false,},
    // {ID: 5, characters:[0, 3, 9], missTurn:false,},
    // {ID: 6, characters:[0, 3, 9], missTurn:false,},
    // {ID: 7, characters:[0, 3, 9], missTurn:false,},
    // {ID: 8, characters:[0, 3, 9], missTurn:false,},
    // {ID: 9, characters:[0, 3, 9], missTurn:false,},
    // {ID: 10, characters:[0, 3, 9], missTurn:false,},
    // {ID: 11, characters:[0, 3, 9], missTurn:false,},
    // {ID: 12, characters:[0, 3, 9], missTurn:false,},
    // {ID: 13, characters:[0, 3, 9], missTurn:false,},
    // {ID: 14, characters:[0, 3, 9], missTurn:false,},
    // {ID: 15, characters:[0, 3, 9], missTurn:false,},

    const [win, setWin] = useState(true);

    const handleClick = (event)=>{
        event.preventDefault();
        setStage(stages.STATS);
    };

    const TwoDimensional = (arr, width) => {
      let res = []; 
      for(var i=0;i < arr.length;i = i+width)
          res.push(arr.slice(i,i+width));
      return res;
    };
    useEffect(()=>{
        // width and height already known
        if (gameData) {
            if (gameData.board) {
                const arr = TwoDimensional(gameData.board, gameData.width);
                console.log('Size by game data', arr);
                setBoard([...arr]);
            }
            else {
                const arr = Array(gameData.height??0).fill(Array(gameData.width??0).fill(0));
                console.log('Fill array with nulls', arr);
                setBoard([...arr]);
            }
        }
    }, [gameData]);
    
    useEffect(()=>{
        //TODO
    });

    return (
        <section className='game'>
            <PlayersList players={players}/>
            <div className='col'>
                <h1>Modified Ludo</h1>
                
                <Board board={board} players={players}/>
                
                { win && 
                    <button onClick={handleClick}>
                        ¡Ver resultados!
                    </button>
                }
            </div>
        </section>
    )
};

export default Game;