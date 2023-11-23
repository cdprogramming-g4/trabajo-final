import React, { useState } from 'react';
import '../css/Config.css';
import { stages } from '../App';

const Config = ({setStage}) => {
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
            numPlayers: config.numPlayers,
            numObstacles: config.numObstacles,
            size: config.width * config.height,
        };
        console.log('cc', currConfig);
      
        try {
            const serverUrl = 'localhost:8080';
            await fetch('http://localhost:8080/config', {
                mode: 'no-cors',
                method: 'POST',
                headers: {
                    // 'Access-Control-Allow-Origin': '*',
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(currConfig)
            })
            .then(resp => {
                // if (!resp.ok) {
                //     throw `Server error: [${resp.status}] [${resp.statusText}] [${resp.url}]`;
                // }
                // return resp.json()
                setStage(stages.GAME);
            })
            .then((data) => {
                const responseData = data;
                console.log('Datos recibidos del servidor:', responseData)
                console.log(data ? JSON.parse(data) : {})
            })
            .catch((error) => {
                console.log('NOT OK', error)
                // reject(error)
            });

        } catch (error) {
          console.error('Error:', error);
        }
    };

    const handleSubmit = (event)=>{
        event.preventDefault();
        // const message = JSON.stringify({tag: 'config', message: currConfig})+'\n\n';
        // console.log('config:', currConfig, message);
        sendDataToServer();
        // setStage(stages.GAME);

        // const socket = new WebSocket(`ws://${serverUrl}`);
        // socket.addEventListener('open', (event) => {
        //     // Abre la conexión WebSocket
        //     console.log('Conexión abierta', 'enviando mensaje');
        //     // Envía el mensaje al servidor Go
        //     socket.send(message);
            
        // });

        // // Escucha las respuestas del servidor Go
        // socket.addEventListener('message', (event) => {
        //     console.log('Mensaje recibido del servidor:', event.data);
        //     // setResponse(event.data);
            
        //     // Cierra la conexión WebSocket después de recibir la respuesta
        //     socket.close();
        // });
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