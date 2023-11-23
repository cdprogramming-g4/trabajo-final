// import logo from './logo.svg';
import './css/App.css';
import { useEffect, useState } from 'react';
import Config from './components/Config';
import Game from './components/Game';
import Stats from './components/Stats';
import Credits from './components/Credits';
// import io from 'socket.io-client';

const stages = {
  CONFIG: 0,
  GAME: 1,
  STATS: 2,
  CREDITS: 3,
};

function App() {
  const [stage, setStage] = useState(stages.CONFIG);
  const [cellJsx, setCellJsx] = useState(<></>);

  useEffect(()=>{
    console.log('stage', stage);
    switch(stage) {
      case stages.CONFIG:
        setCellJsx(<Config setStage={setStage}/>);
      break;
      case stages.GAME:
        setCellJsx(<Game setStage={setStage}/>);
      break;
      case stages.STATS:
        setCellJsx(<Stats setStage={setStage}/>);
      break;
      case stages.CREDITS:
        setCellJsx(<Credits setStage={setStage}/>);
        break;
      default:
        setCellJsx(<></>);
      break;
    }
  }, [stage]);

  return (
    <div className="App">
      <header className="header">
      </header>
      <main>
        {cellJsx}
      </main>
    </div>
  );
}

export default App;
export {stages};