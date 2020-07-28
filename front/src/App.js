import React from 'react';
import './styles.scss';
import { Container, Row, Col } from 'react-bootstrap'
import { Resources, Warehouse, Builds, Info } from './components';

function App() {
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
          <Resources />
        </Col>
        <Col md="4">
          <Warehouse />
        </Col>
        <Col md="4">
          <Builds />
        </Col>
      </Row>
      <Row>
        <Col>
          <Info />
        </Col>
      </Row>
    </Container>
  );
}

export default App;
