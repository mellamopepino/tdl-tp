import React from 'react';
import { Alert, Row, Col, Button } from 'react-bootstrap'
import { useSelector } from 'react-redux';
import {
  FarmerEmoji,
  WorkerEmoji,
} from '../../emojis/'

const Info = ({ start }) => {
  const info = useSelector((state) => state.info)

  function startWebSocket() {
    start()
  }

  return (
    <Alert variant="info">
      <Row>
        <Col md={6} className="d-flex flex-column justify-content-around align-items-center">
          <h1 className="align-self-center">
            {info.totalTime && `Time: ${info.totalTime}` || `Working...`}
          </h1>
          <Button variant="primary" onClick={startWebSocket}>Start</Button>
        </Col>
        <Col className="mx-2">
          <h4>Jobs</h4>
          <p> In progress: {info.jobsInProgress} </p>
          <p> Finished: {info.jobsFinished} </p>
          <p> Failed builds: {info.failedBuilds} </p>
        </Col>
        <Col className="mx-2">
          <h4>Total workers</h4>
          <p>
            <FarmerEmoji/> [{info.totalGatherers}]
          </p>
          <p>
            <WorkerEmoji/> [{info.totalWorkers}]
          </p>
        </Col>
      </Row>
    </Alert>
  )
}

export default Info;
