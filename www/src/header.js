import React, { Component } from 'react';

import './App.css';
import './header.css';

export class Header extends Component {
  constructor(props) {
    super(props);
    this.state = {
      networkTitleStyle: 'network-title'
    };
  }

  componentDidMount() {
    window.addEventListener('scroll', this.handleScroll.bind(this));
  }

  componentWillUnmount() {
    window.removeEventListener('scroll', this.handleScroll.bind(this));
  }

  handleScroll(event) {
    console.log(event);

    var offset;
    if (event.path === undefined) {
      offset = event.pageY;
    } else {
      offset = event.path[1].pageYOffset;
    }

    console.log(offset);

    if (offset > 150) {
      this.setState({networkTitleStyle: 'sm-network-title'});
    }
    else if (offset < 120) {
      this.setState({networkTitleStyle: 'network-title'});
    }
  }

  render() {
    return (
      <div className="header">

        <h2 className={this.state.networkTitleStyle}>
          Internet2 Advanced Layer 2 <small>RouterProxy</small>
        </h2>
        <p className={this.state.networkTitleStyle}>
          A service of the Internet2 Network Operations Center
        </p>

        <div>
          <form className="form-inline header-elems">
            <select className="form-control outer-elem">
              <option>show version</option>
              <option>show bgp</option>
            </select>

            <select className="form-control center-elem">
              <option></option>
              <option>default</option>
              <option>vrf</option>
            </select>

            <input type="text" className="form-control cmd-elem" id="query" placeholder="Command"/>

            <button type="submit" className="btn btn-default outer-elem">Submit</button>
          </form>
        </div>

      </div>
    );
  }
}
