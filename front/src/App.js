import React from 'react';
import './styles.scss';
import { Container, Row, Col } from 'react-bootstrap'
import { Resources, Warehouse, Builds, Info } from './components';
import { useDispatch } from 'react-redux';
import getAction from './store/actions'

function App() {
  const dispatch = useDispatch()
  let ws = null

  function startWebSocket(){
    ws = new WebSocket("ws://127.0.0.1:8080/send");
    ws.onmessage = (e) => {
      dispatch(getAction(e.data))
    }
  }

  return (
    <Container>
      <Row>
        <Col>
          <h1 className="title">Age of empires</h1>
          <hr/>
        </Col>
      </Row>
      <Row className="my-5">
        <Col md="4">
          <Resources ws={ws}/>
        </Col>
        <Col md="4">
          <Warehouse ws={ws}/>
        </Col>
        <Col md="4">
          <Builds ws={ws}/>
        </Col>
      </Row>
      <Row>
        <Col>
          <Info start={startWebSocket} />
        </Col>
      </Row>
    </Container>
  );
}

export default App;
