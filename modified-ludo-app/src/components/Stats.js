import React from 'react';
import '../css/Stats.css';
import { stages } from '../App';

const Stats = ({setStage}) => {
    const handleClick = (event)=>{
        event.preventDefault();
        setStage(stages.CREDITS);
    };

    return (
        <section className='stats col'>
            <h1>Estadísticas</h1>

            <div className='row highest'>
                <div className='col'>
                    <h5>3° puesto</h5>
                    <p className='circle smallest'>3</p>
                </div>
                <div className='col'>
                    <h4>Mejor jugador</h4>
                    <p className='circle big'>1</p>
                </div>
                <div className='col'>
                    <h5>2° puesto</h5>
                    <p className='circle small'>2</p>
                </div>
            </div>

            <div className='col card merit'>
                <h5>Órdenes de mérito</h5>
                <ol>
                    <li></li>
                </ol>
            </div>

            <button onClick={handleClick}>Ver créditos</button>
        </section>
  )
};

export default Stats;