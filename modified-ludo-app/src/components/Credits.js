import React from 'react';
import '../css/Credits.css';
import { stages } from '../App';

const Credits = ({setStage}) => {
    
    const handleClick = (event)=>{
        event.preventDefault();
        setStage(stages.STATS);
    };
    
    return (
        <section className='credits col'>
            <h1>Créditos</h1>
            <ul className='card'>
                <li>Anto Chávez, Carolain</li>
                <li>Maguiña Bernuy, Richard José</li>
                <li>Atarama Leon, Diego Sebastian</li>
            </ul>
            <button onClick={handleClick}>¡Listo!</button>
        </section>
    )
};

export default Credits;