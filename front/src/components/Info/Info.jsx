import React, { useState, useEffect } from 'react';
import { Alert, Row, Col } from 'react-bootstrap'
import {
  FarmerEmoji,
  WorkerEmoji,
} from '../../emojis/'

const Info = () => {
  const [time, setTime] = useState(0);

  useEffect(() => {
    const interval = setInterval(() => {
      setTime((time) => time + 1)
    }, 1000)

    clearInterval(time)
  }, [])

  
  return (
    <Alert variant="info">
      <Row>
        <Col className="d-flex justify-content-center">
          <h1 className="align-self-center">
            Time {time}
          </h1>
        </Col>
        <Col className="mx-2">
          <h4>Goroutines</h4>
          <p> Initial: 1000 </p>
          <p> Current: 1000 </p>
        </Col>
        <Col className="mx-2">
          <h4>Jobs</h4>
          <p> In progress: 42 </p>
          <p> Finished: 42 </p>
        </Col>
        <Col className="mx-2">
          <h4>Total workers</h4>
          <p>
            <FarmerEmoji/> [10]
          </p>
          <p>
            <WorkerEmoji/> [4]
          </p>
        </Col>
      </Row>
    </Alert>
  )
}

export default Info;
