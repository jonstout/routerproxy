import React, { Component } from 'react';

import './device.css';
import Toggle from 'react-bootstrap-toggle';

export class Device extends Component {
  constructor(props) {
    super(props);
    this.state = {
      toggleActive: false
    };

    this.onToggle = this.onToggle.bind(this);
  }

  onToggle() {
    this.setState({toggleActive: !this.state.toggleActive});
  }

  render() {
    const collapseId  = "collapse-" + this.props.id.toString();
    const collapseRef = "#collapse-" + this.props.id.toString();

    const message = "text\ntext\ntext";

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
      <div className="device-card">

        <div className="device-header">

          <div className="device-desc" data-toggle="collapse" href={collapseRef}>
            <h4 className="device-name">
              {statusSym} {this.props.name} - {this.props.address} <small>{this.props.location}</small>
            </h4>
          </div>

          <div className="device-toggle">
            <Toggle onClick={this.onToggle} on="Run" off="Skip" size="sm" active={this.state.toggleActive} />
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

  tick() {
    console.log(this.state);
    fetch('/api/device').then(function(response) {
      return response.json();
    }).then(function(json) {
      this.setState((prevState) => {
        return {devices: json};
      });
    }.bind(this)).catch(function(error) {
      console.log(error);
    });
  }

  componentDidMount() {
    this.tick();
    this.interval = setInterval(() => this.tick(), 3000);
  }

  render() {
    const devices = this.state.devices.map((device) => {
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
