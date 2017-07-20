import React, { Component } from 'react';

import Toggle from 'react-bootstrap-toggle';

import './device.css';

import { DeviceStatus } from './device/status.js';

export class Device extends Component {
  constructor(props) {
    super(props);
    this.state = {};

    this.onToggle = this.onToggle.bind(this);
  }

  onToggle() {
    this.props.setIsSelected(this.props.device.id, !this.props.device.isSelected);
  }

  render() {
    const device = this.props.device;

    const collapseId  = "collapse-" + device.id.toString();
    const collapseRef = "#collapse-" + device.id.toString();

    return (
      <div className="panel panel-default">

        <div className="device-header">
          <div className="device-desc" data-toggle="collapse" href={collapseRef}>
            <h4 className="device-name">
              <DeviceStatus status={device.responseStatus}/> {device.hostname} - {device.ip_address} <small>{device.location}</small>
            </h4>
          </div>

          <div className="device-toggle">
            <Toggle onClick={this.onToggle} on="Run" off="Skip" size="sm" active={device.isSelected} />
          </div>
        </div>

        <div className="device-content collapse" id={collapseId}>
          <pre>
            {device.response}
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

  render() {
    const devices = this.props.devices.map((device) => {
      return (<Device
              key={device.id}
              device={device}
              setIsSelected={this.props.setIsSelected}/>);
    });

    return (
      <div className="device-list">
        <h2 className="device-group-title">Routers</h2>
        {devices}
      </div>
    );
  }
}
