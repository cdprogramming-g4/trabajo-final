import React, { useEffect, useState } from 'react';
import '../css/Board.css';
import Wall from './squares/Wall';
import Trap from './squares/Trap';
import Creature from './squares/Creature';
import Path from './squares/Path';

const square = {
  PATH     : 0,
  WALL     : 1,
  TRAP     : 2,
  CREATURE : 3,
};

const Square = ({cell=0, x=0, y=0, size={w:0, h:0}}) => {
  const [cellJsx, setCellJsx] = useState(<></>);

  useEffect(()=>{
    switch(cell) {
      case square.WALL:
        setCellJsx(<Wall h={size.h-1} w={size.w-1} x={x} y={y}/>);
      break;
      case square.TRAP:
        setCellJsx(<Trap/>);
      break;
      case square.CREATURE:
        setCellJsx(<Creature/>);
      break;
      default:
        setCellJsx(<></>);
      break;
    }
  }, [cell, size]);

  return (<div className='board-cell'>{cellJsx}</div>);
};

const Board = ({board=[[]], players=[]}) => {
  
  const [size, setSize] = useState({w: 0, h:0});
  
  useEffect(()=>{
    const _size = {
      h: board.length,
      w: board[0].length,
    };
    setSize(_size);
    console.log(board, _size);
  }, [board.length > 0]);
  
  return (
    <div className='board'>
      <div className='board-background'
        style={{
          backgroundImage: 'linear-gradient(30deg, #8fa6f140, #ffffff60), url(/images/board/path.jpg)',
        }}
      ></div>

      {board.map((row, y)=>{
        return <div className='board-row' key={`r-${y}`}>
          {row.map((cell, x) =>
            cell === square.PATH ?
              <Path key={`p-${x}.${y}`}
                players={players}
                w={size.w} x={x} y={y}
              /> :
              <Square key={`b-${x}.${y}`}
                cell={cell} size={size}
                x={x} y={y}
              />
          )}
        </div>
      })}
    </div>
  )
};

export default Board;