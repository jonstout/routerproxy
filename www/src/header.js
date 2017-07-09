import React, { Component } from 'react';

import './App.css';
import './header.css';

export class Header extends Component {
  constructor(props) {
    super(props);
    this.state = {
      networkTitleStyle: 'network-title'
    };

    this.onExecuteCommand = this.onExecuteCommand.bind(this);
  }

  componentDidMount() {
    window.addEventListener('scroll', this.handleScroll.bind(this));
  }

  componentWillUnmount() {
    window.removeEventListener('scroll', this.handleScroll.bind(this));
  }

  handleScroll(event) {
    var offset;
    if (event.path === undefined) {
      offset = event.pageY;
    } else {
      offset = event.path[1].pageYOffset;
    }

    if (offset > 75) {
      this.setState({networkTitleStyle: 'sm-network-title'});
    }
    else if (offset < 120) {
      this.setState({networkTitleStyle: 'network-title'});
    }
  }

  onExecuteCommand() {
    let args = {
      route_table: this.routeTable.value,
      selection:   this.command.value,
      text:        this.commandText.value
    };

    this.props.onExecuteCommand(args);
  }

  render() {
    return (
      <div className="header">

        <h1 className={this.state.networkTitleStyle}>
          Internet2 Advanced Layer 2
        </h1>
        <p className={this.state.networkTitleStyle}>
          a service of the Internet2 Network Operations Center
        </p>

        <div>
          <form className="form-inline header-elems">
            <select className="form-control outer-elem"
                    ref={(input) => { this.command = input; }}>
              <option>show version</option>
              <option>show bgp</option>
            </select>

            <select className="form-control center-elem"
                    ref={(input) => { this.routeTable = input; }}>
              <option></option>
              <option>default</option>
              <option>vrf</option>
            </select>

            <input type="text" className="form-control cmd-elem"
                   id="query"
                   placeholder="Command"
                   ref={(input) => { this.commandText = input; }} />

            <button type="button" onClick={this.onExecuteCommand} className="btn btn-default outer-elem">Execute</button>
          </form>
        </div>

      </div>
    );
  }
}
