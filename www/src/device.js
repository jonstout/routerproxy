import React, { Component } from 'react';

import './device.css';

export class Device extends Component {
  constructor(props) {
    super(props);
    this.state = {
      id:       props.id,
      name:     props.name,
      address:  props.address,
      location: props.location
    };
  }

  render() {
    const collapseId  = "collapse-" + this.state.id.toString();
    const collapseRef = "#collapse-" + this.state.id.toString();

    const message = 'text\ntext\ntext';

    let status = 'ok';
    let statusSym = null;
    if (status === 'ok') {
        statusSym = (<span className="glyphicon glyphicon-ok"></span>);
    } else if (status === 'wait') {
      statusSym = (<span className="glyphicon glyphicon-refresh"></span>);
    } else if (status === 'error') {
      statusSym = (<span className="glyphicon glyphicon-remove"></span>);
    } else {
      let style = { color: 'transparent' };
      statusSym = (<span className="glyphicon glyphicon-ok" style={style}></span>);
    }

    return (
      <div>
        <div className="device-header">

          <div className="device-desc" data-toggle="collapse" href={collapseRef}>
            <h4 className="device-name">
              {statusSym} {this.state.name} - {this.state.address} <small>{this.state.location}</small>
            </h4>
          </div>

          <div className="device-toggle">
            <input type="checkbox" data-toggle="toggle" data-size="small" data-on="Run" data-off="Skip"/>
          </div>

        </div>

        <div className="device-content collapse" id={collapseId}>
          <pre>
            {message}
          </pre>
        </div>
      </div>
    );
  }
}

export class DeviceList extends Component {
  constructor(props) {
    super(props);
    this.state = {devices: []};
  }

  componentDidMount() {
    fetch('/api/device').then(function(response) {
      return response.json();
    }).then(function(json) {
      this.setState({devices: json});
    }.bind(this)).catch(function(error) {
      console.log(error);
    });
  }

  render() {
    const devices = this.state.devices.map(function(device) {
      return (<Device
              id={device.id}
              key={device.id}
              name={device.hostname}
              address={device.ip_address}
              location={device.location}/>);
    });

    return (
      <div className="device-list">
        <h2>Core Routers</h2>
        {devices}
      </div>
    );
  }
}
