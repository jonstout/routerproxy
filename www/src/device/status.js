import React, { Component } from 'react';

import './status.css';

export class DeviceStatus extends Component {
  constructor(props) {
    super(props);
  }

  render() {
    let style  = { color: 'transparent' };
    let classes = 'glyphicon glyphicon-remove';

    if (this.props.status === 'ok') {
      style   = {};
      classes = 'glyphicon glyphicon-ok';
    }

    if (this.props.status === 'load') {
      style   = {};
      classes = 'glyphicon glyphicon-refresh loading';
    }

    if (this.props.status === 'error') {
      style   = {};
      classes = 'glyphicon glyphicon-remove';
    }

    return (<span className={classes} style={style}></span>);
  }
}
