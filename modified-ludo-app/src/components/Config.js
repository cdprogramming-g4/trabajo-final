import React, { useState } from 'react';
import '../css/Config.css';
import { connData, stages } from '../App';

const Config = ({setStage, setGameData}) => {
    const [config, setConfig] = useState({
        width: 8,
        height: 8,
        numPlayers: 4,
        numObstacles: 10,
    });

    const handleChangeWidth = (event)=>{
        setConfig({...config, width: event.target.value});
    };
    const handleChangeHeight = (event)=>{
        setConfig({...config, height: event.target.value});
    };
    const handleChangeNumPlayers = (event)=>{
        setConfig({...config, numPlayers: event.target.value});
    };
    const handleChangeNumObstacles = (event)=>{
        setConfig({...config, numObstacles: event.target.value});
    };

    const sendDataToServer = async () => {
        const currConfig = {
            type: 'config',
            numPlayers: config.numPlayers,
            numObstacles: config.numObstacles,
            size: config.width * config.height,
        };
        console.log('cc', currConfig);
        
        const socket = new WebSocket(`ws://${connData.serverURL}/ws`);

        socket.onopen = () => {
            console.log('Conexión establecida');
            socket.send(JSON.stringify(currConfig));
        };

        socket.onmessage = (event) => {
            // Manejar mensajes recibidos del servidor
            const data = JSON.parse(event.data);
            console.log('Mensaje del servidor:', data);
            if (data.type === 'Board') {
                config.board = data.array;
                console.log('config', config);
                setGameData({...config});
                setStage(stages.GAME);
            }
            socket.close();
        };

        socket.onclose = () => {
            console.log('Conexión cerrada');
        };
    };

    const handleSubmit = (event)=>{
        event.preventDefault();
        sendDataToServer();
    };

    return (
        <section className='config col'>
            <h1>Configuración</h1>
            
            <form onSubmit={handleSubmit} className='card'>
                <div className='row'>
                    <div className='input'>
                        <label>Ancho</label>
                        <input id='board-width' name='board-width'
                            type='number' min={3} max={10}
                            value={config?.width}
                            onInput={handleChangeWidth}
                        />
                    </div>
                    <div className='input'>
                        <label>Alto</label>
                        <input id='board-height' name='board-height'
                            type='number' min={3} max={10}
                            value={config?.height}
                            onInput={handleChangeHeight}
                        />
                    </div>
                </div>

                <div className='input'>
                    <label>Número de jugadores a esperar</label>
                    <input id='board-height' name='board-height'
                        type='number' min={2}
                        value={config?.numPlayers}
                        onInput={handleChangeNumPlayers}
                    />
                </div>

                <div className='input'>
                    <label>Número de obstáculos</label>
                    <input id='board-height' name='board-height'
                        type='number' min={2} max={(config?.width * config?.height)*0.75}
                        value={config?.numObstacles}
                        onInput={handleChangeNumObstacles}
                    />
                </div>
                
                <button type='submit'>¡JUGAR!</button>
            </form>
        </section>
    );
};

export default Config;