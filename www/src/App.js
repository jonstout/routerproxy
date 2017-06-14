import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';

import {Header} from './header';
import {DeviceList} from './device';

class App extends Component {
  constructor(props) {
    super(props);

    let msg = 'stuff and\nstuff and stuff and\nstuff and\nstuff and stuff and\nstuff and\nstuff and stuff and\n';

    this.state = {
      devices: [
        {
          id:      1,
          name:    'app-1.sample.com',
          address: '127.0.0.1',
          location: {
            state: 'IN',
            city:  'Bloomington'
          },
          response: {
            status:  'ok',
            message: msg
          }
        },
        {
          id:      2,
          name:    'server-1.sample.com',
          address: '127.0.0.1',
          location: {
            state: 'IN',
            city:  'Bloomington'
          },
          response: {
            status:  'wait',
            message: msg
          }
        },
        {
          id:      3,
          name:    'server-2.sample.com',
          address: '127.0.0.1',
          location: {
            state: 'IN',
            city:  'Bloomington'
          },
          response: {
            status:  'error',
            message: msg
          }
        },
        {
          id:      4,
          name:    'server-4.sample.com',
          address: '127.0.0.1',
          location: {
            state: 'IN',
            city:  'Bloomington'
          },
          response: {
            status:  '',
            message: msg
          }
        }
      ]
    };
  }

  putSelectedDevice(deviceId) {

  }

  delSelectedDevice(deviceId) {

  }

  render() {
    return (
      <div className="App">
        <Header />

        <div className="content">
          <DeviceList devices={this.state.devices} />
          <DeviceList devices={this.state.devices} />
          <DeviceList devices={this.state.devices} />
          <DeviceList devices={this.state.devices} />
          <DeviceList devices={this.state.devices} />
          <DeviceList devices={this.state.devices} />
        </div>
      </div>
    );
  }
}

export default App;
