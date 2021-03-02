import { Component } from 'react';
import { Grid, Form, Button } from 'semantic-ui-react';
import './App.css';
import config from './config.js';

import 'semantic-ui-css/semantic.min.css';

class App extends Component {

  constructor(props) {
    super(props);
    this.state = {
      error: "",
      message: "",
      history: [],
      ws: null,
    };  
  }

  componentDidMount() {
    this.connect()
  }

  connect() {
    var ws = new WebSocket(config.WS_URL);
    ws.onopen = () => {
      this.setState({ ws: ws });
    };

    ws.onclose = () => {
      this.setState({ws: null})
    }

    ws.onmessage = (event) => {
      if (!event.data) {
        return
      }
      this.setState({history: event.data.split(",")})
    };
  }

  sendMessage = () => {
    return this.state.ws.send(this.state.message);
  };

  handleChangeData = (event) => {
    this.setState({[event.target.name]: event.target.value});
  };

  render() {
    return (
      <Grid container style={{ padding: '2em 0em' }}>
        <Grid.Row>
          <Form.Field>
            <Form.Input
              placeholder="Message"
              name="message"
              onChange={(e) => this.handleChangeData(e)}/>
          </Form.Field>
        </Grid.Row>
        <Grid.Row>
          <Form.Field>
            <Button 
              positive 
              icon='checkmark' 
              labelPosition='right' 
              content='Send it!' 
              onClick={() => this.sendMessage()}/>
            </Form.Field>
        </Grid.Row>
        <Grid.Row>
          <ul>
            {this.state.history.map((value, index) => {
              return <li key={index}>{value}</li>
            })}
          </ul>
        </Grid.Row>
      </Grid>
    );
  }
}

export default App;
