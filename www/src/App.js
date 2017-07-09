import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';

import {Header} from './header';
import {DeviceList} from './device';

class App extends Component {
  constructor(props) {
    super(props);

    this.state = {};

    this.socket = new WebSocket('ws://jonstout-dev7.grnoc.iu.edu:8080/api/ws');

    this.socket.addEventListener('open', function (event) {
      console.log('Connected to websocket!');
    }.bind(this));

    this.socket.addEventListener('message', function (event) {
      const result = JSON.parse(event.data);

      let device = this.state[result.device_id];
      device.response = result.output;
      device.responseStatus = 'ok';

      this.setState({ [device.id]: device });
    }.bind(this));

    this.setIsSelected  = this.setIsSelected.bind(this);
    this.executeCommand = this.executeCommand.bind(this);
    this.getDevices     = this.getDevices.bind(this);
  }

  componentDidMount() {
    this.getDevices();
    this.interval = setInterval(() => this.getDevices(), 3000);
  }

  setIsSelected(deviceID, isSelected) {
    let device = this.state[deviceID];
    device.isSelected = isSelected;

    this.setState({ [deviceID]: device });
  }

  getDevices() {
    fetch('/api/device').then(function(response) {
      return response.json();
    }).then(function(json) {
      for (let i = 0; i < json.length; i++) {
        let deviceID = json[i].id;
        let device   = this.state[deviceID];

        if (device === undefined) {
          device = json[i];
          device.response = '';
          device.responseStatus = null;
          device.isSelected = false;
        }

        this.setState({ [deviceID]: device });
        console.log(this.state);
      }
    }.bind(this)).catch(function(error) {
      console.log(error);
    });
  }

  executeCommand(command) {
    let deviceIDs = Object.values(this.state).filter(function(device) {
      if (!device.isSelected && device.responseStatus === 'ok') {
        device.responseStatus = null;
      } else if (device.isSelected) {
        device.responseStatus = 'load';
      }

      this.setState({ [device.id]: device });
      return device.isSelected;
    }.bind(this)).map(function(device) {
      return device.id;
    });

    const req = {
      command: command,
      devices: deviceIDs
    };

    console.log(req);
    this.socket.send(JSON.stringify(req));
  }

  render() {
    const devices = Object.values(this.state);

    return (
      <div className="App">
        <Header onExecuteCommand={this.executeCommand}/>

        <div className="content">
          <DeviceList devices={devices} setIsSelected={this.setIsSelected}/>
        </div>
      </div>
    );
  }
}

export default App;
