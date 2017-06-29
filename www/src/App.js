import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';

import {Header} from './header';
import {DeviceList} from './device';

class App extends Component {
  constructor(props) {
    super(props);

    this.state = {
      devices: []
    };

    const socket = new WebSocket('ws://jonstout-dev7.grnoc.iu.edu:8080/api/ws');

    socket.addEventListener('open', function (event) {
      socket.send('Hello Server!');
    });

    socket.addEventListener('message', function (event) {
      console.log('Message from server', event.data);
    });
  }

  render() {
    return (
      <div className="App">
        <Header />

        <div className="content">
          <DeviceList />
        </div>
      </div>
    );
  }
}

export default App;
